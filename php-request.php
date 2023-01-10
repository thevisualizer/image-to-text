<?php

$curl = curl_init();

curl_setopt_array($curl, [
	CURLOPT_URL => "https://ocr-image-to-text4.p.rapidapi.com/image?etype=image",
	CURLOPT_RETURNTRANSFER => true,
	CURLOPT_FOLLOWLOCATION => true,
	CURLOPT_ENCODING => "",
	CURLOPT_MAXREDIRS => 10,
	CURLOPT_TIMEOUT => 30,
	CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
	CURLOPT_CUSTOMREQUEST => "POST",
	CURLOPT_POSTFIELDS => "-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"image\"\r\n\r\n\r\n-----011000010111000001101001--\r\n\r\n",
	CURLOPT_HTTPHEADER => [
		"X-RapidAPI-Host: ocr-image-to-text4.p.rapidapi.com",
		"X-RapidAPI-Key: YOUR-RAPID-API-KEY",
		"content-type: multipart/form-data; boundary=---011000010111000001101001"
	],
]);

$response = curl_exec($curl);
$err = curl_error($curl);

curl_close($curl);

if ($err) {
	echo "cURL Error #:" . $err;
} else {
	echo $response;
}
