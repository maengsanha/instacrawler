package crawler

// Parser is a struct type that performs parsing from Instagram page source.
type Parser struct {
	pageSource string
}

// NewParser is constructor of Parser.
func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) parseInstaPageSource(url string) error {
	return nil
}
