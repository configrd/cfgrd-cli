/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"log"
	"net/url"
)

type get struct {
	uri     *url.URL
	profile string
	cascade bool
	format  string
	host    string
}

// cfgrd env list -r default -f text -t true /env/dev/default.properties
// cfgrd env list -r default -p dev -f text
// cfgrd env list /env/dev/
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List environment variables and secrets",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		getParams := get{}
		getParams.cascade, _ = cmd.Flags().GetBool("cascade")
		getParams.format, _ = cmd.Flags().GetString("format")
		getParams.profile, _ = cmd.Flags().GetString("profile")
		getParams.host, _ = cmd.PersistentFlags().GetString("host")

		if len(args) == 1 {
			uri, err := url.Parse(args[0])
			if err != nil {
				log.Fatal(err)
			}
			getParams.uri = uri
		} else if len(args) == 0 {
			uri, err := url.Parse("/")
			if err != nil {
				log.Fatal(err)
			}
			getParams.uri = uri
		} else {
			log.Fatal("Invalid number of arguments. Expected 0 or 1: uri")
		}
		fmt.Println("host "+getParams.host)
	},
}

func init() {
	envCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("profile", "p", "", "named configuration profile")
	listCmd.Flags().BoolP("cascade", "c", true, "whether to traverse parent paths")
	listCmd.Flags().StringP("format", "f", "text", "response format text, json, yaml")
}
