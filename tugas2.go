package main
import "fmt"
func main(){
	var d float32
	
	fmt.Print("Masukan Mata Uang Dolar = ")
	fmt.Scanf("%f", &d)
	
	var r float32
	
	r= d*13510
	fmt.Println("Hasil Konfersi Mata Uang Rupiah = Rp.", r)
}

