var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AMsnp_KW3uPWTf6gN4GDNl7XAoOShdRL2.VK05aaDbN1FeG%2BScGHtOuxENv5s2ABoZZzLpqN%2FUbZs");

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

hasil=""
txt=""

trtaabel=`
<tr class="hover:bg-gray-50">
        <th class="flex gap-3 px-6 py-4 font-normal text-gray-900">
          <div class="text-sm">
            <div class="font-medium text-gray-700">#TEXT#</div>
          </div>
        </th>
      </tr>`;

fetch("https://api.spacexdata.com/v5/launches/latest", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

function tampilkan(result){
  console.log(result);
  hasil=JSON.parse(result);   
  // crot = trtaabel.replace("#TEXT#",hasil[0].text)
  txt=hasil.crew.forEach(isitabelnya);
}

function isitabelnya(value){
  console.log(value)
  txt+= trtaabel.replace("#TEXT#",value.role);
  document.getElementById("karpet").innerHTML=txt;
  return txt
}

