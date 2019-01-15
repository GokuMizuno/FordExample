//Gordon Stangler
//Goroutines, with a dash of extern C

package main

import(
	"fmt"
)

//extern C
//extern func Cdisp(sentArray string) string

func coarrayEX(sentArray string){
//	sentArray []string = "hello world"

	//extern C
	returnedArray := Cdisp(sentArray)
	fmt.Printf("%s\n", returnedArray)
}

func main(){
	go coarrayEX("hello")
	go coarrayEX("world")
}