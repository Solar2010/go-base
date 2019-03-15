package main

import(
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"<h1>你好 晓晨!</h1>")
}

func main(){
	http.HandleFunc("/",IndexHandler)
	http.ListenAndServe("127.0.0.1:8888",nil)
}
