package main

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "autoCommit",
		Short: "自动提交工具",
		Run:   run,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	// 选择使用哪种AI
	seleceModel := promptui.Select{
		Label: "请选择 AI 模型",
		Items: []string{"豆包", "ChatGPT"},
	}
	seleceModel.Run()

	// 选择使用哪种模式
	selectMode := promptui.Select{
		Label: "请选择输出模式",
		Items: []string{"简单的描述", "详细的描述"},
	}

	selectMode.Run()

	s := spinner.New(spinner.CharSets[21],
		100*time.Millisecond,
		spinner.WithColor("green"))

	s.Start()

	// 模拟API调用
	callAPI()

	s.Stop()

	fmt.Println("🎉处理完成!")
}

func callAPI() string {
	time.Sleep(5 * time.Second)
	return "API调用成功"
}
