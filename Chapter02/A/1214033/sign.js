function PushButton(){
  first_name=document.getElementById("first_name").value;
  last_name=document.getElementById("last_name").value;
  password=document.getElementById("password").value;
  negara=document.getElementById("negara").value;
  kota=document.getElementById("kota").value;
  PostSignUp(first_name,last_name,password,negara,kota);
}

function PostSignUp(first_name,last_name,password,negara,kota){
  var myHeaders = new Headers();
  myHeaders.append("login", "ifahpunya");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    "first_name": "Isnaenti",
    "last_name": "Latifah",
    "password": "12345",
    "negara": "Indonesia",
    "kota": "Bandung"
  });

  var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
  };

  fetch("https://eoexo7x8b7agsha.m.pipedream.net", requestOptions)
  .then(response => response.text())
  .then(result => GetResponse(result))
  .catch(error => console.log('error', error));
}

function GetResponse(result){
  document.getElementById("formsignup").innerHTML = result;
}