package tests

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "Data"
		fmt.Println("Berhasil !!")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Channel Here!"
}

func TestGiveMeResponse(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Masuk nih"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(4 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "String 1"
		channel <- "String 2"
		channel <- "String 3"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Selesai")

	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // jika ini tidak ada maka deadlock karena akan selalu mengecek data selanjutnya padahal data sudah tidak ada

	}()

	for data := range channel {
		fmt.Println("Menerima Data ", data)
	}

	fmt.Println("End")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	count := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1", data)
			count++
		case data := <-channel2:
			fmt.Println("Channel 2", data)
			count++
		}

		if count == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	count := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1", data)
			count++
		case data := <-channel2:
			fmt.Println("Channel 2", data)
			count++
		default:
			fmt.Println("Menunggu Data!")
		}

		if count == 2 {
			break
		}
	}

}
