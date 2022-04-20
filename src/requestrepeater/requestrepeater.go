package requestrepeater

import (
	"fmt"
	"net/http"
	"sync"

	WordListReader "github.com/michaelknudsen/WordListReader/wordlistreader"
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
				defer res.Body.Close()
				results = append(results, outpututil.ParseResponseIntoResult(res, r.URL.String(), method))
			}

		}
		resultchan <- results

	}
}
