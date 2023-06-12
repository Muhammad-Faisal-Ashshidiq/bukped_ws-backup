var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

  hasil = ""
txt=""

trtaabel =  `<tr class="hover:bg-gray-50">
<th class="flex gap-3 px-6 py-4 font-normal text-gray-900">
  <div class="text-sm">
    <div class="font-medium text-gray-700">#crot#</div>
  </div>
</th>
</tr>`;


  fetch("https://parallelum.com.br/fipe/api/v1/carros/marcas", requestOptions)
    .then(response => response.text())
    .then(result => tampilkan(result))
    .catch(error => console.log('error', error));


function tampilkan(result){
    console.log(result)
    hasil=JSON.parse(result);
    // crot = trtaabel.replace("#text#", hasil.data[1].nome)
    txt = hasil.forEach(isiintabelnya);
}

function isiintabelnya(value){
    console.log(value)
    txt+= trtaabel.replace("#crot#", value.nome)
    document.getElementById("cakep").innerHTML=txt;
    return txt

}
