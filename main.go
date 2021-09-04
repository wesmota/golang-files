package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	// reading directory
	dirname := "Datasets-master"
	var wg sync.WaitGroup
	c := make(chan string)
	f, err := os.Open(dirname)
	var data string
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	for _, file := range files {
		wg.Add(1)
		go readFileContent(dirname, file.Name(), &wg, c)
		data = <-c
		fmt.Println(data)

	}
	wg.Wait()
	close(c)
	finish := time.Since(start)
	fmt.Println(finish.Seconds())

}

func readFileContent(dirname string, filename string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	//fmt.Println(filename)
	c <- filename
	_, err := ioutil.ReadFile(dirname + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))
	// fmt.Println("***********************************************")
}

func createCSVFile(dirname string, filename string, wg *sync.WaitGroup) {

}
