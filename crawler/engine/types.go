package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Request []Request
	Items   []interface{}
}

func NilParser(contents []byte) ParserResult {
	return ParserResult{}
}
