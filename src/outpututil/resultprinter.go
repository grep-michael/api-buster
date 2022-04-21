package outpututil

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/michaelknudsen/api-buster/src/flagparser"
)

type ResultPrinter struct {
	file     io.Writer
	openFile sync.Once
}

func ParseResponseIntoResult(res *http.Response, url string, method string) Result {
	return Result{
		Method:     method,
		Url:        url,
		StatusCode: res.StatusCode,
	}
}

func (rp *ResultPrinter) initFile() {
	if flagparser.Output != "" {
		//file specified
		f, err := os.OpenFile(flagparser.Output, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		f.Write([]byte("["))
		rp.file = f

	} else {
		rp.file = os.Stdout
	}
}
func (rp *ResultPrinter) Close() {
	if flagparser.Output != "" {
		rp.file.(*os.File).WriteString("]")
		rp.file.(*os.File).Close()
	}
}

func (rp *ResultPrinter) PrintResultList(results ResultList) {
	rp.openFile.Do(rp.initFile)
	for _, v := range results {
		fmt.Fprintf(rp.file, "{Url:\"%s\",Method:\"%s\",Status:%d},\n", v.Url, v.Method, v.StatusCode)
	}

}
