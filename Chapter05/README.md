# Deployment Aplikasi Frontend

Kita akan membuat frontend yang sudah kita buat berkomunikasi dengan backend yang sudah dibuat sebelumnya. Contoh :
https://gocroot.herokuapp.com/presensi
Adapun langkah yang akan kita lakukan :
1. Membuat tabel di dalam kontainer dari tampilan yang sudah kita buat
2. Menambahkan jscroot di dalam file html yang kita buat
3. Membuat dan mengisi kerangka jscroot

## Kontainer, Tabel, tag, Content

Pertama pakai frontend yang sudah kemaren kita buat, deploy di github pages. Kita buat di dalam kontainer sebuah tabel untuk mengkonsumsi API.
Contoh frontend : https://github.com/adorableproject/presensi dengan alamat github pages : https://adorableproject.github.io/presensi/ . Hal yang pertama kali kita lakukan adalah membuka file index.html dan menambahkan tag di dalam section head untuk memanggil file js/croot.js.
```html
<script type="module" src="js/croot.js"></script>
```
![image](https://user-images.githubusercontent.com/11188109/224804727-4e5bf910-a563-4d38-85b9-6acb3a24f5df.png)

### Penggunaan Framework JSCroot
Kita akan menggunakan framework frontent jscroot, bisa dilihat di situs : https://jscroot.github.io/

pertama kita edit file index.html kita tambahkan id dengan value namadivisi di salah satu judul tabel
```html
<p id="namadivisi" class="font-semibold text-xl text-coolGray-800">Karyawan</p>
```
kita coba isi croot.js dengan
```js
import { setInner } from "https://jscroot.github.io/element/croot.js";
setInner("namadivisi","Dari croot.js");
```
maka judul berubah, berarti JSCroot sudah berjalan dengan baik

![image](https://user-images.githubusercontent.com/11188109/224890410-27a737dd-ec0d-416d-9614-2a0fd913f090.png)

### Pengambilan Data dari Backend

Kita buat file croot.js di dalam folder js yang berisi.
```js
import { get } from "https://jscroot.github.io/api/croot.js";
import { setInner } from "https://jscroot.github.io/element/croot.js";

let URLPresensi = "https://gocroot.herokuapp.com/presensi";

get(URLPresensi,isiTablePresensi);

function isiTablePresensi(results){
    console.log(results);
}
setInner("namadivisi","Dari croot.js");
```
Kita commit dan push, kemudian tunggu hingga centang hijau pertanda github pages sudah terdeploy dengan baik

![image](https://user-images.githubusercontent.com/11188109/224800920-4d2e520a-e74a-4e2e-b305-3b599a356470.png)

kemudian kita akses url github pages nya. Kita lakukan inspect dan masuk ke tab console terdapat error CORS tampak sebagai berikut.

![image](https://user-images.githubusercontent.com/11188109/224802984-7d9bad9e-1390-4e63-a317-6becf62a670d.png)

Artinya kita perlu mendaftarkan url 'https://adorableproject.github.io' ke dalam config cors.go di repo backend 'https://gocroot.herokuapp.com/presensi'.

![image](https://user-images.githubusercontent.com/11188109/224807482-5ab9ef55-33d3-42d7-8f65-699372aebc77.png)

lalu tambahkan AllowHeaders seperti nama header yang akan diinputkan
```
var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Request-Headers, token",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
```

Simpan, commit dan push ke heroku kemudian kita ujicoba lagi front end kita. Liat console kembali sudah terlihat hasil dari backend

![image](https://user-images.githubusercontent.com/11188109/224811597-f1c34899-ab79-4ce0-b87b-8f4354da97ad.png)

Karena hasil dari backend berupa array dari json object. maka kita ubah kode program tambahkan looping foreach pada croot.js
```js
import { get } from "https://jscroot.github.io/api/croot.js";
import { setInner } from "https://jscroot.github.io/element/croot.js";

let URLPresensi = "https://gocroot.herokuapp.com/presensi";

get(URLPresensi,isiTablePresensi);

function isiTablePresensi(results){
    console.log(results);
    results.forEach(isiRow);
}

function isiRow(value){
    console.log(value)
}



setInner("namadivisi","Dari croot.js");
```
Hasilnya kita dapatkan object yang keluar dari consol.log fungsi isiRow

![image](https://user-images.githubusercontent.com/11188109/224831618-a416f2cf-7902-429d-8c76-651d1f1edca4.png)

### Memasukkan data ke dalam tabel

Sekarang kita masukkan ke dalam tabel. Sebelumnya edit index.html tambahkan id pada tabelnya
```html
<table class="w-full" id="presensi">
```

Kita ambil script TR Tabel HTML yang kemudian kita taruh di file table.js yang di deklarasikan ke dalam variabel presensiTag,presensiClass dan presensiContent.
* presensiTag : berisi tag yang digunakan untuk row table, disini menggunakan tag tr
* presensiClass : nilai value dari class yang ada di presensiTag
* presensiContent : isi didalam presensiTag.

![image](https://user-images.githubusercontent.com/11188109/224797531-a0354504-2050-44a5-9d6b-3e2ff2772ff8.png)

isi file table.js
```js
export let presensiTag="tr";
export let presensiClass="h-18 border-b border-coolGray-100";
export let presensiContent=`
<th class="whitespace-nowrap px-4 bg-white text-left">
  <div class="flex items-center -m-2">
    <div class="w-auto p-2">
      <input class="w-4 h-4 bg-white rounded" type="checkbox">
    </div>
    <div class="w-auto p-2">
      <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ADO</div>
    </div>
    <div class="w-auto p-2">
      <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
      <p class="text-xs font-medium text-coolGray-500">#PHONENUMBER#</p>
    </div>
  </div>
</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#LOKASI#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#KET#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-green-500 text-left">#MASUK#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-green-500 text-left">#PULANG#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#DURASI#</th>
<th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
  <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
  </svg>
</th>
`
```
Dimana #NAMA#, #PHONENUMBER#, #LOKASI#,#KET# dst adalah variabel yang akan kita replace dengan data dari API. File table.js ditaruh di folder template di dalam folder js, sehingga strukturnya tampak sebagai berikut.

![image](https://user-images.githubusercontent.com/11188109/224798778-fbd7e885-48c2-44eb-bd0c-aeabeea71726.png)

kita ubah croot.js lagi menjadi
```js
import { get } from "https://jscroot.github.io/api/croot.js";
import {setInner,addChild } from "https://jscroot.github.io/element/croot.js";
import {getRandomColor,getRandomColorName} from "https://jscroot.github.io/image/croot.js";
import {presensiTag,presensiClass,presensiContent} from "./template/table.js"


let URLPresensi = "https://gocroot.herokuapp.com/presensi";

get(URLPresensi,isiTablePresensi);

function isiTablePresensi(results){
    results.forEach(isiRow);
}

function isiRow(value){
    let content=presensiContent.replace("#NAMA#",value.Biodata.Nama).replace("#PHONENUMBER#",value.Phone_number).replace("#LOKASI#",value.Location).replace("#KET#",value.Checkin).replace("#MASUK#",value.Datetime).replace("#PULANG#",value.Datetime).replace("#DURASI#",value.Datetime).replace("#WARNA#",getRandomColor()).replace(/#WARNALOGO#/g,getRandomColorName());
    addChild("presensi",presensiTag,presensiClass,content);
}


setInner("namadivisi","Dari croot.js");
```
jadilah tabel presensi karyawan sudah terisi

![image](https://user-images.githubusercontent.com/11188109/224889495-d6271680-eba6-4485-8e81-09ac18ab9841.png)

### Membuat boilerplate framework frontend

Sekarang kita rapihkan kode program js kita, sehingga bisa terbaca orang lain dan sesuai dengan sudut pandang framework pada umumnya. Kita ubah dengan struktur folder config,controller,template dengan isi file masing-masing seperti terlihat pada gambar di bawah ini.

![image](https://user-images.githubusercontent.com/11188109/224894184-74bc9192-d635-47e1-bd02-f7dbe16a0d39.png)

file croot.js berisi
```js
import { get } from "https://jscroot.github.io/api/croot.js";
import {setInner } from "https://jscroot.github.io/element/croot.js";
import {isiTablePresensi} from "./controller/table.js";
import {URLPresensi} from "./config/url.js";

get(URLPresensi,isiTablePresensi);

setInner("namadivisi","Dari croot.js");
```
file config/url.js berisi :
```js
export let URLPresensi = "https://gocroot.herokuapp.com/presensi";
```
file controller/table.js berisi :
```js
import { addChild } from "https://jscroot.github.io/element/croot.js";
import {getRandomColor,getRandomColorName} from "https://jscroot.github.io/image/croot.js";
import {presensiTag,presensiClass,presensiContent} from "../template/table.js";

export function isiTablePresensi(results){
    results.forEach(isiRow);
}

function isiRow(value){
    let content=presensiContent.replace("#NAMA#",value.Biodata.Nama).replace("#PHONENUMBER#",value.Phone_number).replace("#LOKASI#",value.Location).replace("#KET#",value.Checkin).replace("#MASUK#",value.Datetime).replace("#PULANG#",value.Datetime).replace("#DURASI#",value.Datetime).replace("#WARNA#",getRandomColor()).replace(/#WARNALOGO#/g,getRandomColorName());
    addChild("presensi",presensiTag,presensiClass,content);
}
```
file template/table.js berisi :
```js
export let presensiTag="tr";
export let presensiClass="h-18 border-b border-coolGray-100";
export let presensiContent=`
<th class="whitespace-nowrap px-4 bg-white text-left">
  <div class="flex items-center -m-2">
    <div class="w-auto p-2">
      <input class="w-4 h-4 bg-white rounded" type="checkbox">
    </div>
    <div class="w-auto p-2">
      <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ADO</div>
    </div>
    <div class="w-auto p-2">
      <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
      <p class="text-xs font-medium text-coolGray-500">#PHONENUMBER#</p>
    </div>
  </div>
</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#LOKASI#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#KET#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-green-500 text-left">#MASUK#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-green-500 text-left">#PULANG#</th>
<th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#DURASI#</th>
<th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
  <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
  </svg>
</th>
`
```

## Kerjakan
* Buatlah frontend yang sudah sesuai dengan aturan boiler plate JSCroot:
  * Terdiri dari 3 folder yaitu template, controller dan config.(30)
  * setting CORS untuk akses dari github pages (10)
  * Sudah bisa diakses dengan github pages (10)
  * menampilkan data dari backend yang dibuat(30) 
  * Pull Request Frontend ditaruh di repo https://github.com/gocroot/app sesuaikan dengan folder topik pekerjaan (20)
* Pull Request ke repo ini dengan Judul 5-KELAS-NPM-NAMA, menyertakan file README.md dalam folder Chapter05/KELAS/NPM/ yang berisi :
  * URL github pages fork dan asli 
  * Skrinsutan dari frontend yang sudah berhasil mengambil data dari backend
