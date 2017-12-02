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
var Thanking = [...]string{"Thank u", "Thanks", "Thank you", "Merci"}

// Goodbyes ... array
var Goodbyes = [...]string{"bye", "au revoir", "salam"}

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
	if strings.Contains(job, "job") {
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
	if i == 0 {
		err := errors.New("No jobs found")
		return result, err
	}

	return result, err
}

//SearchForLocalJobsWuzzuf ... function
func SearchForLocalJobsWuzzuf(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(job, "job") {
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
	if i == 0 {
		err := errors.New("No jobs found")
		return result, err
	}
	return result, err
}

// SearchForGlobalJobs ... function
func SearchForGlobalJobs(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(job, "job") {
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
	if i == 0 {
		err := errors.New("No jobs found")
		return result, err
	}
	return result, err
}

// SearchForGlobalJobsGlassdoor ... function
func SearchForGlobalJobsGlassdoor(searchWord string, job string, country string) (models.Response, error) {
	var kind string
	var link string
	if strings.Contains(job, "job") {
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
	if i == 0 {
		err := errors.New("No jobs found")
		return result, err
	}
	return result, err
}

// HandleJobs ... function
func HandleJobs(session models.Session, input string) (string, models.Response, error) {
	_, arrayFound := session["preferences"]
	if !arrayFound {
		var array []string
		session["preferences"] = array
	}

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
		switch strings.ToLower(input) {
		case "job", "internship":
			userInputs = append(userInputs, input)
			session["preferences"] = userInputs
			counter++
			session["counter"] = counter
			return "which country?", resp, nil
		default:
			return "Please choose Job or Internship", resp, nil
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
		if strings.ToLower(userInputs[2]) == "egypt" {
			messageResp, err = SearchForLocalJobs(userInputs[0], userInputs[1], userInputs[2])
			messageResp2, err = SearchForLocalJobsWuzzuf(userInputs[0], userInputs[1], userInputs[2])
		} else {
			messageResp, err = SearchForGlobalJobs(userInputs[0], userInputs[1], userInputs[2])
			messageResp2, err = SearchForGlobalJobsGlassdoor(userInputs[0], userInputs[1], userInputs[2])
		}

		// concatenating 2 results
		i := 0
		for i < len(messageResp2.Items) {
			messageResp.Items = append(messageResp.Items, messageResp2.Items[i])
			i = i + 1
		}
		userInputs = userInputs[:0]
		session["preferences"] = userInputs
		counter = 0
		session["counter"] = counter

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
		err := errors.New("No Courses found")
		return result, err
	}
	return result, err
}

// HandleCourses ... function
func HandleCourses(session models.Session, input string) (string, models.Response, error) {
	_, arrayFound := session["courses"]
	if !arrayFound {
		var array []string
		session["courses"] = array
	}

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
		switch strings.ToLower(input) {
		case "beginner", "intermediate", "advanced":
			newInput := strings.Replace(input, " ", "%20", -1)
			userInputs = append(userInputs, newInput)
			fmt.Println(userInputs[0], userInputs[1])
			messageResp, err := SearchForCourses(userInputs[0], userInputs[1])
			userInputs = userInputs[:0]
			session["courses"] = userInputs
			counter = 0
			session["coursesCounter"] = counter
			return "", messageResp, err
		default:
			return "Please choose one of the following beginner, intermediate or advanced", resp, nil
		}
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
		err := errors.New("No " + kind + " found")
		return result, err
	}
	return result, err
}

// HandleDegrees ... function
func HandleDegrees(session models.Session, input string) (string, models.Response, error) {
	_, arrayFound := session["degrees"]
	if !arrayFound {
		var array []string
		session["degrees"] = array
	}

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
		switch strings.ToLower(input) {
		case "bachelor", "masters", "phd":
			userInputs = append(userInputs, input)
			session["degrees"] = userInputs
			counter++
			session["degreesCounter"] = counter
			return "which country?", resp, nil
		default:
			return "Please enter one of the 3 Choices PHD, Masters or Bachelor", resp, nil
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

	for _, c := range Thanking {
		if strings.ToLower(input) == strings.ToLower(c) {
			return "7byby teslam", resp, nil
		}
	}

	for _, c := range Goodbyes {
		if strings.ToLower(input) == strings.ToLower(c) {
			return "Bye", resp, nil
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
		switch strings.ToLower(input) {
		case "jobs":
			session["counter"] = 0
			validate++
			session["validate"] = validate
			scenario = 0
			session["scenario"] = scenario
		case "courses":
			session["coursesCounter"] = 0
			validate++
			session["validate"] = validate
			scenario = 1
			session["scenario"] = scenario
		case "degrees":
			session["degreesCounter"] = 0
			validate++
			session["validate"] = validate
			scenario = 2
			session["scenario"] = scenario
		case "restart":
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
		return "Please Choose jobs, degrees, courses OR choose from the Buttons", resp, nil
	default:
		switch scenario {
		case 0:
			message, resp, err = HandleJobs(session, input)
		case 1:
			message, resp, err = HandleCourses(session, input)
		case 2:
			message, resp, err = HandleDegrees(session, input)
		}
	}

	return message, resp, err
}
