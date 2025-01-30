package cli

import (
	"fmt"
	"os"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/config"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/help"
	"github.com/spf13/cobra"
)

type Argsparser struct {
	Url         string
	List        string
	Verbose     bool
	Concurrency int
	Headers     []string
	Method      string
	Config      string
	Proxy       string
	No_Color    bool
	Output      string
	Redirect    bool
	Maxr        int
	Exceptions  error
	Timeouts    int
	Randomagent bool
	Notify      bool
	SlackURL    string
	Ratelimit   int
	Silent      bool
	Delay       int
}

var Opts = &Argsparser{}
var rootCmd = &cobra.Command{
	Use: "crlfix",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func customHelp(cmd *cobra.Command, args []string) {
	helpMessage := help.Helper()
	fmt.Println(helpMessage)
	os.Exit(0)
}

func Execute() Argsparser {
	err := rootCmd.Execute()
	if err != nil {
		return Argsparser{Exceptions: err}
	}
	if !rootCmd.Flags().Changed("config") {
		defaultconfig, err := config.GetConfig()
		if err != nil {
			Opts.Config = ""
		} else {
			Opts.Config = defaultconfig
		}
	}
	return *Opts
}

func init() {
	rootCmd.SetHelpFunc(customHelp)

	rootCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		customHelp(cmd, nil)
		return nil
	})

	rootCmd.Flags().StringVarP(&Opts.Url, "url", "u", "", "url for making http request and print response code and others")
	rootCmd.Flags().StringVarP(&Opts.List, "list", "l", "", "file that contains urls")
	rootCmd.Flags().BoolVarP(&Opts.Verbose, "verbose", "v", false, "Verbose to print failed request errors and reason")
	rootCmd.Flags().IntVarP(&Opts.Concurrency, "concurrency", "c", 10, "Concurrency")
	rootCmd.Flags().StringSliceVarP(&Opts.Headers, "headers", "H", nil, "Headers to pass")
	rootCmd.Flags().StringVarP(&Opts.Method, "method", "X", "get", "Request method")
	rootCmd.Flags().StringVarP(&Opts.Config, "config", "C", "", "config-file path")
	rootCmd.Flags().StringVarP(&Opts.Output, "save", "O", "", "output file to write output")
	rootCmd.Flags().StringVarP(&Opts.Proxy, "proxy", "p", "", "proxy url")
	rootCmd.Flags().BoolVarP(&Opts.Redirect, "follow-redirect", "r", false, "Verbose to print failed request errors and reason")
	rootCmd.Flags().IntVarP(&Opts.Maxr, "max-redirect", "M", 20, "Max redirections to follow")
	rootCmd.Flags().IntVarP(&Opts.Timeouts, "timeout", "T", 20, "Max Timeout for requests")
	rootCmd.Flags().BoolVarP(&Opts.Randomagent, "random-agent", "R", false, "Random User agents for requests")
	rootCmd.Flags().BoolVarP(&Opts.Notify, "notify", "N", false, "Notification if crlf injection found")
	rootCmd.Flags().IntVarP(&Opts.Ratelimit, "rate-limit", "L", 150, "Specify the ratelimit for concurrent request")
	rootCmd.Flags().IntVarP(&Opts.Delay, "delay", "D", 0, "Specify the delay between each request")
	rootCmd.Flags().BoolVarP(&Opts.Silent, "silent", "s", false, "Enable to disable banner and version logging")

}