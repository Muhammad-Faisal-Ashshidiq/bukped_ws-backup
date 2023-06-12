function PostSignUp(namalengkap, alamatemail, nohp, username, password) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "GiparAnakKuat");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    "namalengkap": namalengkap,
    "alamatemail": alamatemail,
    "nohp": nohp,
    "username": username,
    "password": password
  });

  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
  };

  fetch("https://eo822xg7dl2nv3f.m.pipedream.net", requestOptions)
    .then(response => response.text())
    .then(result => {
      GetResponse(result);
      document.getElementById("formsignup").style.display = 'none';
      document.getElementById("formsuccess").style.display = 'block';
    })
    .catch(error => console.log('error', error));
}

function PushButton() {
  var namalengkap = document.getElementById("namalengkap").value;
  var alamatemail = document.getElementById("alamatemail").value;
  var nohp = document.getElementById("nohp").value;
  var username = document.getElementById("username").value;
  var password = document.getElementById("password").value;
  PostSignUp(namalengkap, alamatemail, nohp, username, password);
}

function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}
