package main

import "fmt"

type Blacklist func(string) bool

func RegistrasiUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu masuk dalam daftar blacklist", name)
	} else {
		fmt.Println("Selamat Datang Calon Investor", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "budiono siregar"
	}

	RegistrasiUser("mamat", blacklist)
	RegistrasiUser("budiono siregar", blacklist)

}
