# Crlfix - An accurate and concurrent CRLF Injection Vulnerability Scanner

**Fast, Accurate, and Concurrent CRLF Injection Scanner**  

![GitHub last commit](https://img.shields.io/github/last-commit/RevoltSecurities/Crlfix) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/RevoltSecurities/Crlfix) [![GitHub license](https://img.shields.io/github/license/RevoltSecurities/Subprober)](https://github.com/RevoltSecurities/Crlfix/blob/main/LICENSE) 

## Overview:  
Crlfix is an advanced CRLF injection vulnerability scanner designed for penetration testers and security researchers. It efficiently detects CRLF injection points in web applications while offering high concurrency and customization options.  

## Features:

✅ **Accurate Scanning** – Detects CRLF injection points effectively.  
✅ **High Concurrency** – Supports multiple concurrent scans with adjustable rate limits.  
✅ **Custom Headers & Configuration** – Allows setting HTTP headers and custom configurations.  
✅ **Proxy Support** – Enables scanning through proxies for anonymity.  
✅ **Optimized Performance** – Adjustable timeout, delay, and random User-Agent support.  
✅ **Notification System** – Alerts when a CRLF injection is detected.  
✅ **Input Flexibility** – Supports single URLs, list files, and stdin/stdout input.  

## Installation:  

### Using Go:  
```sh
go install -v github.com/RevoltSecurities/Crlfix/crlfix@latest
```

### Manual Installation  
1. Clone the repository:  
   ```sh
   git clone https://github.com/RevoltSecurities/Crlfix.git && cd Crlfix
   ```
2. Build the binary:  
   ```sh
   go build -o crlfix
   ```
3. Run the tool:  
   ```sh
   ./crlfix -h
   ```

## Usage:

```sh
crlfix -h
```

```yaml

                |    _|  _)  \ \  /
   __|    __|   |   |     |   \  /
  (      |      |   __|   |      \
 \___|  _|     _|  _|    _|   _/\_\


                    - RevoltSecurities
           
[DESCRIPTION]:  An accurate and concurrent CRLF injection scanner


[USAGE]:  

    crlfix [flags]

[FLAGS]:  


    [INPUT]:  

        -u,  --url                      :  Target URL for CRLF injection scan
        -l,  --list                     :  File containing a list of URLs for CRLF injection scan
        stdin/stdout                    :  crlfix also supports stdin/stdout for URL input/output.

    [OUTPUT]:  

        -O,  --save                     :  File to save the results of the scan.

    [CONCURRENCY]:  

        -c,  --concurrency              :  Set concurrency level (default: 10).
        -L,  --rate-limit               :  Specify the rate limit for concurrent requests.

    [OPTIMIZATION]:  

        -T,  --timeout                  :  Maximum timeout for requests (default: 20 seconds).
        -D,  --delay                    :  Delay (in milliseconds) between each request.
        -R,  --random-agent             :  Use random User-Agent headers for requests.

    [CONFIGURATION]:  

        -C,  --config                   :  Path to the custom configuration file of crlfix.
        -H,  --headers                  :  Headers to include in the HTTP requests (e.g., -H cookie:values, -H X-Orgin-Host:127.0.0.1).

    [NETWORK]:  

        -p,  --proxy                    :  Proxy URL for requests.
        -r,  --follow-redirect          :  Follow HTTP redirections (default: false).
        -M,  --max-redirect             :  Maximum number of redirects to follow (default: 20).

    [NOTIFICATIONS]:  

        -N,  --notify                   :  Enable notifications if CRLF injection is found.
    
    [DEBUGGING]:  

        -v,  --verbose                  :  Print errors and reasons for failed requests and other informations.
        -s,  --silent                   :  Disable banner and version logging.
```

### Basic Scan  
Scan a single URL for CRLF injection:  
```sh
crlfix -u https://example.com
```
### Scanning Multiple URLs  
Provide a file containing multiple URLs:  
```sh
crlfix -l urls.txt
```
### Using stdin/stdout  
```sh
cat urls.txt | crlfix
```
### Adjusting Concurrency  
```sh
crlfix -u https://example.com -c 20
```
### Using Proxy  
```sh
crlfix -u https://example.com -p http://127.0.0.1:8080
```
### Custom Headers  
```sh
crlfix -u https://example.com -H "X-Forwarded-For: 127.0.0.1"
```
### Saving Output  
```sh
crlfix -u https://example.com -O results.txt
```

## Security  
Crlfix is built for ethical hacking and security testing. It is designed to **not cause harm** and should only be used with **proper authorization**. The tool does not exploit vulnerabilities but identifies potential risks for remediation.  

## License  
Crlfix is open-source and released under the [MIT License](https://github.com/RevoltSecurities/Crlfix/blob/main/LICENSE). 

### About:  

The **Crlfix** is a cutting-edge CRLF injection vulnerability scanner developed by **RevoltSecurities** to empower Security Researchers and Penetration Testers. Designed for speed, accuracy, and efficiency, Crlfix automates the detection of CRLF injection vulnerabilities, making security assessments more effective. Released under the MIT License, it reflects our commitment to fostering innovation and collaboration within the open-source community.  

At **RevoltSecurities**, we strive to equip researchers with advanced automation tools that simplify complex security testing, enabling professionals to focus on securing modern web applications and infrastructures.
