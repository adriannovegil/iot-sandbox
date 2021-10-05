package system

// InfluxDBConfiguration configuration
type InfluxDBConfiguration struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	SSL  bool   `yaml:"ssl"`
}
