# Pass by Value
- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya
- jadi ketika ada perubahan di awal itu datanya akan aman di awal karena datanya di duplikat
```bash
type Address struct {
	Kota      string
	Kecamatan string
	KodePos   int
}

func main() {
	addressKu := Address{"Malang", "Kromengan", 65165}
	addressChange := addressKu

	addressChange.KodePos = 77777

	fmt.Println(addressKu)
	fmt.Println(addressChange)
}
output :
go run main.go
{Malang Kromengan 65165}
{Malang Kromengan 77777}
```
