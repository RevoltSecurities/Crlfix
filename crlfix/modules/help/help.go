package help

import (
	"fmt"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/banner"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
)

func Helper() string {
	msg := fmt.Sprintf(`
%s

                    %s
           
%s

%s
    %s

%s

    %s
        %s

    %s
        %s

    %s
        %s

    %s
        %s

    %s
        %s

    %s
        %s

    %s
        %s
    
    %s
        %s
`,
		banner.BannerGenerator("crlfiX"),
		logger.Bolder("- RevoltSecurities"),
		logger.Loader("An accurate and concurrent CRLF injection scanner", "DESCRIPTION"),
		logger.Loader("", "USAGE"),
		logger.Bolder("crlfix [flags]"),
		logger.Loader("", "FLAGS"),
		logger.Loader("", "INPUT"),
		logger.Bolder(`-u,  --url                      :  Target URL for CRLF injection scan
        -l,  --list                     :  File containing a list of URLs for CRLF injection scan
        stdin/stdout                    :  crlfix also supports stdin/stdout for URL input/output.`),
		logger.Loader("", "OUTPUT"),
		logger.Bolder(`-O,  --save                     :  File to save the results of the scan.`),
		logger.Loader("", "CONCURRENCY"),
		logger.Bolder(`-c,  --concurrency              :  Set concurrency level (default: 10).
        -L,  --rate-limit               :  Specify the rate limit for concurrent requests.`),
		logger.Loader("", "OPTIMIZATION"),
		logger.Bolder(`-T,  --timeout                  :  Maximum timeout for requests (default: 20 seconds).
        -D,  --delay                    :  Delay (in milliseconds) between each request.
        -R,  --random-agent             :  Use random User-Agent headers for requests.`),
		logger.Loader("", "CONFIGURATION"),
		logger.Bolder(`-C,  --config                   :  Path to the custom configuration file of crlfix.
        -H,  --headers                  :  Headers to include in the HTTP requests (e.g., -H cookie:values, -H X-Orgin-Host:127.0.0.1).`),
		logger.Loader("", "NETWORK"),
		logger.Bolder(`-p,  --proxy                    :  Proxy URL for requests.
        -r,  --follow-redirect          :  Follow HTTP redirections (default: false).
        -M,  --max-redirect             :  Maximum number of redirects to follow (default: 20).`),
		logger.Loader("", "NOTIFICATIONS"),
		logger.Bolder(`-N,  --notify                   :  Enable notifications if CRLF injection is found.`),
		logger.Loader("", "DEBUGGING"),
		logger.Bolder(`-v,  --verbose                  :  Print errors and reasons for failed requests and other informations.
        -s,  --silent                   :  Disable banner and version logging.`),
	)
	return msg
}
