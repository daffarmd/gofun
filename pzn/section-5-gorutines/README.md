# Parallel Programming adalah 
teknik memecah masalah menjadi bagian-bagian kecil yang dijalankan secara bersamaan, memanfaatkan prosesor multicore. Contohnya menjalankan banyak aplikasi sekaligus, beberapa koki memasak, atau banyak teller melayani nasabah.

## Process dan Threads:
* *Process* adalah eksekusi program yang berdiri sendiri, memakai memori besar, saling terisolasi, dan lebih lambat dijalankan/dihentikan.
* *Thread* adalah bagian dari process, lebih ringan, berbagi memori dalam process yang sama, serta lebih cepat dijalankan/dihentikan.

## Parallel vs Concurrenc:
* *Parallel* menjalankan beberapa pekerjaan secara benar-benar bersamaan dan biasanya membutuhkan banyak thread.
* *Concurrency* menjalankan beberapa pekerjaan secara bergantian dalam satu waktu, dengan jumlah thread lebih sedikit.

## CPU-bound
* Algoritma yang bergantung pada kecepatan CPU (misalnya Machine Learning).
* Lebih cocok menggunakanParallel Programming, kurang efektif dengan concurrency.

## I/O-bound
* Aplikasi yang bergantung pada kecepatan input/output (file, database).
* Lebih efektif menggunakanConcurrency Programming karena waktu menunggu I/O bisa dimanfaatkan untuk tugas lain.


# Goroutine 
adalah unit eksekusi ringan di Go yang dijalankan oleh Go Scheduler di atas thread. Jumlah thread yang digunakan dibatasi oleh GOMAXPROCS (biasanya sesuai jumlah core CPU). Goroutine bukan pengganti thread, melainkan berjalan di atas thread. Keunggulannya, developer tidak perlu mengelola thread secara manual karena semuanya diatur otomatis oleh Go Scheduler.

Dalam Go Scheduler dikenal tiga komponen utama:

* G (Goroutine): tugas yang dijalankan
* M (Machine): thread sistem operasi
* P (Processor): pengatur eksekusi goroutine ke thread


# Membuat Goroutine
* Untuk membuat goroutine di Golang sangatlah sederhana
* Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
* Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
* Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai
* Goroutine tidak cocok jika sebuah function mengembalikan return karena return value nya tidak bisa ditangkap oleh goroutine

# Goroutine Sangat Ringan
* Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
* Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
* Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

# Pengenalan Channel
* Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
* Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
* Saat melakukan pengiriman data ke Channel, goroutine akan ter-block, sampai ada yang menerima data tersebut
* Maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking)
* Channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa bahasa pemrograman lain
* Channel menunggu data yang di kirim oleh goroutine

# Karakteristik Channel
* Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di channel diambil
* Channel hanya bisa menerima satu jenis data
* Channel bisa diambil dari lebih dari satu goroutine
* Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak

# Membuat Channel
```
	channel := make(chan string)
```
* Channel di Go-Lang direpresentasikan dengan tipe data chan
* Untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
* Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan kedalam channel tersebut

# Mengirim dan Menerima Data dari Channel
* Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
* Untuk mengirim data, kita bisa gunakan kode : channel <- data
* Sedangkan untuk menerima data, bisa gunakan kode : data <- channel
* Jika selesai, jangan lupa untuk menutup channel menggunakan function close()
* Kalau kita membuat sebuah channel pastikan ada yang mengirim data dan menerima kalau tidak ada salah satu dari itu maka channel akan error

# Channel Sebagai Parameter
* Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
* Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference). 
* Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut

# Channel In dan Out
* Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
* Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya  dapat digunakan untuk menerima data
* Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out * (menerima data)

# Buffered Channel
* Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
* Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
* Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
* Untungnya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

# Buffer Capacity
* Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
* Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
* Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
* Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data

# Range Channel
* Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
* Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
* Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
* Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
* Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual

# Select Channel
* Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
* Lalu kita ingin mendapatkan data dari semua channel tersebut
* Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
* Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random


