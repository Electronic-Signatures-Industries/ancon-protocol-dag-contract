package contract

import (
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

func ToDeployment(schema graphql.ExecutableSchema) ([]byte, error) {
	bz, err := json.Marshal(schema.Schema())

	if err != nil {
		return nil, err
	}

	return bz, nil
}

func HandleSchema(schema []byte) {
	var schemaObj ast.Schema
	err := json.Unmarshal(schema, &schemaObj)
	fmt.Println(schemaObj)
	fmt.Println(err)
}
