package requestformatter

import (
	flagparser "github.com/michaelknudsen/api-buster/src/flagparser"
	"net/http"
)

func FormatRequest(path string, method string) error{
	//body := strings.NewReader("")
	r, err := http.NewRequest(method, flagparser.Url+path, nil)
	
	if err != nil{
		return err
	}

	client := http.Client{}
	client.Do(r)
	return nil
}