/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/stkali/utility/errors"
	"github.com/stkali/utility/tool"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	Version   string
	License   string
	Build     string
	Author    string
	CommitID  string
	ChangeLog string
)

var versionInfo = map[string]string{
	"name":       Program,
	"version":    Version,
	"build":      Build,
	"license":    License,
	"author":     Author,
	"commit_id":  CommitID,
	"change_log": ChangeLog,
}

func Json() string {
	data, err := json.Marshal(versionInfo)
	errors.CheckErr(err)
	return tool.ToString(data)
}

func VersionString() string {
	format := `%s %s build on %s
-----------------------------------------
  License  : %s
  Author   : %s
  CommitID : %s

ChangeLog:
-----------------------------------------
%s
`
	return fmt.Sprintf(format, Program, Version, Build, License, Author, CommitID, ChangeLog)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		info := map[string]string{
			"name":    Program,
			"version": Version,
			"build":   Build,
		}
		if cmd.Flag("detail").Changed {
			info["license"] = License
			info["author"] = Author
			info["commit_id"] = CommitID
			info["change_log"] = ChangeLog
		}
		jsonFlag, formatFlag := cmd.Flag("json").Changed, cmd.Flag("format").Changed
		if jsonFlag && formatFlag {
			errors.CheckErr("--json and --format cannot be used together")
		}
		if jsonFlag {
			data, err := json.Marshal(versionInfo)
			errors.CheckErr(err)
			fmt.Println(tool.ToString(data))
			return
		}
		if formatFlag {
			format := cmd.Flag("format").Value.String()
			if format[len(format)-1] != '\n' {
				format += "\n"
			}
			tmpl, err := template.New("_").Parse(format)
			errors.CheckErr(err)
			err = tmpl.Execute(os.Stdout, versionInfo)
			errors.CheckErr(err)
		}

		fmt.Println(VersionString())
	},
}

func init() {
	versionCmd.Flags().Bool("json", false, "Print json string")
	versionCmd.Flags().String("format", "", "Specify print format")
	versionCmd.Flags().String("detail", "", "Print detail information")
	rootCmd.AddCommand(versionCmd)
}
