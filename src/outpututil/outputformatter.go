package outpututil

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/michaelknudsen/api-buster/src/flagparser"
)

func ParseResponseIntoResult(res *http.Response, url string, method string) Result {
	return Result{
		Method:     method,
		Url:        url,
		StatusCode: res.StatusCode,
	}
}

func PrintResultList(result ResultList) {
	if flagparser.Output != "" {
		//file specified
		f, err := os.Create(flagparser.Output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		printResultList(result, f)

	} else {
		printResultList(result, os.Stdout)
	}

}

func printResultList(results ResultList, out io.Writer) {
	for _, v := range results {
		fmt.Fprintf(out, "{Url:%s,Method:%s,Status:%d}\n", v.Url, v.Method, v.StatusCode)
	}

}
