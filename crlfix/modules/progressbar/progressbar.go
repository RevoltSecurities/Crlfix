package progressbar

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	tcl = "\r\x1b[2K"
)

type Progressbar struct {
	ReqCount   int
	ReqTotal   int
	StartedAt  time.Time
	ErrorCount int
	Race       sync.Mutex
}

func (pb *Progressbar) Bar(reqincr, errincr, ttlincr int) {
	pb.Race.Lock()
	defer pb.Race.Unlock()
	pb.ReqCount += reqincr
	pb.ErrorCount += errincr
	pb.ReqTotal += ttlincr
}

func Render(stats *Progressbar) {
	percentage := float64(stats.ReqCount) / float64(stats.ReqTotal) * 100                 //calculate percentage for porgressing tasks here
	blockcounts := int(percentage / 5)                                                    // we calculate this for block to be added in progressbar
	ProgressBar := strings.Repeat("█", blockcounts) + strings.Repeat(" ", 20-blockcounts) //adding the strings with empty space
	runningseconds := time.Since(stats.StartedAt).Seconds()
	rate := float64(stats.ReqCount) / runningseconds // request speed calculation done here so we see how much concurrent rate request are going
	hours := int(runningseconds) / 3600
	minutes := (int(runningseconds) % 3600) / 60
	seconds := int(runningseconds) % 60
	Printer(ProgressBar, percentage, hours, minutes, seconds, stats.ReqCount, stats.ReqTotal, rate, stats.ErrorCount)
}

func Printer(progressbar string, percentage float64, hours int, minutes int, seconds int, reqcounts int, totalreq int, rate float64, errors int) {
	fmt.Fprintf(os.Stderr, "%s[%s] %6.2f%% [%02d:%02d:%02d] [%d/%d] %.1f req/sec Errors: %d\r",
		tcl, progressbar, percentage, hours, minutes, seconds, reqcounts, totalreq, rate, errors)
}
