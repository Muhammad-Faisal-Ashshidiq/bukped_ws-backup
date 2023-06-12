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
  myHeaders.append("login", "ratu");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    "first_name": "Ratu",
    "last_name": "sukmakomala",
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

  fetch("https://eohlk0xzf926su9.m.pipedream.net/", requestOptions)
  .then(response => response.text())
  .then(result => GetResponse(result))
  .catch(error => console.log('error', error));
}

function GetResponse(result){
  document.getElementById("formsignup").innerHTML = result;
}