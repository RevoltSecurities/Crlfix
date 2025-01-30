package logger

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
)

const (
	tcl = "\r\x1b[2K"
)

func Logger(message string, level string) { //used for mostly helps and logs before executions
	var leveler aurora.Value
	if level == "info" {
		leveler = aurora.Bold(aurora.Blue("INFO"))
	} else if level == "warn" {
		leveler = aurora.Bold(aurora.Yellow("WRN"))
	} else if level == "verbose" {
		leveler = aurora.Bold(aurora.Green("VERBOSE"))
	} else if level == "error" {
		leveler = aurora.Bold(aurora.Red("ERROR"))

	} else {
		leveler = aurora.Bold(aurora.Blue(level))
	}
	formatted := fmt.Sprintf("%s%s%s %s\n",
		aurora.Bold(aurora.White("[")),
		leveler,
		aurora.Bold(aurora.White("]: ")),
		aurora.Bold(aurora.White(message)),
	)
	fmt.Fprintf(os.Stderr, "%s", formatted)
}

func Bolder(message string) string {
	msg := fmt.Sprintf(`%s`, aurora.Bold(aurora.White(message)))
	return msg
}

func Loader(message string, level string) string {
	var leveler aurora.Value
	if level == "info" {
		leveler = aurora.Bold(aurora.Blue("INFO"))
	} else if level == "warn" {
		leveler = aurora.Bold(aurora.Red("WRN"))
	} else if level == "verbose" {
		leveler = aurora.Bold(aurora.Green("VERBOSE"))
	} else {
		leveler = aurora.Bold(aurora.Blue(level))
	}
	formatted := fmt.Sprintf("%s%s%s %s\n",
		aurora.Bold(aurora.White("[")),
		leveler,
		aurora.Bold(aurora.White("]: ")),
		aurora.Bold(aurora.White(message)),
	)
	return formatted
}

func Stdlogger(message string, level string) { // used for after execution to log messages so doesn't interfere with progress bar
	if level == "info" {
		fmt.Fprintf(os.Stderr, "%s", tcl)
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s\n", tcl, aurora.Bold(aurora.White("[")), aurora.Bold(aurora.Blue("INFO")), aurora.Bold(aurora.White("]: ")), aurora.Bold(aurora.White(message)))
	} else if level == "warn" {
		fmt.Fprintf(os.Stderr, "%s", tcl)
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s\n", tcl, aurora.Bold(aurora.White("[")), aurora.Bold(aurora.Red("WRN")), aurora.Bold(aurora.White("]: ")), aurora.Bold(aurora.White(message)))
	} else if level == "error" {
		fmt.Fprintf(os.Stderr, "%s", tcl)
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s\n", tcl, aurora.Bold(aurora.White("[")), aurora.Bold(aurora.Red("ERR")), aurora.Bold(aurora.White("]: ")), aurora.Bold(aurora.White(message)))
	} else if level == "vuln" {
		fmt.Fprintf(os.Stderr, "%s", tcl)
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s\n", tcl, aurora.Bold(aurora.White("[")), aurora.Bold(aurora.Green("VULN")), aurora.Bold(aurora.White("]: ")), aurora.Bold(aurora.White(message)))
	} else {
		fmt.Fprintf(os.Stderr, "%s", tcl)
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s\n", tcl, aurora.Bold(aurora.White("[")), aurora.Bold(aurora.Blue(level)), aurora.Bold(aurora.White("]: ")), aurora.Bold(aurora.White(message)))
	}
}

func Vlogger(level string, toolname string, version string) {
	if level == "latest" {
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s%s\n",
			aurora.Bold(aurora.White("[")),
			aurora.Bold(aurora.Blue("Version")),
			aurora.Bold(aurora.White("]: ")),
			aurora.Bold(aurora.White(fmt.Sprintf("%s current version %s (", toolname, version))),
			aurora.Bold(aurora.Green("latest")),
			aurora.Bold(aurora.White(")")),
		)
	} else {
		fmt.Fprintf(os.Stderr, "%s%s%s%s%s%s\n",
			aurora.Bold(aurora.White("[")),
			aurora.Bold(aurora.Blue("Version")),
			aurora.Bold(aurora.White("]: ")),
			aurora.Bold(aurora.White(fmt.Sprintf("%s current version %s (", toolname, version))),
			aurora.Bold(aurora.Red("outdated")),
			aurora.Bold(aurora.White(")")),
		)
	}
}

func Bannerizer(banner string, color string) aurora.Value {
	if color == "blue" {
		return aurora.Bold(aurora.Blue(banner))
	} else if color == "white" {
		return aurora.Bold(aurora.White(banner))
	} else if color == "magenta" {
		return aurora.Bold(aurora.Magenta(banner))
	} else if color == "green" {
		return aurora.Bold(aurora.Green(banner))
	} else if color == "cyan" {
		return aurora.Bold(aurora.Cyan(banner))
	} else {
		return aurora.Bold(aurora.Blue(banner))
	}
}
