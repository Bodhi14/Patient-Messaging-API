package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func sendSMS(apikey string, numbers string, sender string, message string) {
	resp, err := http.Get("https://api.textlocal.in/send/?")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(resp)
	url := "apikey" + "numbers" + "string" + "message"
	data := EncodeParams(url)
	data = EncodeString(data)

}

func EncodeParams(s string) string {
	return url.QueryEscape(s)
}

func EncodeString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func main() {

}
