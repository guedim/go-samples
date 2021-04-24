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

func main() {

	FILE_PATH = "fix_txs.csv"
	URL = "activemq-houston.payulatam.com:8161/api/message?destination=queue://pps.acknowledge.notification"
	USER = "admin"
	PASSWORD = "4QZboDsXmjMd7"
	HEADERS = "{'content-type': 'application/json'}"
	SEPARATOR = ";"

	processFile()
}

func processFile() {
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

	headerLine := textFile[0]
	fmt.Println("Sending ", len(textFile), " records...")
	println(headerLine)

	positionTransactionId := 0
	positionOrderId := 1
	startTime := time.Now()

	textFile = textFile[1:]
	for _, fileLine := range textFile {
		splitLine := strings.Split(fileLine, SEPARATOR)
		doPost(splitLine[positionTransactionId], splitLine[positionOrderId])
	}

	finalTime := time.Since(startTime)
	fmt.Println("\nThe process has finished")
	fmt.Printf("Time: %.2f seconds", float64(finalTime)/float64(time.Second))

}

func doPost(transactionId string, orderId string) {

	message := map[string]interface{}{
		"transactionId": transactionId,
		"orderId":       orderId,
	}

	bytesmessage, err := json.Marshal(message)
	check(err)

	fmt.Print("Sending: ", string(bytesmessage))

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post("https://"+USER+":"+PASSWORD+"@"+URL, HEADERS, bytes.NewBuffer(bytesmessage))
	if err != nil {
		log.Fatal("Error sending orderId: "+string(orderId), err)
		check(err)
	}
	fmt.Println(" got http status code: ", resp.StatusCode)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
