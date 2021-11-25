package wasm
import (
	"context"
	

		
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vektah/gqlparser/v2/parser"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/Electronic-Signatures-Industries/ancon-protocol-dag-contract/graph"
	"github.com/Electronic-Signatures-Industries/ancon-protocol-dag-contract/graph/generated"
)
//export ApplyExec
func ApplyExec(op, query string) *graphql.Response {
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	exec := executor.New(schema)
	ctx := context.Background()
	rc, err := exec.CreateOperationContext(ctx, &graphql.RawParams{
		Query:         query,
		OperationName: op,
	})
	if err != nil {
		return exec.DispatchError(ctx, err)
	}

	resp, ctx2 := exec.DispatchOperation(ctx, rc)

	return resp(ctx2)
}
