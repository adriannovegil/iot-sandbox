package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"inditex.com/protector/pkg/config/system"
	"inditex.com/protector/pkg/protector/rules"
	"inditex.com/protector/pkg/util/constants"

	// "github.com/armon/go-metrics"
	"github.com/influxdata/influxql"
	"github.com/rs/zerolog/log"
)

type errorResponses struct {
	Results []*errorResponse `json:"results"`
}

type errorResponse struct {
	Error string `json:"error"`
}

var (
	// Server configuration
	serverPort        = ":" + system.GetServerPortConfiguration()
	serverBuckets     = system.GetServerBucketsConfiguration()
	serverMaxDuration = system.GetServerMaxDurationConfiguration()
	serverSlowQuery   = system.GetServerSlowQueryConfiguration()
	// InfluxDB Configuration
	target = system.GetTargetURL()
)

func main() {
	// flags
	port := flag.String("port", serverPort, constants.DefaultPortUsage)
	maxbuckets := flag.Int("maxbuckets", serverBuckets, constants.DefaultBucketsUsage)
	maxduration := flag.Duration("maxduration", serverMaxDuration, constants.DefaultMaxDurationUsage)
	slowquery := flag.Int64("slowquery", serverSlowQuery, constants.DefaultSlowQueryUsage)
	target := flag.String("target", target, constants.DefaultTargetUsage)
	flag.Parse()

	// metrics.NewGlobal(metrics.DefaultConfig("influx-protector"), sink)
	log.Info().Msgf("Influx Protector Version %s", system.GetVersion())
	log.Info().Msgf("server will run on: %s", *port)
	log.Info().Msgf("redirecting to: %s", *target)
	log.Debug().Msgf("go runtime version: %s", runtime.Version())
	log.Debug().Msgf("CLI args: %#v", os.Args)

	purl, _ := url.Parse(*target)
	proxy := httputil.NewSingleHostReverseProxy(purl)
	options := &rules.Options{
		MaxBuckets:  *maxbuckets,
		SlowQuery:   *slowquery,
		MaxDuration: *maxduration,
	}
	// proxy
	http.HandleFunc("/query", queryFunc(options, proxy))
	http.HandleFunc("/ping", pingFunc(options, proxy))
	http.HandleFunc("/write", writeFunc(options, proxy))
	// Launch server
	http.ListenAndServe(*port, nil)
}

// queryFunc delegate to manage the query requests
func queryFunc(options *rules.Options, proxy *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		inputQuery := strings.TrimSpace(r.URL.Query().Get("q"))

		defer showQueryInfo(time.Now(), inputQuery, options)

		// defer metrics.MeasureSince([]string{"queries", "timing"}, time.Now())

		query, err := influxql.NewParser(strings.NewReader(inputQuery)).ParseStatement()

		if err != nil {
			log.Error().Msg(inputQuery)
			log.Error().Msgf("%v", err)
			// metrics.IncrCounter([]string{"queries", "malformed"}, 1)
			generateErrorResp(err, w)
			return
		}

		if err := rules.RunRules(query, options); err != nil {
			log.Error().Msg(inputQuery)
			log.Error().Msgf("%v", err)
			// metrics.IncrCounter([]string{"queries", "blocked"}, 1)
			generateErrorResp(err, w)
			return
		}

		// metrics.IncrCounter([]string{"queries", "accepted"}, 1)
		proxy.ServeHTTP(w, r)
	}
}

// showQueryInfo logs the query info
func showQueryInfo(start time.Time, query string, options *rules.Options) {
	log.Debug().Msgf("[QUERY] %s", query)

	elapsed := int64(time.Since(start).Seconds() * 1000)
	if elapsed > options.SlowQuery {
		log.Info().Msgf("[SLOWQUERY] '%s' took %dms", query, elapsed)
	}
}

// pingFunc delegate to manage the ping requests
func pingFunc(options *rules.Options, proxy *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		proxy.ServeHTTP(w, r)
	}
}

// writeFunc delegate to manage the write requests
func writeFunc(options *rules.Options, proxy *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		// defer metrics.MeasureSince([]string{"writes", "timing"}, time.Now())

		buf, _ := ioutil.ReadAll(r.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))

		// body := string(buf[:len(buf)])
		// lines := strings.Split(body, "\n")

		r.Body = rdr1

		// datapoints := float32(len(lines))
		// metrics.IncrCounter([]string{"writes", "points"}, datapoints)
		// metrics.SetGauge([]string{"writes", "batchsize"}, datapoints)

		proxy.ServeHTTP(w, r)
	}
}

// setHeaders delegate the request headers
func setHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Influx-Protector-Version", system.GetVersion())
}

// generateErrorResp generate the error response
func generateErrorResp(err error, w http.ResponseWriter) {
	body, jsErr := json.Marshal(&errorResponses{
		Results: []*errorResponse{&errorResponse{
			Error: err.Error(),
		}},
	})

	if jsErr != nil {
		http.Error(w, jsErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
