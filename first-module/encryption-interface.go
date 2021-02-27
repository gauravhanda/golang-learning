package main

import "fmt"

type Encoding interface {
	Encode() string
	Decode() string
}

type MyBase64 struct {
	message string
}

func (b *MyBase64) Encode() string {
	return fmt.Sprintf("base64 encode \" %s \"", b.message)
}

func (b *MyBase64) Decode() string {
	return fmt.Sprintf("base64 decode \" %s \"", b.message)
}
