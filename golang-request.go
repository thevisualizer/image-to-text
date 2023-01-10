package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://ocr-image-to-text4.p.rapidapi.com/image?etype=image"

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"image\"\r\n\r\n\r\n-----011000010111000001101001--\r\n\r\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Add("X-RapidAPI-Key", "YOUR-RAPID-API-KEY")
	req.Header.Add("X-RapidAPI-Host", "ocr-image-to-text4.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
