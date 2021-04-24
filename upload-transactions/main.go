package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
Script to unload a transaction to payment workflow  Gigaspaces

File sample
transactionId;orderId
1d344150-bfa5-4690-8e49-213cfe751cb4;1150486395
6fbca7ca-c082-4fbc-8b31-307579c6cb42;392652473
*/

var FILE_PATH string
var URL string
var USER string
var PASSWORD string
var HEADERS string
var SEPARATOR string
var WORKERS_NUMBER int

func main() {

	FILE_PATH = "fix_txs.csv"
	URL = "activemq-houston.payulatam.com:8161/api/message?destination=queue://pps.acknowledge.notification"
	USER = "admin"
	PASSWORD = "4QZboDsXmjMd7"
	HEADERS = "{'content-type': 'application/json'}"
	SEPARATOR = ";"
	WORKERS_NUMBER = 3

	textFile := readFile()

	headerLine := textFile[0]
	fmt.Println("Sending ", len(textFile), " records...")
	println(headerLine)

	startTime := time.Now()

	processFile(textFile)

	finalTime := time.Since(startTime)
	fmt.Println("\nThe process has finished")
	fmt.Printf("Time: %.2f seconds", float64(finalTime)/float64(time.Second))
}

func processFile(textFile []string) {

	jobs := make(chan map[string]string, 100)
	results := make(chan string, 100)

	// Define only WORKERS_NUMBER
	for i := 0; i < WORKERS_NUMBER; i++ {
		go doPost(jobs, results)
	}

	textFile = textFile[1:]
	for _, fileLine := range textFile {
		splitLine := strings.Split(fileLine, SEPARATOR)
		message := map[string]string{
			"transactionId": splitLine[0],
			"orderId":       splitLine[1],
		}
		jobs <- message
	}
	close(jobs)

	// collect all the results of the work.
	for a := 1; a <= len(textFile); a++ {
		fmt.Println(<-results)
	}

}

func readFile() []string {
	fmt.Println("starting...")

	file, err := os.Open(FILE_PATH)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textFile []string

	for scanner.Scan() {
		textFile = append(textFile, scanner.Text())
	}
	file.Close()
	return textFile
}

func doPost(jobs <-chan map[string]string, results chan<- string) {

	for j := range jobs {

		bytesmessage, err := json.Marshal(j)
		check(err)

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.Post("https://"+USER+":"+PASSWORD+"@"+URL, HEADERS, bytes.NewBuffer(bytesmessage))
		if err != nil {
			log.Fatal("Error sending message: "+string(bytesmessage), err)
			check(err)
		}
		results <- "Sending:" + string(bytesmessage) + " got http status code: " + strconv.Itoa(resp.StatusCode)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
