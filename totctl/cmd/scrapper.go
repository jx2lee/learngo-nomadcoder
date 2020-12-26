/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/jx2lee/learngo-nomadcoder/totctl/pkg/scrapper"
	"github.com/spf13/cobra"
)

// scrapperCmd represents the scrapper command
var scrapperCmd = &cobra.Command{
	Use:   "scrapper",
	Short: "#4 JOB SCRAPPER",
}

var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "A single scrapper",
	Run: func(cmd *cobra.Command, args []string) {
		scrapper.Scrape("python")
	},
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "A web Scrapper using Echo",
	Run: func(cmd *cobra.Command, args []string) {
		scrapper.ScrapperEcho()
	},
}

func init() {
	rootCmd.AddCommand(scrapperCmd)
	scrapperCmd.AddCommand(singleCmd)
	scrapperCmd.AddCommand(webCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
