package save

import (
	"fmt"
	"os"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/cli"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
)

func Save(content string) {
	file, err := os.OpenFile(cli.Opts.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		if cli.Opts.Verbose {
			logger.Stdlogger(fmt.Sprintf("Exception occured in the save module due to: %s", err.Error()), "warn")
		}
		return
	}
	defer file.Close()
	if _, err = file.WriteString(content + "\n"); err != nil {
		if cli.Opts.Verbose {
			logger.Stdlogger(fmt.Sprintf("Exception occured in the save module when writing output due to: %s", err.Error()), "warn")
		}
		return
	}
}
