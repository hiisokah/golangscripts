package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileExists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("(+) File exists")
			return true
		}
	}
	return false
}

func listDir(path string, pathReport string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("(+) error open file", err)
	}
	for _, f := range files {
		line := f.Name()
		hashOfFile := computeHash(line)
		writeHashToFile(pathReport, hashOfFile)
	}
}

func computeHash(line string) string {
	hash := md5.New()
	hash.Write([]byte(line))
	fmt.Printf("(+) File: %s --> Hash: %x \n", line, hash.Sum(nil))
	return hex.EncodeToString(hash.Sum(nil))

}

func writeHashToFile(path string, hash string) {

	if fileExists(path) {
		f, _ := os.Create(path)
		fmt.Printf("(+) File created %v", f)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("(+) Error write file", err)
	}
	defer f.Close()

	_, err = f.WriteString(hash + "\n")
	if err != nil {
		log.Println("(+) Error write file", err)
	}
}

func readFile(file string) {
	fmt.Println("(+) Reading file ", file)
	f, err := os.Open(file)
	if err != nil {
		log.Println("(+) Error reading file", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s \n", line)
	}
}

func main() {
	dir := os.Args[1]
	if len(dir) > 1 && len(dir) < 1 {
		fmt.Println("(+) provide one dir to continue: ./main.go /home/$USER/Documents")
		os.Exit(1)
	}

	//oldHashFile := oldHash(dir)
	reportFile := os.Args[1] + "/analysis_report.txt" // computing hashes
	fileOutPut := os.Args[1] + "/file_list.txt"       // used to store the list of hashes of file
	oldFileOutPut := os.Args[1] + "/file_list.old"    // used to store the list of hashes of file

	if !fileExists(reportFile) {
		fmt.Println("(+) not exists")
	}

	if !fileExists(fileOutPut) {
		fmt.Println("(+) not exists")
	}

	if !fileExists(oldFileOutPut) {
		fmt.Println("(+) not exists")
	}

	//fmt.Printf("%x \n", oldHashFile)
	listDir(os.Args[1], reportFile)
	readFile(reportFile)

}
