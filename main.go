package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/comail/colog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	config "monitoring-http/config"
)

func init() {
	// logの出力設定
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime,
	})
	colog.Register()
}

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

func check_server(targetPath string) int {
	// 成功した場合は１、失敗した場合は0を返す
	errorNum := 0
	checkNum := 0
	fatalNum := 2
	errorMessage := ""
	for checkNum < fatalNum {
		checkNum += 1
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Timeout: 5 * time.Second, Transport: tr}
		req, err := http.NewRequest("GET", targetPath, nil)
		if err != nil {
			log.Printf("error: %v", err)
			return 0
		}

		targetDomain := strings.Replace(targetPath, "https://", "", 1)
		targetDomain = strings.Replace(targetPath, "http://", "", 1)

		req.Header.Add("Host", targetDomain)
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("error: %v", err)

			return 0
		}

		if resp.StatusCode != 200 {
			errorNum += 1
			if errorNum >= fatalNum {
				errorMessage += targetPath + "returns " + fmt.Sprint(resp.StatusCode) + "\n"
			}
		} else {
			return 1
		}

		defer resp.Body.Close()
	}

	if errorMessage != "" {
		log.Printf("error: %v", errorMessage)
	}

	return 0
}

func exporter_server() {
	//Exporterサーバ起動
	const port = ":8080"

	log.Printf("info: ListenAndServe of Exporter is http://localhost%s/metrics.", port)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	log.Printf("info: Monitoring Service is Starting...")

	ExecutionInterval := GetENV()
	log.Printf("info: ExecutionInterval is %d minutes.", ExecutionInterval)

	// 定期実行時間の指定
	ticker := time.NewTicker(time.Duration(ExecutionInterval) * time.Minute)
	defer ticker.Stop()

	go exporter_server()

	targetConfig := config.GetTargets()

	// 実行回数の記録
	var cont int = 1
	//prometheusのGauge用配列
	var gauge []prometheus.Gauge
	for {
		select {
		case <-ticker.C:
			log.Printf("info: %d回目の定期実行します", cont)
			// gauge用の配列数
			var contGauge int = 0
			//定期実行する関数
			for _, target := range targetConfig.Targets {
				log.Printf("info: target: %v", target.Name)
				result := check_server(target.URL)
				if cont == 1 {
					// append(gauge, )
					gauge = append(gauge, prometheus.NewGauge(prometheus.GaugeOpts{
						Namespace: target.Namespace,
						Name:      target.Name,
						Help:      target.Help,
					}))

					// 例えば、Prometheusに登録する場合
					prometheus.MustRegister(gauge[contGauge])
				}
				gauge[contGauge].Set(float64(result))
				contGauge++
			}
			cont++
		}
	}

}
