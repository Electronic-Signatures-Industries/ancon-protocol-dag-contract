package main

import (
	"encoding/json"

	"github.com/Electronic-Signatures-Industries/ancon-protocol-dag-contract/contract"
	"github.com/Electronic-Signatures-Industries/ancon-protocol-dag-contract/graph"
	"github.com/Electronic-Signatures-Industries/ancon-protocol-dag-contract/graph/generated"
)

func main() {}

//export New
func New(data []byte) (*contract.DAGContract, error) {

	// Require
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		return nil, err
	}

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers:  graph.New(jsonData),
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	})

	return &contract.DAGContract{Schema: schema, JsonData: jsonData}, nil
}
