package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Target struct {
	Method  string   `json:"method"`
	Path    string   `json:"path"`
	Headers []Header `json:"headers"`
	Body    string   `json:"body"`
}

type SLA struct {
	Latency     int64   `json:"latency"`
	SuccessRate float64 `json:"successRate"`
}

type StressTest struct {
	Rate     uint64 `json:"rps"`
	Duration uint64 `json:"duration"`
	Target   Target `json:"target"`
	SLA      SLA    `json:"sla"`
}

type Config struct {
	BaseURL string       `json:"baseUrl"`
	Tests   []StressTest `json:"tests"`
}

func main() {
	overridenBaseUrl := flag.Arg(0)

	runtime.GOMAXPROCS(runtime.NumCPU())

	tests, err := ioutil.ReadFile("config.json")

	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	config := Config{}

	json.Unmarshal(tests, &config)

	if overridenBaseUrl != "" {
		config.BaseURL = overridenBaseUrl
	}

	for i := 0; i < len(config.Tests); i++ {
		test := config.Tests[i]
		rate := test.Rate
		duration := time.Duration(test.Duration) * time.Second

		headers := &http.Header{}

		for j := 0; j < len(test.Target.Headers); j++ {
			header := test.Target.Headers[j]
			headers.Set(header.Name, header.Value)
		}

		body, err := base64.StdEncoding.DecodeString(test.Target.Body)

		if err != nil {
			os.Stdout.WriteString(err.Error())
			os.Exit(1)
		}

		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: test.Target.Method,
			URL:    config.BaseURL + test.Target.Path,
			Header: *headers,
			Body:   body,
		})
		attacker := vegeta.NewAttacker()

		var metrics vegeta.Metrics
		for res := range attacker.Attack(targeter, rate, duration) {
			metrics.Add(res)
		}
		metrics.Close()

		reporter := vegeta.NewTextReporter(&metrics)
		reporter.Report(os.Stdout)

		if metrics.Success*100 < test.SLA.SuccessRate {
			os.Exit(1)
		}

		if metrics.Latencies.P99.Nanoseconds() > test.SLA.Latency*time.Millisecond.Nanoseconds() {
			os.Exit(1)
		}
	}
}
