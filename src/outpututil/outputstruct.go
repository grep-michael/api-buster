package outpututil

type ResultList []Result

type Result struct {
	Method     string
	Url        string
	StatusCode int
}
