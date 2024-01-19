package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	config "github.com/tochiman/monitoring-http/config"
)

func main() {
	for _, target := range config.HttpTargets() {
		check(target)
	}
}

func check(target config.HttpConfig) {
	errorNum := 0
	checkNum := 0
	fatalNum := 2
	errorMessage := ""
	for checkNum < fatalNum {
		checkNum += 1
		targetPath := target.Proto + "://" + target.Host + target.Path
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Timeout: 5 * time.Second, Transport: tr}
		req, err := http.NewRequest("GET", targetPath, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Host", target.Domain)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.StatusCode != 200 {
			errorNum += 1
			if errorNum >= fatalNum {
				errorMessage += targetPath + " [" + target.Name + "] " + "returns " + fmt.Sprint(resp.StatusCode) + "\n"
			}
		} else {
			break
		}

		defer resp.Body.Close()
	}

	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
}
