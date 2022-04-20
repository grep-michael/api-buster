package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	WordListReader "github.com/michaelknudsen/WordListReader/wordlistreader"
	flagparser "github.com/michaelknudsen/api-buster/src/flagparser"
	"github.com/michaelknudsen/api-buster/src/outpututil"
	"github.com/michaelknudsen/api-buster/src/requestrepeater"
)

const (
	THREAD_COUNT = 5
)

func main() {
	flagparser.Init()
	// I prefer explicit init function rather than built in init function

	if flagparser.Wordlist == "" || flagparser.Url == "" {
		fmt.Println("Missing wordlist and/or url")
		os.Exit(1)
	}

	wlr := WordListReader.MakeNewWordListReader(flagparser.Wordlist)
	defer wlr.Close()
	rlistener := outpututil.ResultListener{}
	rlistener.Init()
	var wg sync.WaitGroup
	for i := 0; i < THREAD_COUNT; i++ {
		wg.Add(1)
		go requestrepeater.Do(&wlr, &wg, rlistener.GetResultChannel())
	}
	go rlistener.Listen()
	wg.Wait()
	rlistener.Done()
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
