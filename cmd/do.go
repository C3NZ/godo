// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/C3NZ/godo/db"
	"github.com/spf13/cobra"
	"strconv"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			// Attempt to convert the argument from a str to an integer
			id, err := strconv.Atoi(arg)

			// If there is an error, inform the user.
			// Append to the list of ids otherwise
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				fmt.Println("You've completed task:", id)
				db.DeleteTask(id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
