package player

import (
	"fmt"
	"net/http"
	"strings"
)

func Hash_handler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("schema",r.URL.Scheme)
	fmt.Println("host", r.URL.Host)
	for k,v := range r.Form{
		fmt.Println("k:",k)
		fmt.Println("v:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello world!")
}
