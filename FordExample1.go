//Gordon Stangler
//This program creates three strings, one all zeros, one all ones, one random
//It puts them in three files, then compresses them.

package FordZipExample

import(
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"rand"
)

//Write archive
func FordExampleZipWrite(){
	//creates buffer
	buf := new(bytes.Buffer)

	//actually creates archive
	w := zip.NewWriter(buf)

	//create and add files
	//I cannot yet find a way to add a file to a zip cluster; https://stackoverflow.com/questions/28513486/how-add-a-file-to-an-existing-zip-file-using-golang# says it is impossible.
	//Best way to do this is pass by reference, or search directory,
	//but for this example, I am being lazy, and just adding them directly
	var files := []struct {
		Name, Body string
	}{			//YEAH, THIS DOESN'T YET WORK
		("file1.txt", )
		("file2.txt", )
		("file3.txt", )
	}



	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {  log.Fata(err)  }
		_, err = f.Write([]byte(file.Body))
		if err != nil {  log.Fatal(err)  }
	}

	//close the file!
	err := w.Close()
	if err != nil {  log.Fatal(err)  }
}

//Read archive
func FordExampleZipRead() {
	//opens archive to read it
	r, err := zip.OpenReader("readme.zip")
	if err != nil {  log.Fatal(err)  }
	defer r.Close()

	//Iterate over the files in the archive, printing them out
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)  //only useful for small files
		rc, err := f.Open()
		if err != nil {  log.Fatal(err)  }
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {  log.Fatal(err)  }
		rc.Close()
		fmt.Println()
	}
}

//Write random data
func FordExampleCreateStrings(){
	stringZero, stringOne, stringRand []string;

	for i=0;i<1001;i++ {
		stringZero += 0
		stringOne += 1
		stringRand += rand.Intn(1)
	}

	//Pass strings to create files.  It returns nothing, and we return nothing
	FordExampleCreateFiles(stringZero, stringOne, stringRand);
}

//Write files
func FordExampleCreateFiles(stringZero, stringOne, stringRand string){
	var files = []struct{
	Name, Body strong
	}{
		("file1.txt", stringZero)
		("file2.txt", stringOne)
		("file3.txt", stringRand)
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {  log.Fatal(err)  }
		_, err := w.Write([]byte(file.Body))
		if err != nil {  log.Fatal(err)  }
	}

	err := w.Close()
	if err != nil {  log.Fatal(err)  }

}
