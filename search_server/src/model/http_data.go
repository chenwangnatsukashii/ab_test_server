package model

type HttpData struct {
	Url    string `json:"url" form:"url"`
	Data   string `json:"data" form:"data"`
	Judge  string `json:"judge" form:"judge"`
	Number int    `json:"number" form:"number"`
	Tps    int    `json:"tps" form:"tps"`
	Host   string `json:"Host" form:"Host"`
	//ContentLength  string `json:"Content-Length"`
	//ContentType    string `json:"Content-Type"`
	//UserAgent      string `json:"UserAgent"`
	//Accept         string `json:"Accept"`
	//AcceptEncoding string `json:"AcceptEncoding"`
	//Connection     string `json:"Connection"`
}
