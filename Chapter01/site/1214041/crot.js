var requestOptions = {
  method: 'GET',
  redirect: 'follow'
};

hasil="";
txt="";

fetch("https://api.openbrewerydb.org/breweries", requestOptions)
  .then(response => response.json())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

function tampilkan(result){
  console.log(result);
  hasil = result;
  txt = hasil.map(isitabel).join("");
  document.getElementById("konten").innerHTML = txt;
}

tabel= `
<tr class="hover:bg-gray-50">
<th class="flex gap-3 px-6 py-4 font-normal text-gray-900">
  <div class="relative h-10 w-10">
    <span class="absolute right-0 bottom-0 h-2 w-2 rounded-full bg-green-400 ring ring-white"></span>
  </div>
  <div class="text-sm">
    <div class="font-medium text-gray-700" id="name">#TEXT#</div>
  </div>
</th>
</tr>`;

function isitabel(value){
  console.log(value)
  return tabel.replace("#TEXT#", value.name);
}
