package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
	"log"
	"io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome! Andrew \n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Index2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := http.Get("http://api.glassdoor.com/api/api.htm?v=1&format=json&t.p=216693&t.k=bhhitzZ6DNo&action=employers&q=pharmaceuticals&userip=192.168.43.42&useragent=Mozilla/%2F4.0")
    responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w,string(responseData))
}

func main() {
    router := httprouter.New()
	router.GET("/", Index)
	router.GET("/glassdoor", Index2)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
