/*
Copyright © 2023 NAME HERE mkkbnv18@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/mikekbnv/auto-leetcode/internal/config"
	"github.com/spf13/cobra"
)

func ConfigInit() *cobra.Command {
	config := &cobra.Command{
		Use:   "config",
		Short: "Initialize the configuration",
		Long:  `...`,
		Run: func(cmd *cobra.Command, args []string) {
			InitConfigWithParams()
			fmt.Println("Config initialized successfully", config.LeetcodeConfig)
		},
	}

	return config
}
