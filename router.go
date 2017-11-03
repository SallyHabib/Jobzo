package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	cors "github.com/heppu/simple-cors"
)

var (
	sessions  = map[string]Session{}
	processor = sampleProcessor
)

type (
	// Session Holds info about a session
	Session map[string]interface{}

	// JSON Holds a JSON object
	JSON map[string]interface{}

	// Processor Alias for Process func
	Processor func(session Session, message string) (string, error)
)

// Message ... type
type Message struct {
	message string
}

// Result ... type
type Result struct {
	Response string `json:"response"`
}

//Data ... type
type Data struct {
	//Success bool "json:city"

	Status   string `json:status`
	Success  bool   `json:success`
	Response string `json:response`
}

// WriteJSON ... function
func WriteJSON(w http.ResponseWriter, data JSON) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func sampleProcessor(session Session, message string) (string, error) {
	// Make sure a history key is defined in the session which points to a slice of strings
	_, historyFound := session["history"]
	if !historyFound {
		session["history"] = []string{}
	}

	// Fetch the history from session and cast it to an array of strings
	history, _ := session["history"].([]string)

	// Make sure the message is unique in history
	for _, m := range history {
		if strings.EqualFold(m, message) {
			return "", fmt.Errorf("You've already ordered %s before!", message)
		}
	}

	// Add the message in the parsed body to the messages in the session
	history = append(history, message)

	// Form a sentence out of the history in the form Message 1, Message 2, and Message 3
	l := len(history)
	wordsForSentence := make([]string, l)
	copy(wordsForSentence, history)
	if l > 1 {
		wordsForSentence[l-1] = "and " + wordsForSentence[l-1]
	}
	sentence := strings.Join(wordsForSentence, ", ")

	// Save the updated history to the session
	session["history"] = history

	return fmt.Sprintf("So, you want %s! What else?", strings.ToLower(sentence)), nil
}

// ProcessFunc Sets the processor of the chatbot
func ProcessFunc(p Processor) {
	processor = p
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
	sessions[uuid] = Session{}
	WriteJSON(w, JSON{
		"message": "Welcome what do you want to order?",
		"uuid":    uuid,
	})
}

// SearchGlassdoor ... function
func SearchGlassdoor(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("http://api.glassdoor.com/api/api.htm?v=1&format=json&t.p=216693&t.k=bhhitzZ6DNo&action=employers&q=pharmaceuticals")
	/*//	responseData, err := ioutil.ReadAll(response.Body)

	defer response.Body.Close()
	//message, _ := resp["response"].(string)*/

	result := Data{}
	json.NewDecoder(response.Body).Decode(&result)

	// data, err := ioutil.ReadAll(response.Body)

	// if err == nil && data != nil {
	// 	err = json.Unmarshal(data, &result)

	// }
	fmt.Fprint(w, result)
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
	// Make sure a session exists for the extracted UUID
	session, sessionFound := sessions[uuid]
	if !sessionFound {
		http.Error(w, fmt.Sprintf("No session found for: %v.", uuid), http.StatusUnauthorized)
		return
	}
	// Parse the JSON string in the body of the request
	data := JSON{}
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
	// Process the received message
	message, err := processor(session, data["message"].(string))
	if err != nil {
		http.Error(w, err.Error(), 422 /* http.StatusUnprocessableEntity */)
		return
	}
	// Write a JSON containg the processed response
	WriteJSON(w, JSON{
		"message": message,
	})
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
	mux.HandleFunc("/welcome", withLog(SearchGlassdoor))
	mux.HandleFunc("/chat", withLog(Chat))
	mux.HandleFunc("/", withLog(handle))
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", cors.CORS(mux)))
}
