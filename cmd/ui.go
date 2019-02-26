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
    "log"

	"github.com/spf13/cobra"
    ui "github.com/gizak/termui"
    widgets "github.com/gizak/termui/widgets"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Initialize the godo user interface",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        // Ensure the UI can be initialized
        if err := ui.Init(); err != nil {
            log.Fatalf("Failed to intialize termui: %v", err);
        }

        defer ui.Close();
       
        // Initialize a new paragraph
        p := widgets.NewParagraph()
        p.Text = "Hello world!"
	    p.SetRect(0, 0, 25, 5)
	    ui.Render(p)

        for e := range ui.PollEvents() {
            if e.Type == ui.KeyboardEvent {
                break
            }
        }
	    
	},
}

func init() {
	RootCmd.AddCommand(uiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
