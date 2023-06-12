var myHeaders = new Headers();
myHeaders.append("Login", "farhan");
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
  "name": "riziq",
  "email": "riziq@gmail.com",
  "nohp": "081372293289",
  "password": "12345",
  "password_confirmation": "12345"
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("https://eoeff9kcv5svdgd.m.pipedream.net", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));