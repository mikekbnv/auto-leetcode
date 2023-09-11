/*
Copyright © 2023 NAME HERE mkkbnv18@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/mikekbnv/auto-leetcode/internal/config"
	"github.com/spf13/cobra"
)


var (
	delay 	   int
	contest_id string
	cookie 	   string
)

func Auto_Leetcode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-leetcode",
		Short: "...",
		Long: `...`,
		
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is Auto Leetcode Reporter")
			//fmt.Println(config.LeetcodeConfig)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {			
			
			err := config.InitLeetcodeConfig()

			if err != nil {
				fmt.Println("Error initializing config")
			}
		},
		
	}
	
	cmd.AddCommand(Report())
	cmd.AddCommand(ConfigInit())

	cmd.PersistentFlags().IntVarP(&delay, "delay", "d", 500, "Time to wait between each request")
	cmd.PersistentFlags().StringVarP(&contest_id, "contest-id", "c", "", "Contest ID")
	cmd.PersistentFlags().StringVarP(&cookie, "cookie", "k", "", "Session cookie")

	return cmd
}

func InitConfigWithParams() {
	if contest_id != "" {
		config.SetConfigField("contest_id", contest_id)
	}

	if delay != 500 {
		config.SetConfigField("delay", delay)
	}

	if cookie != "" {
		parsedcookie := config.ParseCookies(cookie)
		config.SetConfigField("csrf_token", parsedcookie["csrftoken"])
		config.SetConfigField("jwt_token", parsedcookie["jwt_token"])				
	}
}


