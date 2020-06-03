package main

import (
	"log"
	"strings"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the mutator plugin config.
type Config struct {
	sensu.PluginConfig
	Verbose	bool
}

var (
	mutatorConfig = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-remove-perfdata-mutator",
			Short:    "Sensu Go Mutator to remove Nagios Performance Data from Check Output",
			Keyspace: "sensu.io/plugins/sensu-remove-perfdata-mutator/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		{
			Path:      "verbose",
			Argument:  "verbose",
			Shorthand: "v",
			Default:   false,
			Usage:     "Provide verbose output",
			Value:     &mutatorConfig.Verbose,
		},
	}

)

func main() {
	mutator := sensu.NewGoMutator(&mutatorConfig.PluginConfig, options, checkArgs, executeMutator)
	mutator.Execute()
}

func checkArgs(_ *types.Event) error {
	return nil
}

func executeMutator(event *types.Event) (*types.Event, error) {
	event.Check.Output = strings.TrimSpace(strings.Split(event.Check.Output, "|")[0])
	if mutatorConfig.Verbose {
		log.Printf("Remove Nagios Perf Data, Check output is now %s\n", event.Check.Output)
	}
	return event, nil
}
