package requestrepeater

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	WordListReader "github.com/michaelknudsen/WordListReader/wordlistreader"
	"github.com/michaelknudsen/api-buster/src/flagparser"
	"github.com/michaelknudsen/api-buster/src/methods"
	"github.com/michaelknudsen/api-buster/src/outpututil"
	"github.com/michaelknudsen/api-buster/src/requestformatter"
)

func Do(wlr *WordListReader.WordListReader, wg *sync.WaitGroup, resultchan chan<- outpututil.ResultList) {

	defer wg.Done()
	for word := range wlr.Iter() {
		results := outpututil.ResultList{}
		for _, method := range methods.MethodList {
			r, err := requestformatter.FormatRequest(word, method)
			if err != nil {
				fmt.Printf("Error formatting request for %s :%s\n", word, method)
				continue
			}

			client := http.Client{}
			res, err := client.Do(r)
			if err == nil {
				res.Body.Close()
				//search
				var bi int
				for _, v := range flagparser.BlackList {
					if v == res.StatusCode {
						bi = 1
						break
					}
				}
				//bi := sort.SearchInts(flagparser.BlackList, res.StatusCode)
				var wi int
				for _, v := range flagparser.WhiteList {
					if v == res.StatusCode {
						wi = 1
						break
					}
				}
				//wi := sort.SearchInts(flagparser.WhiteList, res.StatusCode)
				//fmt.Println(bi, wi, res.StatusCode)
				if bi <= 0 && (wi > 0 || len(flagparser.WhiteList) == 0) {

					results = append(results, outpututil.ParseResponseIntoResult(res, r.URL.String(), method))
				}

			}
			if flagparser.Duration > 0 {
				time.Sleep(flagparser.Duration)
			}

		}
		resultchan <- results

	}
}
