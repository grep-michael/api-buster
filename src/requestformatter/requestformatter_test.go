package requestformatter

import (
	"testing"

	flagparser "github.com/michaelknudsen/api-buster/src/flagparser"
)

func formatArgs_1() {
	flagparser.Url = "http://localhost:8080/"
	flagparser.Wordlist = "../rockyou.txt"
}

func TestRequestFormatter_1(t *testing.T) {

	_, err := FormatRequest("test", "GET")
	if err != nil {
		t.Error(`error in testRequestFormatter_1`)
	}
}

func formatArgs_2() {
	flagparser.Url = "http://localhost:8080/"
	flagparser.Wordlist = "../rockyou.txt"
	flagparser.Password = "password"
	flagparser.Username = "username"
}

func TestRequestFormatter_2(t *testing.T) {
	formatArgs_2()

	r, err := FormatRequest("test", "GET")
	if err != nil {
		t.Error(`error in testRequestFormatter_2`)
	}
	if r.Header["Authorization"] == nil {
		t.Error(`Authorization Header is nil`)
	}
	if r.Header.Get("Authorization") != "Basic dXNlcm5hbWU6cGFzc3dvcmQ=" {
		t.Errorf("Authorization header is: %s, should be \"Basic dXNlcm5hbWU6cGFzc3dvcmQK \"", r.Header.Get("Authorization"))
	}
}
