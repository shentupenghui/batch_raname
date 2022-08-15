package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	//RootCmd definition of batch-rename cli
	RootCmd = &cobra.Command{
		Use:   "batch-rename",
		Short: "A tool for batch rename(replace) files in a certain directory",
		Long:  `This tool can rename or delete filename in a directory.`,
		Example: `'batch-rename dir to_be_deleted'         will delete to_be_deleted str in dir/ all files
'batch-rename dir to_be_replace replace' will replace dir/ all filename contains to_be_replace with replace str`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			basePath := args[0]
			fileInfo, err := os.Stat(basePath)
			if err != nil {
				fmt.Printf("'%s' is not a valid direcotry, please input a valid directory", basePath)
				return
			}
			if !fileInfo.IsDir() {
				fmt.Printf("'%s' is not a direcotry, please input a valid directory", basePath)
				return
			}
			files, _ := os.ReadDir(basePath)
			//delete str
			if len(args) == 2 {
				for _, entry := range files {
					oldName := entry.Name()
					newName := strings.ReplaceAll(oldName, args[1], "")
					oldPath := filepath.Join(basePath, oldName)
					newPath := filepath.Join(basePath, newName)
					fmt.Printf("filename before delete: %s, filename after delete: %s", oldPath, newPath)
					os.Rename(oldPath, newPath)
				}
			}
			if len(args) == 3 {
				for _, entry := range files {
					oldName := entry.Name()
					tmpName := strings.ReplaceAll(oldName, args[1], "")
					newName := strings.ReplaceAll(tmpName, args[1], args[2])
					oldPath := filepath.Join(basePath, oldName)
					newPath := filepath.Join(basePath, newName)
					fmt.Printf("filename before replace: %s, filename after replace: %s", oldPath, newPath)
					os.Rename(oldPath, newPath)
				}
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}
