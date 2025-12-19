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