  const form = document.querySelector('form');

  form.addEventListener('submit', (elemen) => {

    elemen.preventDefault();

    const formulir = new FormData(form);
    
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var object = {};
    formulir.forEach(function(value, key){
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
  
    fetch("https://eoixeju2haq0aoe.m.pipedream.net", requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error));

    document.getElementById('form').reset();
  })

  //upload
  