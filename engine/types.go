package engine

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

func NilParser([]byte)ParseResult  {
	return ParseResult{}
}

type Item struct {
	Id string
	Url string
	Type string
	Payload interface{}
}