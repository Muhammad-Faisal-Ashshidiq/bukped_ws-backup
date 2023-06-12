# Deployment Aplikasi Backend

Disini kita akan melakukan :
1. Mendaftarkan diri ke github student pack
2. Menyatukan backend menjadi satu repositori
3. Menyatukan frontend menjadi satu repositori
4. Testing deployment boiler plate dari framework iteung
5. Mengatasi error dan kendala masalah umum yang terjadi pada deployment

## Github Student Pack

Student pack berguna untuk mendapatkan beberapa benefit gratis dari github. Untuk mendapatkan student pack. kita mendaftarkan diri dahulu ke alamat https://education.github.com/pack

![image](https://user-images.githubusercontent.com/11188109/223223891-b76d4099-127e-4d2a-8ff0-6bd79c85951b.png)

Disana terlihat beberapa Experiences yang bisa kita coba secara gratis atau free trial. untuk mendaftarkan diri klik Explore Student Programs. Kemudian pilih Student.

![image](https://user-images.githubusercontent.com/11188109/223224571-51a8dc42-8c7b-4a94-b2b5-9fbe07549f73.png)

Kemudian pilih email kampus ULBI, jika belum memiliki akun email ULBI, bisa meminta ke TIK di ruang Audit Lt.1. Selanjutnya adalah melengkapi keterangan mahasiswa dengan melakukan upload KTM, yang bisa diminta ke Iteung.

![image](https://user-images.githubusercontent.com/11188109/223224816-c6e45175-309b-4fd7-83b8-fb7133e5269b.png)

Proses verifikasi akan dilakukan, dan biasanya selesai 15 menit sudah ada pemberitahuan melalui email.

## Heroku Student offer

Heroku merupakan layanan cloud yang berguna untuk deployment atau pemasangan aplikasi berbasis web yang kita buat. Cara pemasangan aplikasi di Heroku cukup mudah, tinggal hubungkan saja dengan repositori github yang sudah ada.Untuk menikmati layanan heroku secara gratis hubungkan dengan akun github anda yang sudah diaktifkan student pack pada link https://www.heroku.com/github-students

![image](https://user-images.githubusercontent.com/11188109/223225509-4ad4da33-cb20-49fa-b49d-5cf1a5f18782.png)

Kemudian klik Get student offer. Gunakan email kampus anda, dan lanjutkan untuk mendaftarkan diri. Jika dibutuhkan verifikasi pembayaran dengan kartu kredit, bisa menggunakan Virtual Credit Card atau Virtual Debit Card atau Debit Card kartu dari Bank Digital yang bisa didapatkan dengan install aplikasi di handphone. Opsional sebagai pelengkap keamanan akun. Install juga aplikasi Google Authenticator https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=en&pli=1

![image](https://user-images.githubusercontent.com/11188109/223226670-c7afe93a-5bdc-4a74-be12-94945dbc807f.png)

## Dashboard Heroku dan Heroku cli

Setelah login, masuk ke laman https://dashboard.heroku.com/apps. Maka akan muncul list aplikasi yang sudah kita buat. Klik New dan pilih Create new app untuk melakukan deployment aplikasi baru kita.

![image](https://user-images.githubusercontent.com/11188109/223227248-d1e1019f-04ff-4f5e-942b-eb65549302ed.png)

Masukkan nama aplikasi kita, dan pilih lokasi server kita apakah amerika atau eropa. Kemudian klik Create app.

![image](https://user-images.githubusercontent.com/11188109/223227515-c25abd89-5b08-483a-bc47-432e5c4b0b28.png)

Lakukan instalasi Heroku CLI, untuk menghubungkan komputer kita dengan server heroku. Link Instalasi https://devcenter.heroku.com/articles/heroku-cli

![image](https://user-images.githubusercontent.com/11188109/223227952-1e166ee3-78db-4860-8956-6b643a3b57eb.png)

Pilih 64-bit installer

![image](https://user-images.githubusercontent.com/11188109/223228151-831c4735-be1c-413d-ad1c-3ab4583995fa.png)

Lanjutkan sampai selesai langkah instalasinya

![image](https://user-images.githubusercontent.com/11188109/223228251-84e36310-2fac-47b6-aa8d-9d3a6789f7f9.png)

Buka git bash dan ketik heroku login, maka akan muncul button login heroku di browser, klik saja login. Kemudian di gitbash akan ada tulisan Logged in as .....

![image](https://user-images.githubusercontent.com/11188109/223228723-ee5b3b88-69ba-4798-bbde-3abe0fe99265.png)

Login sudah berhasil kita bisa close git bash

![image](https://user-images.githubusercontent.com/11188109/223228974-93862ebe-af41-4612-81ff-3f71b5cf0a79.png)

Kita bisa juga melakukan add SSH key ke heroku kita dengan klik gambar profile di pojok kanan atas > Account Settings > SSH Keys

![image](https://user-images.githubusercontent.com/11188109/223316679-790d7423-0dc1-4e5c-bdae-1a2afb668228.png)


## Deployment Boiler Plate

Disini kita akan mencoba testing deployment ke Heroku. Aplikasi web yang akan dilakukan deploymeny adalah Boiler Plate iTeung yang berada di repo https://github.com/aiteung/iteung

![image](https://user-images.githubusercontent.com/11188109/223230699-f000db95-e50e-4db6-90a0-d257a5ad2494.png)

Kita lakukan fork ke repositori kita, kita beri nama sesuai dengan nama aplikasi di heroku. Kemudian, lakukan clone repo ke komputer kita.

![image](https://user-images.githubusercontent.com/11188109/223230979-cb633354-0db8-46db-a073-43f33c23cc28.png)

Setelah di clone menggunakan git bash kemudian masuk ke direktori repo di PC kita. Lakukan add remote heroku sesuai nama aplikasi yang sudah kita buat di heroku dengan perintah

```sh
heroku git:remote -a gocroot
```

dimana gocroot adalah nama aplikasi kita di heroku.

![image](https://user-images.githubusercontent.com/11188109/223231548-521e42b1-32f9-4559-85d9-a36b98129ca3.png)

Untuk melakukan deployment kita cukup mengetikkan perintah 

```sh
git push heroku main
```

Tentunya disini kita akan menemui error pertama dari Heroku. Tenang saja, jangan panik error merupakan bagian dari pembelajaran tinggal kita baca dan ikuti saja log error yang sudah menjadi petunjuk itu.

![image](https://user-images.githubusercontent.com/11188109/223232196-1f4d36e0-1d3b-4e4b-86bd-260cb2767e27.png)

Terlihat kita belum mendefinisikan bahasa yang digunakan untuk aplikasi ini, untuk itu kita akan membaca cara mendefinisikan bahasa bisa dilihat di menu [Heroku Dev Center](https://devcenter.heroku.com/articles/getting-started-with-go#define-a-procfile). Diharuskan membuat Procfile yang berisi

```sh
main: bin/gocroot
```

dimana gocroot adalah nama repo kita.

![image](https://user-images.githubusercontent.com/11188109/223233562-f2328906-61a4-440f-aabb-3aa8519f0c81.png)

Setelah kita buat procfile kita add commit dan push

![image](https://user-images.githubusercontent.com/11188109/223234333-b4df95f1-2ec4-48d2-85c1-4204cd7ce3c1.png)


Terlihat masih error, ternyata kita belum menentukan builpacks, kunjungi https://devcenter.heroku.com/articles/buildpacks untuk melakukan pemilihan buildpacks. Tampak dibagian web tidak membantu. Dan tersadar kita belum melakukan go mod init, karena terlihat belum ada file gomod dan go.sum di folder repo. Kita lakukan dulu init dan berhasil.

![image](https://user-images.githubusercontent.com/11188109/223236780-0a3079ca-0213-46a1-ab54-7b214cb8ac27.png)

Muncul error berikutnya. Tenang saja jangan panik. Karena error kita jadi belajar. terlihat tiba tiba terdeteksi package iteung, ini berarti kode program di dalam file masih import nama package yang lama. Tinggal kita ganti semua importnya dengan nama package yang di deklarasikan pada saat kita go mod init tadi.

![image](https://user-images.githubusercontent.com/11188109/223238838-2f2ace7f-dbda-4dfa-9cc7-0ce03b990d2a.png)

Kita buka VS Code, kita buka terminal pada bagianPROBLEMS akan mengeluarkan error yang serupa. Selesaikan satu persatu di masing-masing file yang muncul problems.

![image](https://user-images.githubusercontent.com/11188109/223239598-40e6ad6a-8fa4-4ac4-89e1-a149094ed25b.png)

Masalah selesai dengan mengganti import di ketiga file yaitu main.go, controller.go dan url.go. Sekarang kita tinggal add commit dan push ke github dan heroku kembali.

![image](https://user-images.githubusercontent.com/11188109/223240939-18fa83ac-006b-41e8-9477-7cc0692c8a98.png)

Terlihat damai dan tentram akhirnya kita berhasil melakukan deployment ke heroku

![image](https://user-images.githubusercontent.com/11188109/223242344-7e10de9a-7a2e-4a8c-a470-70cb5d90d12d.png)

Oke kita coba buka url aplikasi kita, dicontohkan di heroku CLI disini https://gocroot.herokuapp.com/ kita coba buka apakah sesuai dengan yang diharapkan.

![image](https://user-images.githubusercontent.com/11188109/223242649-e632ffc6-ae38-4ace-960d-fb0e0a39f250.png)

Dan kita menemukan error selanjutnya, alhamdulillah. kita coba jalankan heroku logs --tail pada git bash

![image](https://user-images.githubusercontent.com/11188109/223242915-0a1de7d6-8626-41d9-9f7e-dff92197ec5a.png)

Katanya Code H14 No Web processes running, berdasarkan hasil penelusuran https://stackoverflow.com/questions/41804507/h14-error-in-heroku-no-web-processes-running hal ini terjadi karena kita belum melakukan set web dynos di aplikasi kita pada dashboard heroku. Oleh karena itu kita kunjungi kembali dashboard heroku kita.

![image](https://user-images.githubusercontent.com/11188109/223243600-1cec1175-5a39-43d0-be73-e9658876d133.png)

Pada bagian Menu Resources, dan pada bagian main bin/gocroot kita klik tombol tanda pensil, kita geser tombol bulan dan klik tombol confirm

![image](https://user-images.githubusercontent.com/11188109/223243865-ce1debee-020f-4cbf-af37-6fbcc75daaa3.png)

Kita cek web ternyata masih error, kita cek heroku logs ternyata masih error juga. Coba iseng kita hapus Procfile, kemudian kita push kembali. Dan di bagian dashboard, aktifkan kembali Reourcesnya. 

![image](https://user-images.githubusercontent.com/11188109/223246111-4ef01e7c-ce34-4075-8ddd-4017d65c9a57.png)

Dan akhirnya web sudah bisa diakses.

![image](https://user-images.githubusercontent.com/11188109/223246386-1b9fb839-74c2-4826-88f5-12b5ae2ed379.png)

## Iteung Boiler Plate

### Aktifasi Prefork

![image](https://user-images.githubusercontent.com/11188109/223247620-782c1571-d0e8-4f2a-abd4-89a52f457d69.png)

Sekarang aplikasi sudah berjalan dengan baik. Akan tetapi ada beberapa hal yang harus di konfigurasi kembali. Yang paling pertama untuk dilakukan konfigurasi dari boiler plate iteung adalah melakukan aktifasi Prefork. Kita bisa lihat di heroku logs, terlihat Prefork statusnya masih Disabled. Kita buka file main.go tambahkan config.Iteung pada fungsi fiber.New()

```go
func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
```

Kemudian kita buat file config.go di dalam folder config yang berisi :

```go
package config

import "github.com/gofiber/fiber/v2"

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Iteung",
	AppName:       "Message Router",
}
```

Simpan add, commit push kembali ke repo dan heroku. Terlihat di heroku logs, prefork sudah aktif

![image](https://user-images.githubusercontent.com/11188109/223252919-aacc55b5-897d-4b38-a55c-3a54827b6dff.png)

### Free MariaDB JawsDB

Kita akan menambahkan add on MariaDB dari heroku. Pertama dari dashboard kita klik Find more add-ons dan pilih JawsDB Maria https://elements.heroku.com/addons/jawsdb-maria

![image](https://user-images.githubusercontent.com/11188109/223276199-009e12ae-1a53-43ce-b78f-d4fd706182ff.png)

Klik Install JawsDB Maria

![image](https://user-images.githubusercontent.com/11188109/223276346-7549ae48-4b9b-41a0-bbaf-c8a4e59ac092.png)

Pilih nama aplikasi kita di heroku dan klik Submit Order Form

![image](https://user-images.githubusercontent.com/11188109/223276527-235a3ead-76e8-499b-b0ec-ddeb4dd57423.png)

JawsDB akan terinstall, untuk memakainya kita hanya butuh Connection String yang bisa di dapatkan dengan klik JawsDB Maria di bagian Add-ons

![image](https://user-images.githubusercontent.com/11188109/223276910-83bf42d4-cfc7-4d82-9807-ebd18e0f0013.png)

![image](https://user-images.githubusercontent.com/11188109/223277457-c8e6f6a5-09b6-4546-9376-030280a453ae.png)


### Edit Environtment Variabel

Jika pada windows kita melakukan edit di environtment variabel. Begitu juga heroku memiliki fitur tersebut yang bernama config vars. Tinggal kita masuk ke Dashboard bagian Settings > Config Vars > Reveal Config Vars.

![image](https://user-images.githubusercontent.com/11188109/223269495-94f73eb1-7c3b-44bc-b2e4-380e0847ae1b.png)

Kita masukkan beberapa environment yang didapat pada folder config dari boiler plate iteung seperti :

file db.go
```env
ITEUNGBEV1
MONGOSTRING
MARIASTRINGAKADEMIK
```
file token.go
```env
PUBLICKEY
PRIVATEKEY
```
file cors.go
```env
INTERNALHOST
PORT
```
file api.go
```env
URLAPIWABUTTON
```

Untuk saat ini kita tambahkan MARIASTRINGAKADEMIK dan MONGOSTRING terlebih dahulu contoh :
```env
MARIASTRINGAKADEMIK=username:password@tcp(ao9moanwus0rjiex.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/
MONGOSTRING=mongodb+srv://username:password@cluster0.wghp85v.mongodb.net/test
```

![image](https://user-images.githubusercontent.com/11188109/223279062-9a65fb96-50f4-4639-b1b9-b035823445cf.png)

Edit dan tambahkan konfigurasi nama database yang dipakai pada file db.go folder config

```go
var DBUlbimariainfo = atdb.DBInfo{
	DBString: MariaStringAkademik,
	DBName:   "xia3fhuwzm5wo0zo",
}

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "iteung",
}

var Ulbimariaconn = atdb.MariaConnect(DBUlbimariainfo)

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)
```
![image](https://user-images.githubusercontent.com/11188109/223281074-53d90100-0c52-433c-a676-e443e8014e2d.png)

Jika terdapat error seperti diatas maka modif Connection string menjadi seperti berikut :
```env
MARIASTRINGAKADEMIK=username:password@tcp(ao9moanwus0rjiex.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/
```

### Memanggil package golang

Dari package yang sudah dibuat pada chapter sebelumnya kita akan coba panggil dari aplikasi boiler plate iteung yang sudah di deploy di heroku. Sebagai contoh kita akan mencoba memanggil package musik di https://github.com/aiteung/musik hal yang pertama kali adalah kita membuka file url.go di folder url kita tambahkan baris di dalam func Web sebuah baris kode yang memanggil controller.Homepage yaitu fungsi yang akan kita buat :

```go
page.Get("/", controller.Homepage) //ujicoba panggil package musik
```
pada terminal ketikkan perintah go get package yang akan digunakan
```sh
go get github.com/aiteung/musik
```

kemudian kita buka file controller.go di folder controller kita tambahkan fungsi Homepage
```go
func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
```
Kita buka alamat aplikasi dan terlihat ip address dari host heroku

![image](https://user-images.githubusercontent.com/11188109/223288688-d046ad76-d7fd-4f1e-95d5-c0b559d51335.png)

### Mongodb.com Network Access

Agar aplikasi di heroku bisa mengakses mongodb yang sudah dibuat di mongodb.com maka kita harus memasukkan IP adress host heroku yang didapat pada langkah membuat controller homepage ke mongodb.com

![image](https://user-images.githubusercontent.com/11188109/223288962-3917c06a-8332-41cd-b8ce-634180cee9b3.png)

Setelah itu, silahkan masukkan data dummy pada collection dengan mongo compass yang diakses oleh package yang dikembangkan sebanyak-banyaknya

### Deployment Package

Pada chapter sebelumnya. Package yang sudah dibuat bisa kita panggil di controller. Cukup dengan 3 tahap, yaitu :
1. Definisikan alamat URL untuk akses package beserta nama fungsi di controller yang akan kita buat pada file url.go
   ```go
   page.Get("/presensi", controller.GetPresensiBulanIni)
   ```
2. Go get package yang akan digunakan melalui terminal di vscode
   ```sh
   go get github.com/aiteung/presensi
   ```
3. Buat fungsi di file controller.go
   ```go
   func GetPresensiBulanIni(c *fiber.Ctx) error {
        ps := presensi.GetPresensiCurrentMonth(config.Ulbimongoconn)
        return c.JSON(ps)
   }
   ```
Jangan lupa selalu menjalankan perintah 
```sh
go mod tidy
```
Kemudian add, commit push ke repo dan heroku. Kemudian kita akses melaui url yang kita deklarasikan di file url.go. Untuk method GET bisa menggunakan browser saja, untuk POST menggunakan POSTMAN

![image](https://user-images.githubusercontent.com/11188109/223319800-d4743c4d-27af-4636-9d07-7800555e4007.png)

Berhasil, sampai sini kita sudah melakukan deployment backend. Jika terjadi error maka klik ALLOW ACCESS FROM ANYWHERE > Confirm pada menu Network Access > ADD IP ADDRESS mongodb.com

![image](https://user-images.githubusercontent.com/11188109/223330625-28afbfc3-af79-46dc-a1b1-a502181b79eb.png)

![image](https://user-images.githubusercontent.com/11188109/223330948-06b8493d-a9b7-4f19-86ba-06c35d6aaa96.png)



## Kerjakan

* Pull Request Semua fungsi(di folder module) dan type(di folder model) di pull request ke repo https://github.com/gocroot/kampus . Perwakilan kelas melakukan tagging dan publish package kampus
* Pull REquest Frontend ditaruh di repo https://github.com/gocroot/app sesuaikan dengan folder topik pekerjaan
* Mencoba deploy boiler plate https://github.com/aiteung/iteung ke heroku, sudah diisi dengan pemanggilan package dari repo https://github.com/gocroot/kampus
* Github Pages Sudah UP untuk frontend, contoh : https://gocroot.github.io/app/bimbingan/menusaya
* Pull Request dengan Judul 4-KELAS-NPM-NAMA, menyertakan file README.md dalam folder Chapter04/KELAS/NPM/ yang berisi :
  * Daftar URL Heroku app yang sudah memanggil function di package kampus. [contoh](https://gocroot.herokuapp.com/presensi) beserta skinsutannya (nilai 10/URL)
  * Skrinsutan isi mongodb dari mongocompass yang sudah berisi data dummy kasus (nilai 10)
  * URL Package yang sudah publish di pkg.go.dev dari https://github.com/gocroot/kampus beserta skinsutannya (nilai 10)
  * Daftar URL github pages dari setiap studi kasus pada repo https://github.com/gocroot/app beserta skinsutannya contoh. gocroot.github.io/app/bimbingan/menusaya (nilai 10/URL)
