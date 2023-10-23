/*
Copyright © 2023 yeomc <yjc720@gmail.com>
*/
package cmd

import (
	"github.com/yeom-c/golang-fiber-api/tool/cmd/generate"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "controller, service, repository, model 생성",
	Long: `controller, service, repository, model 을 생성합니다.
예: tool.go generate user,user_mail`,
	Run: generate.Run,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	generateCmd.Flags().StringP("controller", "c", "", "controller 생성 | 예: -c user,user_mail")
	generateCmd.Flags().StringP("service", "s", "", "service 생성 | 예: -s user,user_mail")
	generateCmd.Flags().StringP("repository", "r", "", "repository 생성 | 예: -r user,user_mail")
	generateCmd.Flags().StringP("model", "m", "", "model entity 생성 | 예: -m user,user_mail")
}
