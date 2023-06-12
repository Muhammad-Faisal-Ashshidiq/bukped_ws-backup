var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var fullNameInput = document.querySelector('input[name="fullname"]');
var usernameInput = document.querySelector('input[name="Username"]');
var emailInput = document.querySelector('input[name="email"]');
var passwordInput = document.querySelector('input[name="password"]');
var confirmPasswordInput = document.querySelector('input[name="confirm_password"]');

var createAccountBtn = document.querySelector("#create-account-btn");
createAccountBtn.addEventListener("click", function () {
  var raw = JSON.stringify({
    Fullname: fullNameInput.value,
    Username: usernameInput.value,
    Email: emailInput.value,
    Password: passwordInput.value,
    Confirm_password: confirmPasswordInput.value,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eoi6v8ok1i29298.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log("error", error));

  function GetResponse(result) {
    document.getElementById("formsignup").innerHTML = result;
    document.getElementById("formsignup").style.display = "none";
    document.getElementById("formsignup").style.display = "block";
  }
});
