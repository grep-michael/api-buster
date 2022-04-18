package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	WordListReader "github.com/michaelknudsen/WordListReader/wordlistreader"
	flagparser "github.com/michaelknudsen/api-buster/src/flagparser"
)

func main() {
	flagparser.Init()
	fmt.Println(flagparser.Headers)

	
	/*
	var wg sync.WaitGroup
	wlr := WordListReader.MakeNewWordListReader("rockyou.txt")
	defer wlr.Close()
	wg.Add(1)
	go sendReq(&wg, &wlr, 1)
	wg.Add(1)
	go sendReq(&wg, &wlr, 2)
	wg.Add(1)
	go sendReq(&wg, &wlr, 3)
	wg.Wait()*/
}
// poc
func sendReq(wg *sync.WaitGroup, wlr *WordListReader.WordListReader, id int) {
	defer wg.Done()
	methods := [...]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}
	for word := range wlr.Iter() {
		fmt.Println(id, word)
		for _, v := range methods {
			reader := strings.NewReader("Test Body")
			r, _ := http.NewRequest(v, "http://localhost:8080/"+word, reader)
			client := http.Client{}
			client.Do(r)
		}

	}
}
