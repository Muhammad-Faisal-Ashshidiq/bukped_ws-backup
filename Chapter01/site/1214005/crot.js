var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

hasil =""
txt=""

trtaabel = `<tr class="hover:bg-gray-50">
<th class="flex gap-3 px-6 py-4 font-normal text-gray-900">
  <div class="relative h-10 w-10">
    <span class="absolute right-0 bottom-0 h-2 w-2 rounded-full bg-green-400 ring ring-white"></span>
  </div>
  <div class="text-sm">
    <div class="font-medium text-gray-700" id="nama">#TEXT#</div>
  </div>
</th>
</tr>`;


  fetch("https://pokeapi.co/api/v2/pokemon?limit=100000&offset=0", requestOptions)
    .then(response => response.text())
    .then(result => tampilkan(result))
    .catch(error => console.log('error', error));


function tampilkan(result){
    console.log(result)
    hasil=JSON.parse(result);
    // crot = trtaabel.replace("#text#", hasil.data[1].title)
    txt = hasil.results.forEach(isiintabelnya);
}

function isiintabelnya(value){
    console.log(value)
    txt+= trtaabel.replace("#TEXT#", value.name)
    document.getElementById("judul").innerHTML=txt;
    return txt

}