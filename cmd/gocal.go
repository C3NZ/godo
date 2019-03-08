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
	gocal "github.com/Connor-Cahill/GoCal/cmd"
	calevents "github.com/connor-cahill/goCal/services"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var gocalCmd = &cobra.Command{
	Use:   "gocal",
	Short: "utilize the gocal integration",
	Run: func(cmd *cobra.Command, args []string) {

		calevents.MakeIDMap()

		calevents.UpdateMap()
		gocal.RootCmd.Execute()
	},
}

func init() {
	RootCmd.AddCommand(gocalCmd)
}
