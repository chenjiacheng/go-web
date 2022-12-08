package cmd

import (
	"github.com/spf13/cobra"
	"go-web/pkg/console"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	console.Success("---------")
	console.Error("---------")
	console.Warning("---------")
	console.Exit("---------")
}
