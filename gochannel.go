package main
  
import "fmt"
  
func main() {
  
    // Creating a channel
    // Using var keyword
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel)
    fmt.Printf("Type of the channel: %T ", mychannel)
  
    // Creating a channel using make() function
    mychannel1 := make(chan int)
    fmt.Println("\nValue of the channel1: ", mychannel1)
    fmt.Printf("Type of the channel1: %T ", mychannel1)

	// Go program to illustrate send
    // and receive operation

	fmt.Println("start Main method")
    // Creating a channel
    ch := make(chan int,45)
  
    go myfunc(ch)
    ch <- 23
    fmt.Println("End Main method")

    test()
	
}
func myfunc(ch chan int) {
  
    fmt.Println(234 + <-ch)
}
func test()  {
    a:=10;
    fmt.Println(&a,a)
    
}