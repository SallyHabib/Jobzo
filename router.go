package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
	"log"
	"encoding/json"
	"io/ioutil"
//	"strings"
//	cors "github.com/heppu/simple-cors"
)

type (
	// Session Holds info about a session
	Session map[string]interface{}

	// JSON Holds a JSON object
	JSON map[string]interface{}

	// Processor Alias for Process func
	Processor func(session Session, message string) (string, error)
)

type Message struct {
	message string
}

type Result struct {
    Response string `json:"response"`
}

func writeJSON(w http.ResponseWriter, data JSON) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome! Andrew \n")
}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// mapD := map[string]string{"message": "Welcome! what do you want to order?"}
    // mapB, _ := json.Marshal(mapD)
	// fmt.Fprintf(w, string(mapB))
	writeJSON(w, JSON{
		"message": "Welcome what do you want to order?",
	})
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Index2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, _ := http.Get("http://api.glassdoor.com/api/api.htm?v=1&format=json&t.p=216693&t.k=bhhitzZ6DNo&action=employers&q=pharmaceuticals&userip=192.168.43.42&useragent=Mozilla/%2F4.0")
    responseData , err := ioutil.ReadAll(response.Body)
	if err != nil {
	 	log.Fatal(err)
	 }
	defer response.Body.Close()
	//message, _ := resp["response"].(string)
	result := Result{} 
	json.Unmarshal(responseData, &result)
	fmt.Fprintf(w, result.Response)
}

func chat(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	resp := JSON{}
	json.NewDecoder(r.Body).Decode(&resp)
	defer r.Body.Close()
	message, _ := resp["message"].(string)
	writeJSON(w, JSON{
	 	"message": message,
	})
}

func main() {
    router := httprouter.New()
	router.GET("/", Index)
	router.GET("/welcome", welcome)
	router.GET("/glassdoor", Index2)
	router.GET("/hello/:name", Hello)
	
	router.POST("/chat",chat)

    log.Fatal(http.ListenAndServe(":8080", router))
}
