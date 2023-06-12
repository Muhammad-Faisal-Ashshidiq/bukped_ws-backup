function PushButton(){
    nama=document.getElementById("nama").value;
    email=document.getElementById("email").value;
    phone=document.getElementById("phone").value;
    blood=document.getElementById("blood").value;
    password=document.getElementById("password").value;
    PostSignUp(nama,email,phone,blood,password);
}

function PostSignUp(nama,email,phone,blood,password){
var myHeaders = new Headers();
myHeaders.append("login", "npm1214005");
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
  "nama": nama,
  "email": email,
  "phone": phone,
  "blood": blood,
  "password": password
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("https://eoh66q54brqjgxa.m.pipedream.net", requestOptions)
  .then(response => response.text())
  .then(result => GetResponse(result))
  .catch(error => console.log('error', error));
}

function GetResponse(result){
    document.getElementById("formsignup").innerHTML = result;
}