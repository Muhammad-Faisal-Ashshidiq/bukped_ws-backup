var myHeaders = new Headers();
myHeaders.append("token", "2167bb35-70c9-4009-b5f6-91ce3854841b");
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
  "first name": "",
  "last name": "",
  "email": "",
  "password": "",
  "phone number": ""
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("https://eoyhkfbzqhme8h8.m.pipedream.net", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));

    