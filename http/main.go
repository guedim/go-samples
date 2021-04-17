package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//Make a byte slide and full it with all information
	bs := make([]byte, 999999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))

	// Using io copy
	io.Copy(os.Stdout, resp.Body)
}
