//Gordon Stangler
//GoLang AJAX example
package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

var Index_HTML []byte
var mux *http.ServeMux

func main(){
	mux = http.NewServeMux();
	mux.Handle("/", http.FileServe(http.Dir("./public")))
	//fmt.Println("Starting server on localhost:8888/")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe(":8888", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	log.Println("GET /")
	w.Write(Index_HTML)
}


//Looks for a static file.  Lazy, and stolen, but hey, it works.
func static(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.ContainsRune(r.URL.Path, '.'){
			mux.ServeHTTP(w,r)
			return
		}
		h.ServeHTTP(w,r)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	//checks to see if method is POST, otherwise, return 404.
	if r.Method != "POST" {
		http.NotFound(w,r)
		return
	}

	field := r.FormValue("textfield")
	fmt.Println("textfield: ", field)
	r.ParseForm()
	log.Println("in post handler", r.Form)


	//DB Query!
	//ignoring because we are not doing a DB query, yet.
		//maybe later

	w.Write([]byte(field))
}

func init(){
	Index_HTML, _ = ioutil.ReadFile("./html/index.html")
}
