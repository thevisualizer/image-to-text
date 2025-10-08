package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://image-to-text40.p.rapidapi.com/ocr-by-image"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("x-rapidapi-key", "YOUR-RAPID-API-KEY")
	req.Header.Add("x-rapidapi-host", "image-to-text40.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
