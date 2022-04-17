
# API-Buster
# This project is not finished/started yet only the readme has been finished

A buster designed specifically for brute forcing apiendpoints

https://github.com/OJ/gobuster/blob/master/README.md
Used as template for commands
```text
Usage:
  ./api-buster [flags]

Flags:
  -f, --add-slash               Append / to each request
  -c, --cookies string          Cookies to use for the requests
  -H, --headers stringArray     Specify HTTP headers, -H 'Header1: val1' -H 'Header2: val2'
  -P, --password string         Password for Basic Auth
  -W, --white-list string       status code white list (if set only codes listed here will be shown)
  -B, --blacklist string        status code blacklist (default, will not show these codes; defaut: 404,400)
  -u, --url string              The target URL
  -U, --username string         Username for Basic Auth
  -o, --output string           Output file of type json
      --delay duration          Time each thread waits between requests (e.g. 1500ms)
  -w, --wordlist string         Path to the wordlist
```
# TODO

 - start
 - learn flags
 - plan needed modules
