package main /*to make the package as an executable program. In Go language,
the package main is a special package, which is used with the programs that are executable*/

import "fmt" //fmt=format

//main function di bawah ini, is the entry point of the executable program.
//whenever we run a Go program, Go will automatically call the main function
//So, theres no need to call the main function explicitly
//Every executable program must contain the single main package (kek di line 1) and the main function
func main() {
	fmt.Println("Hello World!")
}
