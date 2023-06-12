const form = document.querySelector('form');

  form.addEventListener('submit', (e) => {

    e.preventDefault();

    const formData = new FormData(form);
    
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("token", "iyas si ganteng kalem");

    var object = {};
    formData.forEach(function(value, key){
        object;
        object[key] = value;
    });

    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: JSON.stringify(object),
      redirect: 'follow',
    };

    console.log(requestOptions);
  
    fetch("https://eo596jk4jc4jncq.m.pipedream.net", requestOptions)
    .then(response => response.text())
    .then(result => GetResponse(result))
    .catch(error => console.log('error', error));

    document.getElementById('form').reset();

    function GetResponse(result){
      document.getElementById("formsignup").innerHTML = result;
      document.getElementById("formsignup").style.display = 'none';
      document.getElementById("formsignup").style.display = 'block';
    }

  })