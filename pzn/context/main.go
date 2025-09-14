package main

import (
	"fmt"
	"time"
)

// worker adalah goroutine yang akan menerima pekerjaan dari channel
func worker(id int, jobs <-chan int, results chan<- string) {
	fmt.Printf("Pekerja %d: Siap menerima tugas.\n", id)
	for j := range jobs { // Menerima pekerjaan dari channel 'jobs'
		fmt.Printf("Pekerja %d: Mengerjakan tugas %d...\n", id, j)
		time.Sleep(500 * time.Millisecond)                                         // Simulasikan pekerjaan
		results <- fmt.Sprintf("Pekerja %d: Selesai mengerjakan tugas %d.", id, j) // Mengirim hasil ke channel 'results'
	}
	fmt.Printf("Pekerja %d: Selesai bekerja.\n", id)
}

func main() {
	// Membuat channel untuk mengirim tugas (jobs) dan menerima hasil (results)
	jobs := make(chan int, 10)   // Channel tanpa buffer untuk 10 tugas
	results := make(chan string) // Channel tanpa buffer untuk hasil
	fmt.Println("[DEBUGING] ini results jobs : ", jobs)
	fmt.Println("[DEBUGING] ini results chan : ", results)

	// Meluncurkan 3 goroutine "worker"
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results) // `go` keyword meluncurkan goroutine baru
	}

	// Mengirim 5 tugas ke channel 'jobs'
	for j := 1; j <= 5; j++ {
		jobs <- j // Mengirim nilai 'j' ke channel 'jobs'
		fmt.Printf("Main: Mengirim tugas %d.\n", j)
	}
	close(jobs) // Menutup channel 'jobs' setelah semua tugas dikirim
	fmt.Println("Main: Semua tugas telah dikirim, channel jobs ditutup.")

	// Menerima hasil dari channel 'results'
	// Kita harapkan 5 hasil, karena 5 tugas dikirim
	for a := 1; a <= 5; a++ {
		res := <-results // Menerima nilai dari channel 'results'
		fmt.Println("Main: Menerima hasil:", res)
	}

	fmt.Println("Main: Semua hasil telah diterima.")
	// Di sini, program utama akan menunggu semua goroutine worker selesai jika kita tidak menutup channel jobs dan memastikan semua hasil diterima.
	// Karena kita menutup jobs dan menunggu results, program akan tahu kapan harus selesai.
}
