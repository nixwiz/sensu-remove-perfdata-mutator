package main

import (
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	event := types.FixtureEvent("entity1", "check1")
	assert.NoError(checkArgs(event))
}

func TestExecuteMutator(t *testing.T) {
	assert := assert.New(t)
	event := types.FixtureEvent("entity1", "check1")
	event.Check.Output = "Output | Perf Data"
	ev, err := executeMutator(event)
	assert.NoError(err)
	assert.Equal(ev.Check.Output, "Output")
}
