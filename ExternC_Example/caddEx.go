//Gordon Stangler
//Goroutines, with a dash of extern C

package main

import(
	"fmt"
)

func coarrayEX(sentArray string){
//	sentArray = "hello world"

	//extern C
	returnedArray string = func Cdisp(*byte sentArray)
		fmt.printf("%s\n", returnedArray)
}

func main
	go coarrayEX("hello")
	go coarrayEX("world")
}