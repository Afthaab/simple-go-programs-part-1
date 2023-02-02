// Golang program to illustrate
// the working of the slice components
package main
 
import "fmt"
 
func main() {
	var fruit=[]string{"apple","Banana","Mouse"}
	fmt.Println(fruit)

	//append
	fruit=append(fruit,"strawberry")
	fmt.Println(fruit)

	//sqaure brases
	newfruit:=fruit[1:3]
	fmt.Println(newfruit)
}