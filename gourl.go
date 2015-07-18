package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	reqBody := ""
	if len(os.Args) > 3 {
		reqBody = os.Args[3]
	}
	req, err := http.NewRequest(os.Args[1], os.Args[2], bytes.NewBufferString(reqBody))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}
