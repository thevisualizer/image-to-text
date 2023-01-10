import requests

url = "https://ocr-image-to-text4.p.rapidapi.com/image"

querystring = {"etype":"image"}

payload = "-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"image\"\r\n\r\n\r\n-----011000010111000001101001--\r\n\r\n"
headers = {
	"content-type": "multipart/form-data; boundary=---011000010111000001101001",
	"X-RapidAPI-Key": "YOUR-RAPID-API-KEY",
	"X-RapidAPI-Host": "ocr-image-to-text4.p.rapidapi.com"
}

response = requests.request("POST", url, data=payload, headers=headers, params=querystring)

print(response.text)
