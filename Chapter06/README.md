# Debug Aplikasi

## Pre Test on Class (Waktu 30 Menit)
* Ubahlah frontend form dari Chapter 2 menggunakan library jscroot dan sesuai dengan aturan boiler plate JSCroot:
  * Folder js, terdiri dari 3 folder yaitu template, controller dan config (30)
  * URL POST menggunakan backend sendiri, bukan menggunakan pipedream (40)
  * menampilkan respon dari backend yang dibuat setelah klik submit (30) 
* Presentasikan dan langsung mendapat penilaian di kelas. Setiap keterlambatan 15 menit akan ada diskon nilai sebanyak 5 poin.
* Gunakan Inspect Element untuk kepentingan debug, gunakan hard refresh(shift+F5) setiap ada perubahan kode javascript.

## Penggunaan Inspect Element

Pada sesi ini kita akan mencoba melakukan debugging dari aplikasi yang sudah kita buat. Alat yang dibutuhkan cukup browser google chrome. Kita akan menggunakan Inspect Element. Sebagai contoh kita buka https://adorableproject.github.io/presensi/ klik kanan masuk ke menu Inspect

![image](https://user-images.githubusercontent.com/11188109/226216392-5f69a5ab-0db1-4fa0-82a9-37b9540805a6.png)

Setelah itu pilih tab Network > Fetch/XHR

![image](https://user-images.githubusercontent.com/11188109/226216442-880d8ec3-0eab-4b6f-a03c-4fc6a6ad022a.png)

Masih terlihat kosong, sekarang kita lakukan hard refresh(Shift+F5) atauh refresh(F5) maka akan muncul sebuah data proses pengambilang yang berjalan ke sisi backend

![image](https://user-images.githubusercontent.com/11188109/226216512-598aacd1-3935-4c15-8291-5e98d268f29c.png)

Kita klik dan buka Name dari daftar tersebut. Akan terbuka 

![image](https://user-images.githubusercontent.com/11188109/226216574-f83858e4-2612-49b4-996d-b005a2fb16a1.png)

Terdapat 5 Buah tab disana yaitu :
* Headers
* Preview
* Response
* Initiator
* Timing

Yang kemudian akan kita jelaskan satu per satu maksud dari setiap tab tersebut.

### Headers

![image](https://user-images.githubusercontent.com/11188109/226216819-9e0b2990-fdbd-4b87-a4dc-6cba5dbfcc4c.png)

Pada tab Headers terdapat 3 section yaitu :
* General : Berisi URL Backend, Method Untuk ke Backend, Satus Code dari Backend, Alamat IP Address dari Backend dan Referrer Policy.
* Response Headers : Header Respon dari Backend
* Request Headers : Header Request dari Frontend

### Preview

![image](https://user-images.githubusercontent.com/11188109/226216827-129f0328-061a-4c1c-aa55-b94ba8015ccd.png)

Berisi data yang dikirimkan dari backend yang sudah berbentuk rapih

### Response

![image](https://user-images.githubusercontent.com/11188109/226216882-6ce048e6-726d-4a67-a3bb-3acf907abbab.png)

Berisi raw data yang dikirimkan dari backend

### Initiator

![image](https://user-images.githubusercontent.com/11188109/226216923-5f611ce7-62d3-49b8-9300-9f11e11e3add.png)

Berisi nama fungsi yang memanggil aksi pemanggilan ke backend, beserta runutan posisi file dari awal hingga pemanggilan.

### Duration

![image](https://user-images.githubusercontent.com/11188109/226216967-75228a6a-0838-4b5a-8fb4-dd256e364482.png)

Berisi perhitungan waktu dari setiap aksi request dari frontend ke backend




