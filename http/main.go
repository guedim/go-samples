package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {

	// option 1:  make a byte slide and full it with all information
	resp := doGoogleGet()
	bs := make([]byte, 999999)
	resp.Body.Read(bs)
	print(1, bs)

	// option 2:  Using  ioutils Read All
	resp = doGoogleGet()
	bs, err := ioutil.ReadAll(resp.Body)
	validateError(err)
	print(2, bs)

	// option 3:  Using  custom copy interface
	resp = doGoogleGet()
	lw := logWritter{}
	io.Copy(lw, resp.Body)
}

func (logWritter) Write(bs []byte) (int, error) {
	print(3, bs)
	return len(bs), nil
}

func doGoogleGet() *http.Response {
	resp, err := http.Get("https://google.com")
	validateError(err)
	return resp

}

func validateError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func print(n int, bs []byte) {
	fmt.Println("**************		OPTION ", n, "	******************")
	fmt.Println(string(bs))
}
