package requestformatter

/*
options that effect the request

Forcerterminal
Cookies
Password
Username
Url
Headers
*/

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

func formatArgs_3() {
	flagparser.Url = "http://localhost:8080/"
	flagparser.Wordlist = "../rockyou.txt"
	flagparser.Password = "password"
	flagparser.Username = "username"
	flagparser.Headers = append(flagparser.Headers, "Connection:close")
	flagparser.Headers = append(flagparser.Headers, "Keep-Alive:timeout=5, max=1000")
}

func TestRequestFormatter_3(t *testing.T) {
	formatArgs_3()
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
	if len(flagparser.Headers) < 1 {
		t.Error("flagparser.Headers length less than 1, should be 2")
	}
	if r.Header.Get("Connection") != "close" {
		t.Error("Connection header malformed")
	}
	if r.Header.Get("Keep-Alive") != "timeout=5, max=1000" {
		t.Error("Keep-Alive header malformed")
	}
}
