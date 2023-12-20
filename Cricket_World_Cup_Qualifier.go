package main
import "fmt"

var input int

func main(){
	// your code goes here
	fmt.Scanf("%d", &input)
	if(input >= 12){
	    fmt.Println("Yes")
	}else {
	    fmt.Println("No")
	}
}
