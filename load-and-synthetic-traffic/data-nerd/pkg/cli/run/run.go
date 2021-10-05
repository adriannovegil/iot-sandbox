package run

import (
	"encoding/csv"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"blacklemur.com/datanerd/pkg/config/nodes"
	"blacklemur.com/datanerd/pkg/util/mqtt"
	stringUtil "blacklemur.com/datanerd/pkg/util/string"
	pahoMQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ()

// NewCmdRun start a load test
func NewCmdRun() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Start a load test",
		Run: func(cmd *cobra.Command, args []string) {

			MQTTNodes := ReadNodesCSV(
				nodes.NodesConfiguration.Nodes.Path,
				nodes.NodesConfiguration.Nodes.Number,
				nodes.NodesConfiguration.Nodes.Identifier,
			)

			client := mqtt.StartClient(nodes.NodesConfiguration.MQTT.Broker)
			
			for node := range MQTTNodes {
				for _, measurment := range nodes.NodesConfiguration.Measurments {
					StartMessaggingRoutine(node, MQTTNodes[node], client, measurment)
				}
			}
			select {}
		},
	}
}

// StartMessaggingRoutine ...
func StartMessaggingRoutine(nodeID string, nodeStatus interface{}, client pahoMQTT.Client, measurement nodes.Measurment) {
	logrus.Info(nodeID)
	logrus.Info(nodeStatus)
	go messageSender()
}

func messageSender() {
	counter := 0
	for {
		counter++
		randomTime := rand.Intn(nodes.NodesConfiguration.Messaging.PeriodDeviation*2) - nodes.NodesConfiguration.Messaging.PeriodDeviation
		sleep := nodes.NodesConfiguration.Messaging.Period + randomTime
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}

// ReadNodesCSV ...
func ReadNodesCSV(path string, nodes int, nodeIdentifier string) map[string]interface{} {
	nodesCsv, err := os.Open(path)
	nodesMap := make(map[string]interface{})
	if err != nil {
		log.Fatal().Str("ERR", "File not found, exiting...").Send()
	}
	defer nodesCsv.Close()
	reader := csv.NewReader(nodesCsv)
	var headers []string
	var line []string
	var idPosition int = -1
	for i := 0; i <= nodes; i++ {
		if i == 0 {
			headers, _ = reader.Read()
			idPosition = stringUtil.InSlice(nodeIdentifier, headers)
			continue
		}
		line, _ = reader.Read()
		nodeData := make(map[string]string)
		if idPosition != -1 {
			for index, header := range headers {
				if index != idPosition {
					nodeData[header] = line[index]
				}
			}
			nodesMap[line[idPosition]] = nodeData
			continue
		}
		for index, header := range headers {
			nodeData[header] = line[index]
		}
		// GENERATE RANDOM ID

		nodeID := uuid.New()
		nodesMap[nodeID.String()] = nodeData
	}
	return nodesMap
}

// SaveFile Will Save the file, magic right?
func SaveFile(myFile []byte, path string) {
	if err := ioutil.WriteFile(path, myFile, os.FileMode(0644)); err != nil {
		panic(err)
	}
}
