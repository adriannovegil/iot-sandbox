package constants

import (
	"time"
)

const (
	// Server

	// DefaultPort default server port
	DefaultPort = "8087"
	// DefaultPortUsage default port usage text
	DefaultPortUsage = "default server port, ':8087', ':8086'..."
	// DefaultBuckets default buckets configuration
	DefaultBuckets = 1000
	// DefaultBucketsUsage default bucket usage text
	DefaultBucketsUsage = "default buckets 1000"
	// DefaultMaxDuration default max duration server configuration
	DefaultMaxDuration = time.Hour * 24 * 90
	// DefaultMaxDurationUsage default max duration usage text
	DefaultMaxDurationUsage = "default max duration 2160h (90 days)"
	// DefaultSlowQuery default server slow query configuration
	DefaultSlowQuery = 1000
	// DefaultSlowQueryUsage default slow query usage text
	DefaultSlowQueryUsage = "default slowquery time in milliseconds 1000"

	// InfluxDB

	// DefaultInfluxdbHost default InfluxDB host configuration
	DefaultInfluxdbHost = "localhost"
	// DefaultInfluxdbPort default InfluxDB port configuration
	DefaultInfluxdbPort = "8086"
	// DefaultTargetUsage default InfluxDB target configuration
	DefaultTargetUsage = "default redirect url, 'http://influxdb-proxy:8086'"
)
