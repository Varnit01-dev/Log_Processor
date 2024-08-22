package main

import (
	"fmt"
	"regexp"
)

func ParseLog() {
	logsExample := `[2024-07-27T07:39:54.173Z] "GET /healthz HTTP/1.1" 200 - 0 61 225 - "111.114.195.106,10.0.0.11" "okhttp/3.12.1" "0557b0bd-4c1c-4c7a-ab7f-2120d67bee2f" "example.com" "172.16.0.1:8080"`

	
	logFormat := `\[(?P<time_stamp>[^]]+)\] "(?P<http_method>\w+) (?P<request_path>[^ ]+) HTTP\/\d\.\d" (?P<response_code>\d{3}) - (?P<bytes_sent>\d+) (?P<duration>\d+) (?P<request_length>\d+) - "(?P<ips>[^"]+)" "(?P<user_agent>[^"]+)" "(?P<request_id>[^"]+)" "(?P<domain>[^"]+)" "(?P<client_ip>[^"]+:\d+)"`

	
	re := regexp.MustCompile(logFormat)

	
	matches := re.FindStringSubmatch(logsExample)

	
	captureGroups := make(map[string]string)
	names := re.SubexpNames()

	for i, name := range names {
		if i != 0 && name != "" {
			captureGroups[name] = matches[i]
		}
	}

	
	for key, value := range captureGroups {
		fmt.Printf("%-15s => %s\n", key, value)
	}
}

func main() {
	ParseLog()
}
