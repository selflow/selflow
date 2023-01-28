package config

import (
	"encoding/json"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	sourceFile := "testdata/flow.yaml"

	sourceFileContent, err := os.ReadFile(sourceFile)
	if err != nil {
		panic(err)
	}

	strictFlow, err := Parse(sourceFileContent)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(strictFlow)
	if err != nil {
		panic(err)
	}

	answer := &Flow{}

	strictFlowFileContent, _ := os.ReadFile("testdata/flow_as_json.json")

	_ = json.Unmarshal(strictFlowFileContent, answer)
	strictFlowAnswer, _ := json.Marshal(answer)

	if string(strictFlowAnswer) != string(b) {
		t.Errorf("parsing error")
	}
}
