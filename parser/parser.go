package parser

type Parser interface {
	Parse(command string) error
}

type SimpleParser struct{}

func NewSimpleParser() SimpleParser {
	return SimpleParser{}
}
