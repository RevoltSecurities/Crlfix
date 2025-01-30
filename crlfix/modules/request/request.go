package request

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/cli"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
)

func Request(url string, args cli.Argsparser, httpx *http.Client, mu *sync.Mutex) (*http.Response, error) {

	request, err := http.NewRequest(strings.ToUpper(args.Method), url, nil)
	if err != nil {
		return nil, err
	}

	for _, headers := range args.Headers {
		keys := strings.SplitN(headers, ":", 2)
		if (len(keys)) != 2 {
			if args.Verbose {
				logger.Stdlogger(fmt.Sprintf("Skipping %s from setting Request Headers due to unexpected headers format", keys), "error")
			}
			continue
		}
		request.Header.Set(keys[0], keys[1])
	}

	if args.Randomagent {
		mu.Lock()
		request.Header.Set("User-Agent", browser.Random())
		mu.Unlock()
	}
	time.Sleep(time.Duration(args.Delay))
	response, err := httpx.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
