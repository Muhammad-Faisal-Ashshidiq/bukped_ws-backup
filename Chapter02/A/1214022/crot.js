const form = document.getElementById('ayam');

form.addEventListener('submit', (event) => {
  event.preventDefault(); // prevent default form submission behavior

  const formData = new FormData(form);

  fetch('https://eoecmq3fbx2nhqa.m.pipedream.net', {
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