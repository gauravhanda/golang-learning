package main

import "fmt"

type Encoding interface {
	Encode() string
	Decode() string
}

type URLEncode struct {
	url string
}

func (u URLEncode) String() string {
	return "URLEncoded String from stringer interface"
}

func (b *URLEncode) Encode() string {
	return fmt.Sprintf("URLEncode encode \" %s \"", b.url)
}

func (b *URLEncode) Decode() string {
	return fmt.Sprintf("URLEncode decode \" %s \"", b.url)
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
