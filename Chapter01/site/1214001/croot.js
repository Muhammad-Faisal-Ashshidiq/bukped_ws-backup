var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3A6SF1Xjgj77f-I04-yM5MzlPDHDiYrKiy.MDhTJt2H%2Fr8He3rihnExBT5S%2BKtGeD0XwWk5r%2Fvb06o");

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

fetch("https://reqres.in/#support-heading", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));