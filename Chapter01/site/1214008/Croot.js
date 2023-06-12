var requestOptions = {
  method: 'GET',
  redirect: 'follow'
};
   
hasil = ""
txt=""

fetch("https://indonesia-public-static-api.vercel.app/api/volcanoes", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));
    
    function tampilkan(result){
      console.log(result);
      hasil=JSON.parse(result);
      trhi=trnyatabel.replace("#TEXT#",hasil[0].text)
      txt=hasil.forEach(isikantabel);
    }
    function isikantabel(value){
      txt= txt+trnyatabel.replace("#TEXT#","Nama : "+value.nama);
      document.getElementById("konten").innerHTML=txt;
    }

    trnyatabel= 
   ` <tr>
       <div class="font-medium text-gray-700" id="nama">
       <th scope="col" class="px-6 py-4 font-medium text-gray-900">#TEXT#</th>
        </div>
     </tr>
    `
    
   
    


   