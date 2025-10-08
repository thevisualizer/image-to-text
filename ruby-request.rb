require 'uri'
require 'net/http'

url = URI("https://image-to-text40.p.rapidapi.com/ocr-by-image")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

request = Net::HTTP::Post.new(url)
request["x-rapidapi-key"] = 'YOUR-RAPID-API-KEY'
request["x-rapidapi-host"] = 'image-to-text40.p.rapidapi.com'
request["Content-Type"] = 'application/x-www-form-urlencoded'

response = http.request(request)
puts response.read_body
