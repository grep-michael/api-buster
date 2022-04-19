package flagparser

import (
	"flag"
	"strings"
	"time"
)

type headerflags []string

func (i headerflags) String() string {
	var sb strings.Builder
	for _, v := range i {
		sb.WriteString(v)
		sb.WriteString(" ")
	}
	return sb.String()
}
func (i *headerflags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var Headers headerflags
var Forceterminal bool
var Cookies string
var Password string
var Whitelist string
var Blacklist string
var Url string
var Username string
var Output string
var Wordlist string
var Duration time.Duration

func Init() {
	flag.BoolVar(&Forceterminal, "f", false, "End each url with /")
	flag.BoolVar(&Forceterminal, "add-slash", false, "End each url with /")

	flag.StringVar(&Cookies, "c", "", "Cookies for request, semicolon seperated")
	flag.StringVar(&Cookies, "cookies", "", "Cookies for request, semicolon seperated")

	flag.StringVar(&Password, "P", "", "Password for Baisc Auth")
	flag.StringVar(&Password, "password", "", "Password for Baisc Auth")

	flag.StringVar(&Whitelist, "W", "", "status code White list (if set only codes listed here will be shown)")
	flag.StringVar(&Whitelist, "white-list", "", "status code White list (if set only codes listed here will be shown)")

	flag.StringVar(&Blacklist, "B", "400,404", "status code blacklist (default, will not show these codes; defaut: 404,400)")
	flag.StringVar(&Blacklist, "black-list", "400,404", "status code blacklist (default, will not show these codes; defaut: 404,400)")

	flag.StringVar(&Url, "u", "", "The target url")
	flag.StringVar(&Url, "url", "", "The target url")

	flag.StringVar(&Username, "U", "", "Username for basic Auth")
	flag.StringVar(&Username, "username", "", "Username for basic Auth")

	flag.StringVar(&Output, "o", "", "Output file of type json")
	flag.StringVar(&Output, "output", "", "Output file of type json")

	flag.StringVar(&Wordlist, "w", "", "Path to wordlist")
	flag.StringVar(&Wordlist, "wordlist", "", "Path to wordlist")

	flag.Var(&Headers, "H", "Specify HTTP headers, -H 'Header1: val1' -H 'Header2: val2")

	flag.DurationVar(&Duration, "-delay", 0*time.Second, "Time each waits between request")

	flag.Parse()
}
