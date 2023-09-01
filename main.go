package main

import (
	"fmt"
)

func main() {

	//Config serivce - Func Read from file, Func Set(key, value), Get config
	_ = CreateLeetcodeConfig()
	fmt.Println(LeetcodeConfig)
	submissiondata := Fetching_Submitions_With_Ids(LeetcodeConfig.ContestID)
	fmt.Println(submissiondata)

	//Detect, group similarity submissions

	Report([]int{}, "", "Code from stackoverflow")

}

//Fetching submitions for checking service
//Checker service

//TODO:
//leetcode-reporter --cookie "cookie value" --delay 10 --contest-id "bi/weekly-contest-358"
//detector-api.go

//config-service.go - sengelton-service
