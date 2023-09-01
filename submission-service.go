package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Fetching_Submitions_With_Ids(contest_id string) map[string][]SubmissionCode {
	solutions := make(map[string][]SubmissionCode)

	page_num := 1
	page_subms, exist := fetching_SubmissionID_Code(page_num, contest_id)

	for question, value := range page_subms {
		solutions[question] = append(solutions[question], value...)
	}
	for exist {
		page_num++
		page_subms, exist = fetching_SubmissionID_Code(page_num, contest_id)

		for question, value := range page_subms {
			solutions[question] = append(solutions[question], value...)
		}
	}

	return solutions
}

func fetching_SubmissionID_Code(page_num int, contest_id string) (map[string][]SubmissionCode, bool) {

	submission_code := make(map[string][]SubmissionCode)
	page := strconv.Itoa(page_num)
	req_link := "https://leetcode.com/contest/api/ranking/" + contest_id + "/?pagination=" + page + "&region=global"
	resp, err := http.Get(req_link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var d Contest
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		panic(err)
	}
	subms := d.Submissions
	for _, el := range subms {
		if len(el) == 0 {
			return submission_code, false
		}
		for question, e := range el {
			if e.DataRegion == "US" {
				var subid_code SubmissionCode
				id := strconv.Itoa(e.SubmissionID)
				sub := "https://leetcode.com/api/submissions/" + id
				resp, err := http.Get(sub)
				if err != nil {
					continue
				} else {
					defer resp.Body.Close()
					var submission Solution
					if err := json.NewDecoder(resp.Body).Decode(&submission); err != nil {
						panic(err)
					}
					subid_code.Code = submission.Code
				}
				subid_code.Submission_ID = e.ID
				submission_code[question] = append(submission_code[question], subid_code)
			}
		}
	}

	return submission_code, true
}
