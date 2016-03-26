package main
import "fmt"

func main(){
	i:=1
	for  i<=100{
	if i>100{
		fmt.Println("Selesai")
	}else {
		switch i{
		case 3:
			fmt.Println("lipat 3")
		case 5:
			fmt.Println("lipat 5")
			default:
			fmt.Println("salah")
		}
	}
	i=i+3
	}
}