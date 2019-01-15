//Gordon Stangler
//Goroutines, with a dash of extern C
//C function is in cadd.c

package main

import "C"
import(
	"fmt"
)

//extern C
func Cdisp(sentArray string) string

func coarrayEX(sentArray string){
	//extern C
	returnedArray := Cdisp(sentArray)
	fmt.Printf("%s\n", returnedArray)
}

func main(){
	go coarrayEX("hello")
	go coarrayEX("world")
}