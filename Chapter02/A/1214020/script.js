var myHeaders = new Headers();
myHeaders.append("Login", "Farhan Kasep");
myHeaders.append("Content-Type", "application/json");

// Pilih/Select Elemen dari Form
const form = document.querySelector('#myForm');

// Menangani Event pengiriman form 
form.addEventListener('submit', (evt) => {
    evt.preventDefault(); // Mencegah Form dikirim melalui HTTP otomatis

    // Create atau membuat objek FormData yang baru dari data pada form 
    const formData = new FormData(form);


    var object = {};
    formData.forEach(function (value, key) {
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

    // Mengirim permintaan POST ke dalam API Endpoint pipedream
    fetch("https://eowvbrqjnfebizu.m.pipedream.net", requestOptions)
        .then(response => response.text())
        .then(result => GetResponse(result))
        .catch(error => console.log('error', error));

    form.reset();

    function GetResponse(result) {
        document.getElementById("formsignup").innerHTML = result;
        document.getElementById("formsignup").style.display = 'none';
        document.getElementById("formsignup").style.display = 'block';
    }
})