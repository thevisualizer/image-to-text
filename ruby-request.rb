require 'uri'
require 'net/http'
require 'openssl'

url = URI("https://ocr-image-to-text4.p.rapidapi.com/image?etype=image")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE

request = Net::HTTP::Post.new(url)
request["content-type"] = 'multipart/form-data; boundary=---011000010111000001101001'
request["X-RapidAPI-Key"] = 'YOUR-RAPID-API-KEY'
request["X-RapidAPI-Host"] = 'ocr-image-to-text4.p.rapidapi.com'
request.body = "-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"image\"\r\n\r\n\r\n-----011000010111000001101001--\r\n\r\n"

response = http.request(request)
puts response.read_body
