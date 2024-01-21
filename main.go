package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/comail/colog"

	config "monitoring-http/config"
)

func GetENV() int {
	const (
		defaultExecutionInterval = 1 //デフォルトの定期実行時間は1分
	)

	ExecutionInterval := os.Getenv("ExecutionInterval")
	if ExecutionInterval == "" {
		return defaultExecutionInterval
	}
	ExecutionIntervalNum, err := strconv.Atoi(ExecutionInterval)
	if err != nil {
		log.Printf("error: ExecutionInterval '%v' acquisition failure...", ExecutionInterval)
		return defaultExecutionInterval
	}
	return ExecutionIntervalNum
}

func check_server(target config.HttpConfig) {
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
			log.Printf("error: %v", err)
			return
		}

		req.Header.Add("Host", target.Domain)
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("error: %v", err)
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
		log.Printf("error: %v", errorMessage)
	}
}

func main() {
	// logの出力設定
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime,
	})
	colog.Register()

	ExecutionInterval := GetENV()

	// 定期実行時間の指定
	ticker := time.NewTicker(time.Duration(ExecutionInterval) * time.Minute)
	defer ticker.Stop()

	log.Printf("info: Monitoring Service is Starting...")

	// 実行回数の記録
	var cont int = 1
	for {
		select {
		case <-ticker.C:
			log.Printf("info: %d回目の定期実行しました", cont)
			cont++
			//定期実行する関数
			for _, target := range config.HttpTargets() {
				check_server(target)
			}
		}
	}

}
