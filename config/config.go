package config

import ()

type Headers map[string] interface{}

type Settings struct {
	URL          string
	Method       string
	Headers		 string
	Data         interface{}
}

type Response struct {
	StatusCode int
	Headers    Headers
	Body       map[string]interface{}
}
