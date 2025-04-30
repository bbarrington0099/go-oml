package gooml

type oml struct {
	keyword string // Word used by the parser to identify the needed literal
	literal string
	ref 	string
}

func newOml(keyword, literal, ref string) *oml {
	return &oml{
		keyword: keyword,
		literal: literal,
		ref:     ref,
	}
}

func (o *oml) getKeyword() (keyword string) {
	keyword =  o.keyword
	return
}

func (o *oml) getLiteral() (literal string) {
	literal = o.literal
	return
}

func (o *oml) getRef() (ref string) {
	ref = o.ref
	return
}