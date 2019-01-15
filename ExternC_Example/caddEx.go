//Gordon Stangler
//Goroutines, with a dash of extern C

package main

import(
	"fmt"
)

func coarrayEX(sentArray string){
//	sentArray []string = "hello world"

	//extern C
	returnedArray []string = Cdisp(*byte sentArray)
	fmt.printf("%s\n", returnedArray)
}

func main(){
	go coarrayEX("hello")
	go coarrayEX("world")
}