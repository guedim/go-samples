package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Script to unload a transaction to payment workflow  Gigaspaces

File sample 1
transactionId;orderId
1d344150-bfa5-4690-8e49-213cfe751cb4;1150486395
6fbca7ca-c082-4fbc-8b31-307579c6cb42;392652473

File sample 2
orderId;transactionId
1150486395;1d344150-bfa5-4690-8e49-213cfe751cb4
392652473;6fbca7ca-c082-4fbc-8b31-307579c6cb42
*/

func main() {

	// Change this
	FILE_PATH := "/Users/mario.guerrero/Downloads/fix_txs.csv"
	//URL := "activemq-houston.payulatam.com:8161/api/message?destination=queue://pps.acknowledge.notification"

	//ORDER_ID_NAME := "orderId"             //Case insensitive
	//TRANSACTION_ID_NAME := "transactionId" //Case insensitive
	//USER := "admin"
	//PASSWORD := "4QZboDsXmjMd7"
	//HEADERS := "{'content-type': 'application/json'}"

	openFile(FILE_PATH)

}

func openFile(filePath string) {
	fmt.Println("starting...")

	file, err := os.Open("sample.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
