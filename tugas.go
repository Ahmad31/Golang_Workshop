package main
import "fmt"
func main(){
	var f float32
	
	fmt.Print("Masukan Nilai Fahrenheit = ")
	fmt.Scanf("%f", &f)
	
	var c float32
	
	c= (f-32)*5/9
	fmt.Println("Hasil Konfersi ke-Celcius = ", c)
}
