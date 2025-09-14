package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Kelas struct {
	NamaKelas string
}

func (p Person) GetName() string {
	return p.Name
}

func (k Kelas) GetName() string {
	return k.NamaKelas
}

func SayHello(value HasName) {
	fmt.Printf("Hello %s\n", value.GetName())
}

func main() {
	person := Person{Name: "Daffa"}
	kelas := Kelas{NamaKelas: "X IPA 1"}

	SayHello(person)
	SayHello(kelas)
}
