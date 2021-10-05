package system

import (
	"os"
	"time"

	"inditex.com/protector/pkg/util/constants"
	"inditex.com/protector/pkg/util/logger"

	"github.com/jinzhu/configor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SystemConfiguration configuration
var SystemConfiguration = struct {
	AppName               string                `yaml:"appname"`
	Version               string                `yaml:"version"`
	LogLevel              string                `yaml:"loglevel"`
	ServerConfiguration   ServerConfiguration   `yaml:"server"`
	InfluxDBConfiguration InfluxDBConfiguration `yaml:"influxdb"`
}{}

func init() {
	// Load configuration file
	configor.Load(&SystemConfiguration, "config.yml")
	// Configure logger
	logger.ConfigLogger(SystemConfiguration.LogLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Debug().
		Interface("SystemConfiguration", SystemConfiguration).
		Msg("System Configuration:")
}

// GetAppName Returns the app name value
func GetAppName() string {
	return SystemConfiguration.AppName
}

// SetAppName set the app name value
func SetAppName(appName string) {
	SystemConfiguration.AppName = appName
}

// GetVersion returns the app version
func GetVersion() string {
	return SystemConfiguration.Version
}

// SetVersion set the app version
func SetVersion(version string) {
	SystemConfiguration.Version = version
}

// GetLogLevel returns the log level value
func GetLogLevel() string {
	return SystemConfiguration.LogLevel
}

// GetServerConfiguration returns the server configuration
func GetServerConfiguration() ServerConfiguration {
	return SystemConfiguration.ServerConfiguration
}

// GetServerPortConfiguration returns the port configuration value.
// Return the default value if parameter not exists in the configuration file
func GetServerPortConfiguration() string {
	if SystemConfiguration.ServerConfiguration.Port != "" {
		return SystemConfiguration.ServerConfiguration.Port
	}
	return constants.DefaultPort
}

// GetServerBucketsConfiguration returns the buckets configuration value.
// Return the default value if parameter not exists in the configuration file
func GetServerBucketsConfiguration() int {
	if SystemConfiguration.ServerConfiguration.Buckets != 0 {
		return SystemConfiguration.ServerConfiguration.Buckets
	}
	return constants.DefaultBuckets
}

// GetServerMaxDurationConfiguration returns the max duration configuration value.
// Return the default value if parameter not exists in the configuration file
func GetServerMaxDurationConfiguration() time.Duration {
	if SystemConfiguration.ServerConfiguration.MaxDuration != 0 {
		return time.Duration(SystemConfiguration.ServerConfiguration.MaxDuration) * time.Hour
	}
	return constants.DefaultMaxDuration
}

// GetServerSlowQueryConfiguration returns the slow query configuration value.
// Return the default value if parameter not exists in the configuration file
func GetServerSlowQueryConfiguration() int64 {
	if SystemConfiguration.ServerConfiguration.SlowQuery != 0 {
		return SystemConfiguration.ServerConfiguration.SlowQuery
	}
	return constants.DefaultSlowQuery
}

// SetServerConfiguration set the server configuration
func SetServerConfiguration(serverConfiguration ServerConfiguration) {
	SystemConfiguration.ServerConfiguration = serverConfiguration
}

// GetInfluxDBConfiguration returns the InfluxDB configuration
func GetInfluxDBConfiguration() InfluxDBConfiguration {
	return SystemConfiguration.InfluxDBConfiguration
}

// GetInfluxDBHostConfiguration returns the slow query configuration value.
// Return the default value if parameter not exists in the configuration file
func GetInfluxDBHostConfiguration() string {
	if SystemConfiguration.InfluxDBConfiguration.Host != "" {
		return SystemConfiguration.InfluxDBConfiguration.Host
	}
	return constants.DefaultInfluxdbHost
}

// GetInfluxDBPortConfiguration returns the slow query configuration value.
// Return the default value if parameter not exists in the configuration file
func GetInfluxDBPortConfiguration() string {
	if SystemConfiguration.InfluxDBConfiguration.Port != "" {
		return SystemConfiguration.InfluxDBConfiguration.Port
	}
	return constants.DefaultInfluxdbPort
}

// GetInfluxDBSSLConfiguration returns is active ssl.
// Return false if parameter not exists in the configuration file
func GetInfluxDBSSLConfiguration() bool {
	return SystemConfiguration.InfluxDBConfiguration.SSL
}

// GetTargetURL return the formatted target URL
// Return the default value if a URL parameter not exists in the configuration
// file
func GetTargetURL() string {
	var targetURL string
	if SystemConfiguration.InfluxDBConfiguration.SSL {
		targetURL += "https://"
	} else {
		targetURL += "http://"
	}
	if SystemConfiguration.InfluxDBConfiguration.Host != "" {
		targetURL += SystemConfiguration.InfluxDBConfiguration.Host
	} else {
		targetURL += constants.DefaultInfluxdbHost
	}
	targetURL += ":"
	if SystemConfiguration.InfluxDBConfiguration.Port != "" {
		targetURL += SystemConfiguration.InfluxDBConfiguration.Port
	} else {
		targetURL += constants.DefaultInfluxdbHost
	}
	return targetURL
}

// SetInfluxDBConfiguration set the InfluxDB configuration
func SetInfluxDBConfiguration(influxDBConfiguration InfluxDBConfiguration) {
	SystemConfiguration.InfluxDBConfiguration = influxDBConfiguration
}
