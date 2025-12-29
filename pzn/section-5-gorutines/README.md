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

