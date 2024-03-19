package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func getContent(link string, secure bool) string {

	// Configure transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: secure},
	}

	// Configure client
	client := &http.Client{Transport: tr}

	// Get link stop if error
	res, err := client.Get(link)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(content)

}

func main() {

	//fmt.Println("Secure true, trusted website")
	//fmt.Println(getContent("https://medium.com/@func25/golang-secret-how-to-add-default-values-to-function-parameters-60bd1e9625dc", true))

	fmt.Println("Secure true, non-trusted website")
	fmt.Println(getContent("https://cybershoke.net/cs2/servers/dm", true))

}
