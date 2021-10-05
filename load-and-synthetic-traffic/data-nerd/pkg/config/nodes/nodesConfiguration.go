package nodes

import (

	// "github.com/jinzhu/configor"

	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
)

// NodesConfiguration configuration
var NodesConfiguration = struct {
	Nodes       Nodes        `yaml:"nodes"`
	Measurments []Measurment `yaml:"measurments"`
	Messaging   Messaging    `yaml:"messaging"`
	MQTT        MQTT         `yaml:"mqtt"`
}{}

// Nodes ...
type Nodes struct {
	Path       string `yaml:"path"`
	Number     int    `yaml:"number"`
	Identifier string `yaml:"identifier"`
}

// Measurment ...
type Measurment struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

// Messaging ...
type Messaging struct {
	Period          int `yaml:"period"`
	PeriodDeviation int `yaml:"periodDeviation"`
	SkipProbability int `yaml:"skipProbability"`
}

// MQTT ...
type MQTT struct {
	Broker string `yaml:"broker"`
}

func init() {
	// Load configuration file
	err := configor.Load(&NodesConfiguration, "config.nodes.yml")
	if err != nil {
		logrus.Warn(err)
		log.Panic().Str("ERR", "Configuration file not found").Send()
	}
	log.Info().Str("MSG", "Nodes Configuration loaded")
}

// // GetPath Returns the app name value
// func GetPath() string {
// 	return NodesConfiguration.Node.Path
// }

// // SetPath ...
// func SetPath(path string) {
// 	NodesConfiguration..NodePath = path
// }

// // GetNumber set the app name value
// func GetNumber() int {
// 	return NodesConfiguration.Number
// }

// // SetNumber set the app name value
// func SetNumber(number int) {
// 	NodesConfiguration.Number = number
// }
