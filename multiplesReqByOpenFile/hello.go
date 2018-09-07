package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
)

func sendRequests(url string) {
	httpURL := "http://" + url
	fmt.Println(httpURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", httpURL, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0)`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("[+] Status: ", resp.StatusCode)

}

func openFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Println("err", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		sendRequests(scanner.Text())
	}

}

func main() {
	fmt.Println(os.Args[1])
	fmt.Println(reflect.TypeOf(os.Args[1]))
	openFile(os.Args[1])
}
