package protol

type Message struct {
	Header Header	`header`
	Content Content `content`
}

type Header struct {
	Type string `type`
	Token string `token`
}

type Content struct {
	Method string `method`
	Params map[string]interface{} `params`
}