package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"time"

	"./controller"
	"./models"

	cors "github.com/heppu/simple-cors"
)

var (
	sessions = map[string]models.Session{}
)

// WriteJSON ... function
func WriteJSON(w http.ResponseWriter, data models.JSON) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// WriteJSON2 ... function
func WriteJSON2(w http.ResponseWriter, data models.Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// withLog Wraps HandlerFuncs to log requests to Stdout
func withLog(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := httptest.NewRecorder()
		fn(c, r)
		log.Printf("[%d] %-4s %s\n", c.Code, r.Method, r.URL.Path)

		for k, v := range c.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)
		c.Body.WriteTo(w)
	}
}

// Welcome ... function
func Welcome(w http.ResponseWriter, r *http.Request) {
	// Generate a UUID.
	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	uuid := hex.EncodeToString(hasher.Sum(nil))
	// Create a session for this UUID
	sessions[uuid] = models.Session{}
	WriteJSON(w, models.JSON{
		"message": "Hiii",
		"uuid":    uuid,
	})
}

// Chat ... function
func Chat(w http.ResponseWriter, r *http.Request) {
	// Make sure only POST requests are handled
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	// Make sure a UUID exists in the Authorization header
	uuid := r.Header.Get("Authorization")
	if uuid == "" {
		http.Error(w, "Missing or empty Authorization header.", http.StatusUnauthorized)
		return
	}
	//Make sure a session exists for the extracted UUID
	session, sessionFound := sessions[uuid]
	if !sessionFound {
		http.Error(w, fmt.Sprintf("No session found for: %v.", uuid), http.StatusUnauthorized)
		return
	}
	// Parse the JSON string in the body of the request
	data := models.JSON{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Couldn't decode JSON: %v.", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// Make sure a message key is defined in the body of the request
	_, messageFound := data["message"]
	if !messageFound {
		http.Error(w, "Missing message key in body.", http.StatusBadRequest)
		return
	}

	message, response, err := controller.HandleSequence(session, data["message"].(string))
	if err != nil {
		WriteJSON(w, models.JSON{
			"message":    err,
			"statusCode": 422,
		})
	} else if message != "" {
		WriteJSON(w, models.JSON{
			"message":    message,
			"statusCode": 200,
		})
	} else {
		WriteJSON2(w, response)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	body :=
		"<!DOCTYPE html><html><head><title>Chatbot</title></head><body><pre style=\"font-family: monospace;\">\n" +
			"Available Routes:\n\n" +
			"  GET  /welcome -> Welcome\n" +
			"  POST /chat    -> Chat\n" +
			"  GET  /        -> handle        (current)\n" +
			"</pre></body></html>"
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(w, body)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", withLog(Welcome))
	mux.HandleFunc("/chat", withLog(Chat))
	mux.HandleFunc("/", withLog(handle))

	//Start the server
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, cors.CORS(mux)))
}
