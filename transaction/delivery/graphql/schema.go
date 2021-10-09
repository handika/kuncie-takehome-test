package graphql

import "github.com/graphql-go/graphql"

// TransactionGraphQL holds transaction information with graphql object
var TransactionGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"grand_total": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// TransactionEdgeGraphQL holds transaction edge information with graphql object
var TransactionEdgeGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TransactionEdge",
		Fields: graphql.Fields{
			"node": &graphql.Field{
				Type: TransactionGraphQL,
			},
			"cursor": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// TransactionResultGraphQL holds transaction result information with graphql object
var TransactionResultGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TransactionResult",
		Fields: graphql.Fields{
			"edges": &graphql.Field{
				Type: graphql.NewList(TransactionEdgeGraphQL),
			},
			"pageInfo": &graphql.Field{
				Type: pageInfoGraphQL,
			},
		},
	},
)

var pageInfoGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PageInfo",
		Fields: graphql.Fields{
			"endCursor": &graphql.Field{
				Type: graphql.String,
			},
			"hasNextPage": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

// Schema is struct which has method for Query and Mutation. Please init this struct using constructor function.
type Schema struct {
	transactionResolver Resolver
}

// Query initializes config schema query for graphql server.
func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetTransactionByID": &graphql.Field{
				Type:        TransactionGraphQL,
				Description: "Get Transaction By ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: s.transactionResolver.GetTransactionByID,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

var TransactionInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "transactionInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"user_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(TransactionDetailInputType),
		},
	},
})

var TransactionDetailInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "transactionDetailInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"product_id": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"qty":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})

// Mutation initializes config schema mutation for graphql server.
func (s Schema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"StoreTransaction": &graphql.Field{
				Type:        graphql.String,
				Description: "Store a new transaction",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(TransactionInputType),
					},
				},
				Resolve: s.transactionResolver.StoreTransaction,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

// NewSchema initializes Schema struct which takes resolver as the argument.
func NewSchema(transactionResolver Resolver) Schema {
	return Schema{
		transactionResolver: transactionResolver,
	}
}
