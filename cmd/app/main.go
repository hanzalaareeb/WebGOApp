package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// link- https://go.dev/doc/tutorial/create-module
	url := "https://go.dev/doc/tutorial/create-module"
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	// Check http response code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status code %d:", resp.StatusCode)
	}

	// read the parsed body
	bytes, err := io.ReadAll(resp.Body)
	check(err)

	// Write the decoded bytes to new file
	file, err := os.Create("go-doc.html")
	check(err)
	defer file.Close()

	// write parsed data to file
	_, err = file.Write(bytes)
	check(err)
	fmt.Println("writing ...")

	/*

	 */

}
