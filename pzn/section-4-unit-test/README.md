# Section 4 – Unit Test

```md
> Referensi materi: Video 92
```
## Tujuan
Belajar unit test di Go menggunakan package testing.

## Yang Dipelajari
- testing.T
- table driven test
- go test

## Catatan Penting
- File test harus diakhiri `_test.go`
- Function test harus diawali `Test`

## Kesalahan yang Sempat Terjadi
- Test tidak terbaca karena nama function salah
- Lupa import testing

## Insight
Unit test di Go itu simpel, tidak perlu framework tambahan.

## Menjalankan Unit Test
Untuk menjalankan unit test kita bisa menggunakan perintah :
```bash
go test
```
Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah :
```bash
go test -v
```
Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah:
```bash
go test -v -run TestNamaFunction
```
## Menjalankan Semua Unit Test
Jika kita ingin menjalankan semua unit test dari top folder module nya, kita bisa gunakan perintah :
```bash
go test./ ...
```

```md
> Referensi materi: Video 93
```
## Menggagalkan Unit Test

· Menggagalkan unit test menggunakan panic bukanlah hal yang bagus, karena test akan langsung berhenti
· Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
· Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test

## t.Fail() dan t.FailNow()
· Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Lantas apa
bedanya?
· Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir
ketika selesai, maka unit test tersebut dianggap gagal
· FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

## t.Error(args ... ) dan t.Fatal(args ... )
· Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
· Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia
akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
· Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai
selesai
· Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil
FailNow(), sehingga mengakibatkan eksekusi unit test berhenti

```md
> Referensi materi: Video 94
```
## Assertion
· Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
· Apalagi jika result data yang harus di cek itu banyak
· Oleh karena itu, sangat disarankan untuk menggunakan Assertion untuk melakukan pengecekan
· Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh
menambahkan library untuk melakukan assertion ini

## Testify
Salah satu library assertion yang paling populer di Go-Lang adalah Testify
· Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
· https://github.com/stretchr/testify
· Kita bisa menambahkannya di Go module kita :
```bash
go get github.com/stretchr/testify
```

## assert vs require
· Testify menyediakan dua package untuk assertion, yaitu assert dan require
· Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya
eksekusi unit test akan tetap dilanjutkan
· Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil
FailNow(), artinya eksekusi unit test tidak akan dilanjutkan

```md
> Referensi materi: Video 95
```

## Skip Test
· Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
· Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
· Untuk membatalkan unit test kita bisa menggunakan function Skip()

```md
> Referensi materi: Video 96
```

## Before dan After Test
· Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit
test dieksekusi
· Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka
membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode
duplikat jadinya
· Untungnya di Go-Lang terdapat fitur yang bernama testing.M
Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga
bisa kita gunakan untuk melakukan Before dan After di unit test

## testing.M
· Untuk mengatur ekeskusi unit test, kita cukup membuat sebuah function bernama TestMain
dengan parameter testing.M
· Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi
function ini tiap kali akan menjalankan unit test di sebuah package
· Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
· Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function
unit test

```md
> Referensi materi: Video 97
```
## Sub Test
. Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
· Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang
lainnya
· Untuk membuat sub test, kita bisa menggunakan function Run()

## Menjalankan Hanya Sub Test
· Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
go test -run TestNamaFunction
· Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
go test -run TestNamaFunction/NamaSubTest
. Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
go test -run /NamaSubTest

```md
> Referensi materi: Video 98
```

## Table Test
· Sebelumnya kita sudah belajar tentang sub test
· Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
· Dan fitur sub test ini, biasa digunaka oleh programmer Go-Lang untuk membuat test dengan
konsep table test
· Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi
hasil dari unit test
· Lalu slice tersebut kita iterasi menggunakan sub test

```md
> Referensi materi: Video 99
```

# Mock
· Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika
dipanggil, dia akan menghasilkan data yang sudah kita program diawal
· Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari
suatu object yang memang sulit untuk di testing
· Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil
API Call ke third party service. Hal ini sangat sulit untuk di test, karena unit testing kita harus
selalu memanggil third party service, dan belum tentu response nya sesuai dengan apa yang kita
mau
· Pada kasus seperti ini, cocok sekali untuk menggunakan mock object

# Testify Mock
· Untuk membuat mock object, tidak ada fitur bawaan Go-Lang, namun kita bisa menggunakan
library testify yang sebelumnya kita gunakan untuk assertion
· Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin
membuat mock object
· Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan
mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik
. Mari kita buat contoh kasus

# Aplikasi Query Ke Database
· Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke
database
. Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan
ke database
· Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface

# Urutan pembuatan mock
Entity -> Repo -> RepoMock -> Service  

