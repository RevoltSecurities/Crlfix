package handler

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/banner"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/cli"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/client"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/config"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/payloads"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/reader"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/scanner"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/utils"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/gitversion"
	"golang.org/x/sync/semaphore"
)

var (
	wg   = sync.WaitGroup{}
	sem  *semaphore.Weighted
	lock sync.Mutex
	V = "v1.0.0"
)

func Version(){
	git, err := gitversion.GitVersion()
	if err != nil{
		logger.Stdlogger(fmt.Sprint("Unable to get the latest version of the crlfix"), "warn")
	}else{
		if git == V {
			logger.Vlogger("latest", "crlfix", V)
		}else{
			logger.Vlogger("outdated", "crlfix", V)
		}
	}
	
}

func Init() {

	args := cli.Execute()

	if !args.Silent {
		fmt.Fprintf(os.Stderr, "%s", banner.BannerGenerator("crlfix"))
		fmt.Fprintf(os.Stderr, "\n              - %s", logger.Bolder("Revoltsecurities"))
		fmt.Println("\n")
		Version()
	}

	logger.Stdlogger(fmt.Sprintf("Loading configuration file from %s", args.Config), "info")

	sem = semaphore.NewWeighted(int64(args.Concurrency))
	Payloads := payloads.Payloads()
	httpx, err := client.HttpxClient()
	if err != nil {
		logger.Logger(fmt.Sprintf("Exception occured in the building http client due to: %s", err.Error()), "error")
		os.Exit(1)
	}

	if args.Output != "" {
		ok,err := utils.IsPermission(args.Output)
		if err != nil && !ok{
			logger.Stdlogger(fmt.Sprintf("error occured in checking file permission due to: %s", err.Error()), "warn")
			os.Exit(1)
		}
	}

	if args.Notify {
		cfg, err := config.SetConfig(args.Config)
		if err != nil {
			if args.Verbose {
				logger.Logger(fmt.Sprintf("Unable to load configurations due to: %s", err.Error()), "warn")
			}
			args.SlackURL = ""
			return
		}

		randurl, err := cfg.GetRandomKey()
		if err != nil {
			if args.Verbose {
				logger.Logger(fmt.Sprintf("Unable to retrieve random Slack URL: %s", err.Error()), "warn")
			}
			args.SlackURL = ""
		}

		args.SlackURL = randurl
	} else {
		args.SlackURL = ""
	}

	if args.Url != "" {
		Start(args.Url, httpx, args, Payloads)
		os.Exit(0)
	}

	if args.List != "" {
		urls, err := reader.Reader(args.List)
		if err != nil {
			logger.Stdlogger(fmt.Sprintf("Exception occured in reading %s due to: %s", args.List, err.Error()), "error")
			os.Exit(1)
		}
		for _, url := range urls {
			Start(url, httpx, args, Payloads)
		}
		os.Exit(0)
	}

	if !utils.IsStdin() {
		logger.Stdlogger("no input passed to crlfix, please use crlfix -h for more information\n", "warn")
		os.Exit(1)
	} else {
		scanned := bufio.NewScanner(os.Stdin)
		for scanned.Scan() {
			url := scanned.Text()
			url = strings.TrimSpace(url)
			if url != ""{
				Start(url, httpx, args, Payloads)
			}
		}
		os.Exit(0)
	}

}

func Start(url string, httpx *http.Client, args cli.Argsparser, payloads []string) {
	var urls []string
	for _, payload := range payloads {

		u, err := utils.Pathadder(url, payload)
		if err != nil {
			continue
		}
		urls = append(urls, u)
	}
	scanner.Scanner(url, urls, httpx, args, &lock, sem, &wg)
}
