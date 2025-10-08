import requests

url = "https://image-to-text40.p.rapidapi.com/ocr-by-image"

payload = {}
headers = {
	"x-rapidapi-key": "YOUR-RAPID-API-KEY",
	"x-rapidapi-host": "image-to-text40.p.rapidapi.com",
	"Content-Type": "application/x-www-form-urlencoded"
}

response = requests.post(url, data=payload, headers=headers)

print(response.json())
