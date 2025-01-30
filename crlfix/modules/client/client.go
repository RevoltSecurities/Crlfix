package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/cli"
	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
)

func HttpxClient() (*http.Client, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:    50,
		IdleConnTimeout: 90 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 60,
			KeepAlive: time.Second * 80,
		}).DialContext,
	}

	if cli.Opts.Proxy != "" {
		proxyurl, err := url.Parse(cli.Opts.Proxy)
		if err != nil {
			if cli.Opts.Verbose{
			logger.Logger(fmt.Sprintf("Exception occured in configuring proxy due to: %s", err.Error()), "warn")
			}
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxyurl)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(cli.Opts.Timeouts) * time.Second,
	}

	if cli.Opts.Redirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) >= cli.Opts.Maxr {
				return http.ErrUseLastResponse
			}
			return nil
		}
	} else {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return client, nil
}
