const form = document.querySelector('form');

    form.addEventListener('submit', (e) => {
      e.preventDefault();
      const formData = new FormData(form);
      Postdata(formData);
      resetform();
    })

  function Postdata(formData){
    var myHeaders = new Headers();
    myHeaders.append("Login", "hilalfakhrid");
    myHeaders.append("Content-Type", "application/json");

    var object = {};
    formData.forEach(function(value, key){
        object;
        object[key] = value;
    });
    
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: JSON.stringify(object),
      redirect: 'follow'
    };
    
    fetch("https://eori9eiuv0of5i3.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => AmbilResponse(result))
      .catch(error => console.log('error', error));
  }

  function AmbilResponse(result) {
    document.getElementById("formsignup").innerHTML = result;  }

  function resetform(){
    document.getElementById('form').reset();
  }

