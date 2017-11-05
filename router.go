package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"time"

	cors "github.com/heppu/simple-cors"
)
import _ "github.com/joho/godotenv/autoload"

var (
	sessions = map[string]Session{}
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

// Response ... type
type Response struct {
	Info  SearchInfo `json:"searchInformation"`
	Items []Job      `json:"items"`
}

// SearchInfo ... type
type SearchInfo struct {
	Num string `json:"totalResults"`
}

// Job ... type
type Job struct {
	Title string    `json:"title"`
	Link  string    `json:"link"`
	Image Thumbnail `json:"pagemap"`
}

// Thumbnail ... type
type Thumbnail struct {
	CseImage []ImageSrc `json:"cse_image"`
}

// ImageSrc ... type
type ImageSrc struct {
	Src string `json:"src"`
}

// WriteJSON ... function
func WriteJSON(w http.ResponseWriter, data JSON) {
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
	sessions[uuid] = Session{}
	WriteJSON(w, JSON{
		"message": "Hiii",
		"uuid":    uuid,
	})
}

// SearchGoogle ... function
func SearchGoogle(searchWord string, job string, country string) (Response, error) {
	var kind string
	var link string
	if strings.Contains(job, "job") {
		kind = "jobs"
	} else {
		kind = "internships"
	}
	link = "https://www.googleapis.com/customsearch/v1?q=" + searchWord + "%20" + kind + "%20in%20" + country + "&key=AIzaSyAeALD2cLr3-NSEoOz2wUjLMhaOOxgLUN0&cx=006422052657745549454:vmlxelexg7y&num=10"
	response, err := http.Get(link)
	result := Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 {
		err := errors.New("No jobs found")
		return result, err
	}
	return result, err
}

// Urls ... function
func Urls(x Response) string {
	i := 0
	message := ""
	for i < len(x.Items) {
		message += "\n" + x.Items[i].Link + "\n"
		i++
	}
	return message
}

// HandleSequence ... function
func HandleSequence(session Session, input string) (string, error) {
	_, arrayFound := session["preferences"]
	if !arrayFound {
		var array []string
		session["preferences"] = array
	}
	_, counterFound := session["counter"]
	if !counterFound {
		session["counter"] = 0
	}
	counter, _ := session["counter"].(int)
	userInputs, _ := session["preferences"].([]string)

	switch counter {
	case 0:
		counter++
		session["counter"] = counter
		return "What field are you interested in?", nil
	case 1:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		session["preferences"] = userInputs
		counter++
		session["counter"] = counter
		return "Are you looking for a job or an internship?", nil
	case 2:
		userInputs = append(userInputs, input)
		session["preferences"] = userInputs
		counter++
		session["counter"] = counter
		return "which country?", nil
	case 3:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		fmt.Println(userInputs[0], userInputs[1], userInputs[2])
		messageResp, err := SearchGoogle(userInputs[0], userInputs[1], userInputs[2])
		message := Urls(messageResp)
		userInputs = userInputs[:0]
		session["preferences"] = userInputs
		counter = 0
		session["counter"] = counter
		return message, err
	}
	return "", nil
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

	message, err := HandleSequence(session, data["message"].(string))
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
