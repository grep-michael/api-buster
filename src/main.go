package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"os"
	WordListReader "github.com/michaelknudsen/WordListReader/wordlistreader"
	flagparser "github.com/michaelknudsen/api-buster/src/flagparser"
	requestformatter "github.com/michaelknudsen/api-buster/src/requestformatter"
)
const (
	THREAD_COUNT = 10
)
func main() {
	flagparser.Init()
	//fmt.Println(flagparser.Headers)
	
	if flagparser.Wordlist == "" || flagparser.Url == "" {
		fmt.Println("Missing wordlist and/or url")
		os.Exit(1)
	}

	wlr := WordListReader.MakeNewWordListReader(flagparser.Wordlist)
	defer wlr.Close()
	s,_ :=  wlr.ReadLine()
	requestformatter.FormatRequest(s,"GET")
	/*
	var wg sync.WaitGroup
	for i:=0;i<THREAD_COUNT;i++{
		wg.Add(1)

		wg.Done()
	}
	wg.Wait()*/
	
}
// poc
func sendReq(wg *sync.WaitGroup, wlr *WordListReader.WordListReader, id int) {
	defer wg.Done()
	methods := [...]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}
	for word := range wlr.Iter() {
		fmt.Println(id, word)
		for _, v := range methods {
			reader := strings.NewReader("")
			r, _ := http.NewRequest(v, "http://localhost:8080/"+word, reader)
			client := http.Client{}
			client.Do(r)
		}

	}
}
