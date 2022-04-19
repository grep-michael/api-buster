package requestformatter

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	flagparser "github.com/michaelknudsen/api-buster/src/flagparser"
)

func FormatRequest(path string, method string) (*http.Request, error) {
	//body := strings.NewReader("")
	r := &http.Request{
		Header: make(http.Header), //init empty header
	}
	var err error

	if flagparser.Forceterminal {
		path += "/"
	}
	if flagparser.Cookies != "" {
		r.Header.Add("cookies", flagparser.Cookies)
	}
	if flagparser.Password != "" && flagparser.Username != "" {
		data := flagparser.Username + ":" + flagparser.Password
		data = base64.StdEncoding.EncodeToString([]byte(data))
		r.Header.Add("Authorization", "Basic "+data)
	}
	if flagparser.Whitelist != "" {

	}
	if flagparser.Blacklist != "" {

	}

	if len(flagparser.Headers) > 0 {
		fmt.Println(flagparser.Headers)
	}

	r.URL, err = url.Parse(flagparser.Url + path)
	r.Method = method
	return r, err
}
