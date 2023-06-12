const form = document.getElementById('form');

var myHeaders = new Headers();
myHeaders.append("token", "92533fbd-9cb0-47a1-acd8-cab6efde2030");
myHeaders.append("Content-Type", "application/json");

form.addEventListener('submit', (event) => {
  event.preventDefault(); // prevent default form submission behavior

  const formData = new FormData(form);

  fetch('https://eoilgbmdny2nho1.m.pipedream.net', {
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