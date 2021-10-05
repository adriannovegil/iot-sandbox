package system

// ServerConfiguration configuration
type ServerConfiguration struct {
	Port        string `yaml:"port"`
	Buckets     int    `yaml:"buckets"`
	MaxDuration int    `yaml:"max-duration"`
	SlowQuery   int64  `yaml:"slow-query"`
}
