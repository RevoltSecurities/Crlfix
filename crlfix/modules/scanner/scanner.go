package scanner

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/cli"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/notify"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/progressbar"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/request"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/save"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/vulnerable"
	"github.com/projectdiscovery/ratelimit"
	"golang.org/x/sync/semaphore"
)

func Scanner(target string, urls []string, httpx *http.Client, args cli.Argsparser, lock *sync.Mutex, semaphore *semaphore.Weighted, wg *sync.WaitGroup) {
	logger.Stdlogger(fmt.Sprintf("Crlf Injection scanning started for: %s", target), "info")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bar := &progressbar.Progressbar{
		ReqTotal:  len(urls),
		StartedAt: time.Now(),
	}
	Ticker := time.NewTicker(1 * time.Second)
	defer Ticker.Stop()

	limiter := ratelimit.New(context.Background(), uint(args.Ratelimit), time.Duration(1*time.Second))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			logger.Stdlogger("CTRL+C Pressed!", "warn")
			os.Exit(1)
		}
	}()

	go func() {
		for range Ticker.C {
			lock.Lock()
			progressbar.Render(bar)
			lock.Unlock()
		}
	}()

	for _, url := range urls {
		wg.Add(1)
		if err := semaphore.Acquire(ctx, 1); err != nil {
			logger.Stdlogger(fmt.Sprintf("Unable to create a runner, exception occured in the semaphore acquiring due to: %s", err.Error()), "warn")
			os.Exit(1)
		}

		go func(url string) {
			defer semaphore.Release(1)
			defer wg.Done()
			limiter.Take()
			response, err := request.Request(url, args, httpx, lock)
			if err != nil {
				if args.Verbose {
					logger.Stdlogger(fmt.Sprintf("Request failed due to: %s", err.Error()), "warn")
				}

				lock.Lock()
				bar.Bar(1, 1, 0)
				lock.Unlock()
			}
			if response == nil {
				return
			}
			header, value, vuln := vulnerable.Vulnerable(response)
			if vuln {
				logger.Stdlogger(fmt.Sprintf("%s [%s:%s] [%d]", url, header, value, response.StatusCode), "vuln")

				if args.Output != "" {
					save.Save(fmt.Sprintf("%s [%s:%s] [%d]", url, header, value, response.StatusCode))
				}

				if args.Notify {
					notify.SendNotify(args.SlackURL, url, args, httpx)
				}
			} else {
				if args.Verbose {
					logger.Logger(fmt.Sprintf("%s is not vulnerable", url), "warn")
				}
			}
			defer response.Body.Close()
			lock.Lock()
			bar.Bar(1, 0, 0)
			lock.Unlock()
		}(url)
	}
	wg.Wait()
	progressbar.Render(bar)
	fmt.Println()
}
