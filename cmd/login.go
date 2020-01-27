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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type Context struct {
	host      string
	authToken string
}

const GITHUB_AUTH = "https://api.github.com/user"

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to configrd",
	Run: func(cmd *cobra.Command, args []string) {

		host, _ := cmd.Flags().GetString("host")
		account, _ := cmd.Flags().GetString("account")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		token, _ := cmd.Flags().GetString("token")

		if strings.Contains(host, "{account}") && account != "" {
			h := &host
			*h = strings.Replace(host, "{account}", account, 1)
		}

		if dir, err := homedir.Dir(); err == nil {
			if exp, err := homedir.Expand(dir); err == nil {
				os.Mkdir(path.Join(exp, ".cfgrd"), os.ModeDir)
			}
		}

		context := &Context{
			host:      host,
			authToken: "",
		}

		if s, err := context.validate(); err == nil && s {
			req, err := http.NewRequest("GET", GITHUB_AUTH, nil)

			if err != nil {
				log.Fatal(err)
			}

			client := &http.Client{}
			if username != "" && password != "" {
				req.SetBasicAuth(username, password)
			} else if username != "" && token != "" {
				req.SetBasicAuth(username + ":" + token, nil)
			}
		}

		fmt.Println("Host:" + host)
		fmt.Println("login called")
	},
}

func (c Context) validate() (bool, error) {
	if c.host != "" {

		_, err := url.Parse(c.host)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().String("host", "https://{account}.api.configrd.io/", "configrd hostname to authenticate into")
	loginCmd.Flags().String("account", "", "your configrd account name (required)")
	loginCmd.MarkFlagRequired("account")
	loginCmd.MarkFlagRequired("host")

	loginCmd.Flags().String("token", "", "token")
	loginCmd.Flags().String("username", "", "username")
	loginCmd.Flags().String("password", "", "password")
}
