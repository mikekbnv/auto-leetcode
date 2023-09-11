/*
Copyright © 2023 NAME HERE mkkbnv18@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/mikekbnv/auto-leetcode/internal/config"
	"github.com/mikekbnv/auto-leetcode/internal/report"
	"github.com/spf13/cobra"
)

// reportCmd represents the report command

func Report() *cobra.Command {
	report := &cobra.Command{
		Use:   "report",
		Short: "Auto report command",
		Long:  `...`,
		Run: func(cmd *cobra.Command, args []string) {
			if config.LeetcodeConfig.CSRFToken == "" || config.LeetcodeConfig.JWTToken == "" {
				fmt.Println("CSRFTOKEN or JWTTOKEN cannot be empty")
				return
			}

			//submissiondata := report.Fetching_Submitions_With_Ids(config.LeetcodeConfig.ContestID)
			submissions := report.Cuncurrent_Fethcing_Submitions_With_Ids(config.LeetcodeConfig.ContestID)

			for k := range submissions {
				fmt.Println(k, len(submissions[k]))
			}
			// for _, submission := range submissions {
			// 	//fmt.Println(submission[0])
			// 	break
			// }
			// Detect, group similarity submissions
			//submission_ids := report.Detecting_Similarity(submissions)
			// Report([]int{}, "", "Code from stackoverflow")
		},
	}

	return report
}
