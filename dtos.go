package main

type Question struct {
	ID           int    `json:"id"`
	Date         int    `json:"date"`
	QuestionID   int    `json:"question_id"`
	SubmissionID int    `json:"submission_id"`
	Status       int    `json:"status"`
	ContestID    int    `json:"contest_id"`
	DataRegion   string `json:"data_region"`
	FailCount    int    `json:"fail_count"`
}

type Contest struct {
	Submissions []map[string]Question `json:"submissions"`
	Questions   []struct {
		ID         int    `json:"id"`
		QuestionID int    `json:"question_id"`
		Credit     int    `json:"credit"`
		Title      string `json:"title"`
		TitleSlug  string `json:"title_slug"`
	} `json:"questions"`
	TotalRank []struct {
		ContestID     int    `json:"contest_id"`
		Username      string `json:"username"`
		CountryCode   string `json:"country_code"`
		CountryName   string `json:"country_name"`
		Rank          int    `json:"rank"`
		Score         int    `json:"score"`
		FinishTime    int    `json:"finish_time"`
		GlobalRanking int    `json:"global_ranking"`
		DataRegion    string `json:"data_region"`
	} `json:"total_rank"`
	UserNum int `json:"user_num"`
}

type Solution struct {
	ID                int    `json:"id"`
	Code              string `json:"code"`
	Lang              string `json:"lang"`
	ContestSubmission int    `json:"contest_submission"`
}

type SubmissionCode struct {
	Submission_ID int
	Code          string
}
