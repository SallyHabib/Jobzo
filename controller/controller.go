package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"../models"
)

// Thanking ... array
var Thanking = [...]string{"thanks", "shokran", "thnx", "sanko", "zanko", "sankyo", "zankyo", "merci", "rbna y5lek", "thank you", "thank u", "shokr"}

// Emojis ... array
var Emojis = [...]string{":D", "😜", ":)", ";P", ":O", "(y)", ":P", "B)", "B-)", "8)", "8-)", "^_^", ":*", "O:)", "😂", ";)", "3:)", "<3"}

// Complements ... array
var Complements = [...]string{"nice", "7byby", "zy el fol", "7abebe", "7bb", "zalfol", "cool", "awesome", "great", "ur good", "danta dma8", "danta dma9", "danta dema8", "danta dma3'", "danta dema3'", "cute", "awsome", "lol", "it is good", "u r good"}

// Goodbyes ... array
var Goodbyes = [...]string{"bye", "au revoir", "salam"}

var jobs = false

// getCountries ... function
func getCountries() []models.Country {
	raw, err := ioutil.ReadFile("./countries.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []models.Country
	json.Unmarshal(raw, &c)
	return c
}

// SearchForLocalJobs ... function
func SearchForLocalJobs(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(strings.ToLower(job), "job") {
		kind = "jobs"
	} else {
		kind = "internships"
	}
	link = "https://www.googleapis.com/customsearch/v1?q=" + searchWord + "%20" + kind + "%20in%20" + country + "&key=AIzaSyAeALD2cLr3-NSEoOz2wUjLMhaOOxgLUN0&cx=006422052657745549454:vmlxelexg7y&num=5"
	response, err := http.Get(link)
	result := models.Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 || i == 1 {
		newSearchWord := strings.Replace(searchWord, "%20", " ", -1)
		err := errors.New("No " + newSearchWord + " " + kind + " found in " + country)
		return result, err
	}

	return result, nil
}

//SearchForLocalJobsWuzzuf ... function
func SearchForLocalJobsWuzzuf(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(strings.ToLower(job), "job") {
		kind = "jobs"
	} else {
		kind = "internships"
	}
	link = "https://www.googleapis.com/customsearch/v1?q=" + searchWord + "%20" + kind + "%20in%20" + country + "&key=AIzaSyAeALD2cLr3-NSEoOz2wUjLMhaOOxgLUN0&cx=006422052657745549454:gj2panfjzja&num=5"
	response, err := http.Get(link)
	result := models.Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 || i == 1 {
		newSearchWord := strings.Replace(searchWord, "%20", " ", -1)
		err := errors.New("No " + newSearchWord + " " + kind + " found in " + country)
		return result, err
	}
	return result, nil
}

// SearchForGlobalJobs ... function
func SearchForGlobalJobs(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(strings.ToLower(job), "job") {
		kind = "jobs"
	} else {
		kind = "internships"
	}
	link = "https://www.googleapis.com/customsearch/v1?q=" + searchWord + "%20" + kind + "%20in%20" + country + "&key=AIzaSyAeALD2cLr3-NSEoOz2wUjLMhaOOxgLUN0&cx=006422052657745549454:xojc8tra6ua&num=5"
	response, err := http.Get(link)
	result := models.Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 || i == 1 {
		newSearchWord := strings.Replace(searchWord, "%20", " ", -1)
		err := errors.New("No " + newSearchWord + " " + kind + " found in " + country)
		return result, err
	}
	return result, nil
}

// SearchForGlobalJobsGlassdoor ... function
func SearchForGlobalJobsGlassdoor(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(strings.ToLower(job), "job") {
		kind = "jobs"
	} else {
		kind = "internships"
	}
	link = "https://www.googleapis.com/customsearch/v1?q=" + searchWord + "%20" + kind + "%20in%20" + country + "&key=AIzaSyAeALD2cLr3-NSEoOz2wUjLMhaOOxgLUN0&cx=006422052657745549454:l5p6gvxphiy&num=5"
	response, err := http.Get(link)
	result := models.Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 || i == 1 {
		newSearchWord := strings.Replace(searchWord, "%20", " ", -1)
		err := errors.New("No " + newSearchWord + " " + kind + " found in " + country)
		return result, err
	}
	return result, nil
}

// HandleJobs ... function
func HandleJobs(session models.Session, input string) (string, models.Response, error) {
	counter, _ := session["counter"].(int)
	userInputs, _ := session["preferences"].([]string)
	resp := models.Response{}

	switch counter {
	case 0:
		counter++
		session["counter"] = counter
		return "What field are you interested in?", resp, nil
	case 1:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		session["preferences"] = userInputs
		counter++
		session["counter"] = counter
		return "Are you looking for a job or an internship?", resp, nil
	case 2:
		if strings.Contains(strings.ToLower(input), "job") {
			userInputs = append(userInputs, "job")
			session["preferences"] = userInputs
			counter++
			session["counter"] = counter
			return "which country?", resp, nil
		} else if strings.Contains(strings.ToLower(input), "internship") {
			userInputs = append(userInputs, "internship")
			session["preferences"] = userInputs
			counter++
			session["counter"] = counter
			return "which country?", resp, nil
		} else {
			return "Please choose either a Job or an Internship", resp, nil
		}

	case 3:
		countries := getCountries()
		found := false

		for _, c := range countries {
			if strings.ToLower(input) == strings.ToLower(c.Name) {
				found = true
				break
			}
		}

		if !found {
			return "Please enter a valid country", resp, nil
		}

		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		fmt.Println(userInputs[0], userInputs[1], userInputs[2])
		messageResp := models.Response{}
		messageResp2 := models.Response{}

		var err error
		var err2 error
		if strings.ToLower(userInputs[2]) == "egypt" {
			messageResp, err = SearchForLocalJobs(userInputs[0], userInputs[1], userInputs[2])
			messageResp2, err2 = SearchForLocalJobsWuzzuf(userInputs[0], userInputs[1], userInputs[2])
		} else {
			messageResp, err = SearchForGlobalJobs(userInputs[0], userInputs[1], userInputs[2])
			messageResp2, err2 = SearchForGlobalJobsGlassdoor(userInputs[0], userInputs[1], userInputs[2])
		}

		userInputs = userInputs[:0]
		session["preferences"] = userInputs
		counter = 0
		session["counter"] = counter

		if err == nil && err2 == nil {
			i := 0
			for i < len(messageResp2.Items) {
				messageResp.Items = append(messageResp.Items, messageResp2.Items[i])
				i = i + 1
			}
			return "", messageResp, err
		} else if err == nil && err2 != nil {
			return "", messageResp, err
		} else if err != nil && err2 == nil {
			return "", messageResp2, err2
		}
		return "", messageResp, err
	}
	return "", resp, nil
}

// SearchForCourses ... function
func SearchForCourses(searchWord string, kind string) (models.Response, error) {
	var link string
	link = "https://www.googleapis.com/customsearch/v1?q=" + kind + "%20" + searchWord + "%20course&key=AIzaSyACe9KwpeP0u7Aubb4TiFJ5sPD0jJ2rvPo&cx=013526896367865193935:qq9kmjp5tr8&num=10"
	response, err := http.Get(link)
	result := models.Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 {
		newSearchWord := strings.Replace(searchWord, "%20", " ", -1)
		err := errors.New("No " + kind + " " + newSearchWord + " courses found")
		return result, err
	}
	return result, err
}

// HandleCourses ... function
func HandleCourses(session models.Session, input string) (string, models.Response, error) {
	counter, _ := session["coursesCounter"].(int)
	userInputs, _ := session["courses"].([]string)
	resp := models.Response{}

	switch counter {
	case 0:
		counter++
		session["coursesCounter"] = counter
		return "What course you want to learn?", resp, nil
	case 1:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		session["courses"] = userInputs
		counter++
		session["coursesCounter"] = counter
		return "Are you searching for beginner, intermediate or advanced course?", resp, nil
	case 2:
		if !strings.Contains(strings.ToLower(input), "beginner") && !strings.Contains(strings.ToLower(input), "intermediate") && !strings.Contains(strings.ToLower(input), "advanced") {
			return "Please specify which type of course you want", resp, nil
		}
		if strings.Contains(strings.ToLower(input), "beginner") {
			userInputs = append(userInputs, "beginner")
		} else if strings.Contains(strings.ToLower(input), "intermediate") {
			userInputs = append(userInputs, "intermediate")
		} else if strings.Contains(strings.ToLower(input), "advanced") {
			userInputs = append(userInputs, "advanced")
		}
		messageResp, err := SearchForCourses(userInputs[0], userInputs[1])
		userInputs = userInputs[:0]
		session["courses"] = userInputs
		counter = 0
		session["coursesCounter"] = counter
		return "", messageResp, err
	}
	return "", resp, nil
}

// SearchForDegrees ... function
func SearchForDegrees(searchWord string, kind string, country string) (models.Response, error) {
	var link string
	link = "https://www.googleapis.com/customsearch/v1?q=" + searchWord + "%20" + kind + "%20in%20" + country + "&key=AIzaSyAeALD2cLr3-NSEoOz2wUjLMhaOOxgLUN0&cx=013526896367865193935:quzhbm-uvts&num=10"
	fmt.Println("LINK: ", link)
	response, err := http.Get(link)
	result := models.Response{}
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&result)
	i, err := strconv.Atoi(result.Info.Num)
	if i == 0 {
		newSearchWord := strings.Replace(searchWord, "%20", " ", -1)
		err := errors.New("No " + newSearchWord + " " + kind + " found in " + country)
		return result, err
	}
	return result, err
}

// HandleDegrees ... function
func HandleDegrees(session models.Session, input string) (string, models.Response, error) {
	counter, _ := session["degreesCounter"].(int)
	userInputs, _ := session["degrees"].([]string)
	resp := models.Response{}

	switch counter {
	case 0:
		counter++
		session["degreesCounter"] = counter
		return "What topic are you interested in?", resp, nil
	case 1:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		session["degrees"] = userInputs
		counter++
		session["degreesCounter"] = counter
		return "Are you looking for a Bachelor , Masters or PHD?", resp, nil
	case 2:
		if !strings.Contains(strings.ToLower(input), "bachelor") && !strings.Contains(strings.ToLower(input), "masters") && !strings.Contains(strings.ToLower(input), "phd") {
			return "Please specify which type of degree you want to achieve", resp, nil
		}
		if strings.Contains(strings.ToLower(input), "bachelor") {
			userInputs = append(userInputs, "bachelor")
		} else if strings.Contains(strings.ToLower(input), "masters") {
			userInputs = append(userInputs, "masters")
		} else if strings.Contains(strings.ToLower(input), "phd") {
			userInputs = append(userInputs, "phd")
		}
		session["degrees"] = userInputs
		counter++
		session["degreesCounter"] = counter
		return "which country?", resp, nil
	case 3:
		countries := getCountries()
		found := false

		for _, c := range countries {
			if strings.ToLower(input) == strings.ToLower(c.Name) {
				found = true
				break
			}
		}

		if !found {
			return "Please enter a valid country", resp, nil
		}

		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		fmt.Println(userInputs[0], userInputs[1], userInputs[2])
		messageResp, err := SearchForDegrees(userInputs[0], userInputs[1], userInputs[2])
		userInputs = userInputs[:0]
		session["degrees"] = userInputs
		counter = 0
		session["degreesCounter"] = counter
		return "", messageResp, err
	}
	return "", resp, nil
}

// HandleSequence ... function
func HandleSequence(session models.Session, input string) (string, models.Response, error) {
	var message string
	var err error
	resp := models.Response{}
	newInput := strings.TrimSpace(input)
	fmt.Println(newInput)

	for _, c := range Thanking {
		if strings.Contains(strings.ToLower(newInput), strings.ToLower(c)) {
			return "Urw ^_^", resp, nil
		}
	}

	for _, c := range Emojis {
		if strings.Contains(strings.ToLower(newInput), strings.ToLower(c)) {
			return input, resp, nil
		}
	}

	for _, c := range Complements {
		if strings.Contains(strings.ToLower(newInput), strings.ToLower(c)) {
			return "7byby teslam >_<", resp, nil
		}
	}

	for _, c := range Goodbyes {
		if strings.Contains(strings.ToLower(newInput), strings.ToLower(c)) {
			return "Bye :D", resp, nil
		}
	}

	_, validateFound := session["validate"]
	if !validateFound {
		session["validate"] = 0
	}
	validate, _ := session["validate"].(int)

	_, Found := session["initialize"]
	if !Found {
		session["initialize"] = 0
	}
	initialize, _ := session["initialize"].(int)

	_, scenarioFound := session["scenario"]
	if !scenarioFound {
		session["scenario"] = -1
	}
	scenario, _ := session["scenario"].(int)

	switch initialize {
	case 0:
		initialize++
		session["initialize"] = initialize
		choices := "What do you want to search for?" + "\n" +
			"1) Jobs & Internships (type jobs)" + "\n" +
			"2) Courses (type courses)" + "\n" +
			"3) Bachelor, Masters & PHD Degrees (type degrees)"
		return choices, resp, nil

	case 1:
		var array []string
		if strings.Contains(strings.ToLower(newInput), "jobs") {
			session["preferences"] = array
			session["counter"] = 0
			validate++
			session["validate"] = validate
			scenario = 0
			session["scenario"] = scenario
		} else if strings.Contains(strings.ToLower(newInput), "courses") {
			session["courses"] = array
			session["coursesCounter"] = 0
			validate++
			session["validate"] = validate
			scenario = 1
			session["scenario"] = scenario
		} else if strings.Contains(strings.ToLower(newInput), "degrees") {
			session["degrees"] = array
			session["degreesCounter"] = 0
			validate++
			session["validate"] = validate
			scenario = 2
			session["scenario"] = scenario
		} else if strings.Contains(strings.ToLower(newInput), "restart") {
			validate = 0
			session["validate"] = validate
			choices := "What do you want to search for?" + "\n" +
				"1) Jobs & Internships (type jobs)" + "\n" +
				"2) Courses (type courses)" + "\n" +
				"3) Bachelor, Masters & PHD Degrees (type degrees)"
			var array []string
			session["preferences"] = array
			session["courses"] = array
			session["degrees"] = array
			session["counter"] = 0
			session["coursesCounter"] = 0
			session["degreesCounter"] = 0

			return choices, resp, nil
		}
	}

	switch session["validate"] {
	case 0:
		return "Please Choose jobs, degrees, courses Or choose from the buttons", resp, nil
	default:
		switch scenario {
		case 0:
			message, resp, err = HandleJobs(session, newInput)
		case 1:
			message, resp, err = HandleCourses(session, newInput)
		case 2:
			message, resp, err = HandleDegrees(session, newInput)
		}
	}

	return message, resp, err
}
