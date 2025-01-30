package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/cli"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
)

type Payload struct {
	Text string `json:"text"`
}

func SendNotify(url string, target string ,args cli.Argsparser, httpx *http.Client) {

	if args.SlackURL == "" {
		return
	}

	paytext := Payload{Text: fmt.Sprintf("Crlf Injection Found ⚡\n%s", target)}
	paybytes, err := json.Marshal(paytext)
	if err != nil {
		if args.Verbose {
			logger.Stdlogger("Json Marshalling exception occured", "warn")
		}
		return
	}

	request, err := http.NewRequest("POST", args.SlackURL, bytes.NewBuffer(paybytes))
	if err != nil {
		if args.Verbose {
			logger.Stdlogger(fmt.Sprintf("Exception occured in the noitify request building module due to: %s", err.Error()), "warn")
		}
		return
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := httpx.Do(request)
	if err != nil {
		if args.Verbose {
			logger.Stdlogger(fmt.Sprintf("Exception occured in the notify request due to: %s", err), "warn")
		}
		return
	}
	defer response.Body.Close()
}
