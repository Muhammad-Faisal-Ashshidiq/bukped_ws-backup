const form = document.getElementById('sign');

form.addEventListener('submit', (event) => {
  event.preventDefault(); // prevent default form submission behavior

  const formData = new FormData(form);

  var myHeaders = new Headers();
  myHeaders.append("token", "fad5dd656f64a5906e2b7eb50407e78e");  

    fetch('https://eorcpkprx5bt8wg.m.pipedream.net',{
    method: 'POST',
    body: formData
  })
    .then(response => response.text())
    .then(data => {
      console.log('Success:', data);
      // do something with the response data
    })
    .catch((error) => {
      console.error('Error:', error);
      // handle any errors that occur during the request
    });
});
