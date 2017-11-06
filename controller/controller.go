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

// Urls ... function
func Urls(x models.Response) string {
	i := 0
	message := ""
	for i < len(x.Items) {
		message += x.Items[i].Link + "<br>"
		i++
	}
	return message
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
func HandleJobs(session models.Session, input string) (string, error) {
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
		countries := getCountries()
		found := false

		for _, c := range countries {
			fmt.Println(strings.ToLower(c.Name))
			if strings.ToLower(input) == strings.ToLower(c.Name) {
				found = true
				break
			}
		}

		if !found {
			return "Please enter a valid country", nil
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
		message := Urls(messageResp)
		message2 := Urls(messageResp2)

		userInputs = userInputs[:0]
		session["preferences"] = userInputs
		counter = 0
		session["counter"] = counter
		resultReturn := message + message2
		return resultReturn, err
	}
	return "", nil
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
func HandleCourses(session models.Session, input string) (string, error) {
	_, arrayFound := session["courses"]
	if !arrayFound {
		var array []string
		session["courses"] = array
	}
	_, counterFound := session["coursesCounter"]
	if !counterFound {
		session["coursesCounter"] = 0
	}
	counter, _ := session["coursesCounter"].(int)
	userInputs, _ := session["courses"].([]string)

	switch counter {
	case 0:
		counter++
		session["coursesCounter"] = counter
		return "What course you want to learn?", nil
	case 1:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		session["courses"] = userInputs
		counter++
		session["coursesCounter"] = counter
		return "Are you searching for beginner, intermediate or advanced course?", nil
	case 2:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		fmt.Println(userInputs[0], userInputs[1])
		messageResp, err := SearchForCourses(userInputs[0], userInputs[1])
		message := Urls(messageResp)
		userInputs = userInputs[:0]
		session["courses"] = userInputs
		counter = 0
		session["coursesCounter"] = counter
		return message, err
	}
	return "", nil
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
func HandleDegrees(session models.Session, input string) (string, error) {
	_, arrayFound := session["degrees"]
	if !arrayFound {
		var array []string
		session["degrees"] = array
	}
	_, counterFound := session["degreesCounter"]
	if !counterFound {
		session["degreesCounter"] = 0
	}
	counter, _ := session["degreesCounter"].(int)
	userInputs, _ := session["degrees"].([]string)

	switch counter {
	case 0:
		counter++
		session["degreesCounter"] = counter
		return "What topic are you interested in?", nil
	case 1:
		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		session["degrees"] = userInputs
		counter++
		session["degreesCounter"] = counter
		return "Are you looking for a Bachelor , Masters or PHD?", nil
	case 2:
		userInputs = append(userInputs, input)
		session["degrees"] = userInputs
		counter++
		session["degreesCounter"] = counter
		return "which country?", nil
	case 3:
		countries := getCountries()
		found := false

		for _, c := range countries {
			fmt.Println(strings.ToLower(c.Name))
			if strings.ToLower(input) == strings.ToLower(c.Name) {
				found = true
				break
			}
		}

		if !found {
			return "Please enter a valid country", nil
		}

		newInput := strings.Replace(input, " ", "%20", -1)
		userInputs = append(userInputs, newInput)
		fmt.Println(userInputs[0], userInputs[1], userInputs[2])
		messageResp, err := SearchForDegrees(userInputs[0], userInputs[1], userInputs[2])
		message := Urls(messageResp)
		userInputs = userInputs[:0]
		session["degrees"] = userInputs
		counter = 0
		session["degreesCounter"] = counter
		return message, err
	}
	return "", nil
}

// HandleSequence ... function
func HandleSequence(session models.Session, input string) (string, error) {
	var message string
	var err error

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
		choices := "What do you want to search for?<br>" +
			"1) Jobs & Internships (type jobs)<br>" +
			"2) Courses (type courses)<br>" +
			"3) Bachelor, Masters & PHD Degrees (type degrees)"
		return choices, nil
	case 1:
		switch strings.ToLower(input) {
		case "jobs":
			scenario = 0
			session["scenario"] = scenario
		case "courses":
			scenario = 1
			session["scenario"] = scenario
		case "degrees":
			scenario = 2
			session["scenario"] = scenario
		}

		switch scenario {
		case 0:
			message, err = HandleJobs(session, input)
		case 1:
			message, err = HandleCourses(session, input)
		case 2:
			message, err = HandleDegrees(session, input)
		}
	}
	return message, err
}
