package main
import "fmt"

func main(){
	for i:=1; i<=100; i++{
	if i%3==0{
		fmt.Println("lipat 3")
	
	}else if i%5==0{
		fmt.Println("lipat 5")
	}else{
		fmt.Println(i)
	}
	}
}