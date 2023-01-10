const fs = require('fs');
const request = require('request');

const options = {
    method: 'POST',
    url: 'https://ocr-image-to-text4.p.rapidapi.com/image',
    qs: {
        etype: 'image'
    },
    headers: {
        'content-type': 'multipart/form-data',
        'X-RapidAPI-Key': 'YOUR-RAPID-API-KEY',
        'X-RapidAPI-Host': 'ocr-image-to-text4.p.rapidapi.com',
        useQueryString: true
    },
    formData: {
        image: {
            value: fs.createReadStream('example.png'),
            options: {
                filename: 'example.png',
                contentType: 'application/octet-stream'
            }
        }
    }
};

request(options, function(error, response, body) {
    if (error) throw new Error(error);

    console.log(body);
});
