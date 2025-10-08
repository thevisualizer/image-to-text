const qs = require('querystring');
const http = require('https');

const options = {
	method: 'POST',
	hostname: 'image-to-text40.p.rapidapi.com',
	port: null,
	path: '/ocr-by-image',
	headers: {
		'x-rapidapi-key': 'YOUR-RAPID-API-KEY',
		'x-rapidapi-host': 'image-to-text40.p.rapidapi.com',
		'Content-Type': 'application/x-www-form-urlencoded'
	}
};

const req = http.request(options, function (res) {
	const chunks = [];

	res.on('data', function (chunk) {
		chunks.push(chunk);
	});

	res.on('end', function () {
		const body = Buffer.concat(chunks);
		console.log(body.toString());
	});
});

req.write(qs.stringify({}));
req.end();
