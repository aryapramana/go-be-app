package model

import "time"

type BinData struct {
	Args    Args    `json:"args"`
	Data    string  `json:"data"`
	Files   Files   `json:"files"`
	Form    Form    `json:"form"`
	Headers Headers `json:"headers"`
	JSON    JSON    `json:"json"`
	Origin  string  `json:"origin"`
	URL     string  `json:"url"`
}
type Args struct {
}
type Files struct {
}
type Form struct {
}
type Headers struct {
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	ContentLength  string `json:"Content-Length"`
	ContentType    string `json:"Content-Type"`
	Host           string `json:"Host"`
	UserAgent      string `json:"User-Agent"`
	XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
}
type JSON struct {
	Brand     string    `json:"brand"`
	Createdat time.Time `json:"createdat"`
	Discount  int       `json:"discount"`
	Display   string    `json:"display"`
	ID        int       `json:"id"`
	Price     int       `json:"price"`
	Rating    int       `json:"rating"`
	Updatedat time.Time `json:"updatedat"`
}
