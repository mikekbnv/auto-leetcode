package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mikekbnv/auto-leetcode/internal/client"
	"github.com/mikekbnv/auto-leetcode/internal/config"
)

const (
	baseURL         = "https://leetcode.com/contest/api/ranking/"
	usSubmissionURL = "https://leetcode.com/api/submissions/"
	cnSubmissionURL = "https://leetcode.cn/api/submissions/"
)

var count int64
var httpClient = &http.Client{}
var mu sync.Mutex

func Cuncurrent_Fethcing_Submitions_With_Ids(contest_id string) map[string][]Solution {
	solutions := make(map[string][]Solution)
	startTime := time.Now()
	fmt.Println("Start fetching submissions...")

	numGoroutines := 20

	done := make(chan struct{})
	results := make(chan map[string][]Solution, numGoroutines)

	processPages := func(startPage, endPage int, results chan<- map[string][]Solution) {
		pageSubmissions := make(map[string][]Solution)
		defer func() {
			results <- pageSubmissions
		}()

		for pageNum := startPage; pageNum < endPage; pageNum++ {
			pageSubms, err := fetchingSubmissionIDCode(pageNum, contest_id)
			if err != nil {
				fmt.Println("Error fetching submissions:", err)
				continue
			}

			for question, value := range pageSubms {
				pageSubmissions[question] = append(pageSubmissions[question], value...)
			}
		}
	}

	pagesPerGoroutine := 10
	for i := 0; i < numGoroutines; i++ {
		startPage := i * pagesPerGoroutine
		endPage := (i + 1) * pagesPerGoroutine
		fmt.Println(startPage, endPage)
		time.Sleep(500 * time.Millisecond)
		go processPages(startPage, endPage, results)
	}

	// Collect results from goroutines
	go func() {
		for i := 0; i < numGoroutines; i++ {
			pageSubmissions := <-results
			for question, value := range pageSubmissions {
				solutions[question] = append(solutions[question], value...)
			}
		}
		close(done)
	}()

	<-done

	elapsedTime := time.Since(startTime)
	fmt.Println("Fetching submissions took:", elapsedTime, "count:", count)
	return solutions
}

func fetchingSubmissionIDCode(pageNum int, contestID string) (map[string][]Solution, error) {
	fetchedsubmissions := make(map[string][]Solution)
	page := strconv.Itoa(pageNum)
	reqLink := baseURL + contestID + "/?pagination=" + page + "&region=global"
	fmt.Println(reqLink)
	resp, err := httpClient.Get(reqLink)
	if err != nil {
		return fetchedsubmissions, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fetchedsubmissions, fmt.Errorf("Received non-OK status code: %s", resp.Status)
	}

	var contestData Contest
	if err := json.NewDecoder(resp.Body).Decode(&contestData); err != nil {
		return fetchedsubmissions, err
	}

	submissions := contestData.Submissions
	if len(submissions) == 0 {
		return fetchedsubmissions, errors.New("No submissions found")
	}

	var wg sync.WaitGroup
	wg.Add(len(submissions))

	for _, submissionList := range submissions {
		time.Sleep(1 * time.Second)
		go func(submissionList map[string]Question) {
			defer wg.Done()

			for question, submission := range submissionList {
				submissionCode, err := fetchSubmissionCode(submission.DataRegion, submission.SubmissionID)
				if err != nil {
					fmt.Printf("Error fetching submission code for question %s: %v\n", question, err)
					continue
				}

				mu.Lock()
				fetchedsubmissions[question] = append(fetchedsubmissions[question], submissionCode)
				mu.Unlock()
			}
		}(submissionList)
	}

	wg.Wait()
	return fetchedsubmissions, nil
}

func fetchSubmissionCode(region string, sub_id int) (Solution, error) {
	var submissionData Solution
	id := strconv.Itoa(sub_id)
	subURL := usSubmissionURL
	if region != "US" {
		time.Sleep(500 * time.Millisecond)
		subURL = cnSubmissionURL
		//return Solution{}, errors.New("not US region")
	}
	subURL += id

	time.Sleep(500 * time.Millisecond)
	resp, err := httpClient.Get(subURL)
	if err != nil {
		return submissionData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return submissionData, fmt.Errorf("received non-OK status code: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&submissionData); err != nil {
		return submissionData, err
	}

	atomic.AddInt64(&count, 1)
	return submissionData, nil
}

// Taking any submission code by submission id, since for contest there is api but for regular submissions need to parse html
func FetchSubmissionCode() (string, error) {
	url := "https://leetcode.com/submissions/detail/1038989061/"
	client := client.NewLeetcodeHttpClient(config.LeetcodeConfig.CSRFToken, config.LeetcodeConfig.JWTToken)

	resp, err := client.Get(url, "application/json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK status code: %s", resp.Status)
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)

	pattern := `submissionCode:\s('[^']*')`

	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(body)

	if len(match) != 2 {
		return "", fmt.Errorf("no submission code found in the response")
	}

	submissionCode := match[1]
	fmt.Println(convertUnicodeEscape(submissionCode))
	return "", nil
}

func convertUnicodeEscape(input string) (string, error) {
	re := regexp.MustCompile(`\\u[0-9A-Fa-f]{4}`)
	result := re.ReplaceAllStringFunc(input, func(matched string) string {
		code := matched[2:]
		unicodeValue := stringToUnicode(code)
		return unicodeValue
	})

	return result, nil
}
func stringToUnicode(s string) string {
	unicodeValue, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		return s
	}
	return string(unicodeValue)
}
