package graph
//go:generate go run github.com/99designs/gqlgen
import 	"github.com/Electronic-Signatures-Industries/ancon-protocol-dag-contract/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
}


func NewRootResolver() *Resolver {
	return &Resolver{
		todos: []*model.Todo{
			{
				ID:       "Todo:1",
				Text:     "Buy a cat food",
			},
			{
				ID:       "Todo:2",
				Text:     "Check cat water",
			},
			{
				ID:       "Todo:3",
				Text:     "Check cat meal",
 			},
		},
	}
}
