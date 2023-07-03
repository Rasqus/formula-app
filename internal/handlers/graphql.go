package handlers

import (
	"errors"
	"github.com/anras5/formula-app-backend/internal/graph"
	"github.com/anras5/formula-app-backend/internal/models"
	"github.com/graphql-go/graphql"
)

type Graph struct {
	QueryString    string
	Variables      map[string]interface{}
	Config         graphql.SchemaConfig
	queryFields    graphql.Fields
	mutationFields graphql.Fields
}

func NewGraph() *Graph {

	var queryFields = graphql.Fields{
		"listDrivers": &graphql.Field{
			Type:        graphql.NewList(graph.DriverType),
			Description: "Get all drivers",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				drivers, err := Repo.DB.SelectDrivers()
				if err != nil {
					return nil, err
				}
				return drivers, nil
			},
		},
		"getDriver": &graphql.Field{
			Type:        graph.DriverType,
			Description: "Get driver by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					driver, err := Repo.DB.SelectDriver(id)
					if err != nil {
						return nil, err
					}
					return driver, nil
				}
				return nil, errors.New("did not provide id")
			},
		},
		"listRaces": &graphql.Field{
			Type:        graphql.NewList(graph.RaceType),
			Description: "Get all races",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				races, err := Repo.DB.SelectRaces()
				if err != nil {
					return nil, err
				}
				return races, nil
			},
		},
	}

	var mutationFields = graphql.Fields{
		"createDriver": &graphql.Field{
			Type:        graph.DriverType,
			Description: "Create new driver",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"surname": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"country": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"start_number": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"team": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				driver := &models.Driver{
					Name:        p.Args["name"].(string),
					Surname:     p.Args["surname"].(string),
					Country:     p.Args["country"].(string),
					StartNumber: p.Args["start_number"].(int),
					Team:        p.Args["team"].(string),
				}
				driver, err := Repo.DB.InsertDriver(*driver)
				if err != nil {
					return nil, err
				}
				return driver, nil
			},
		},
		"createSeason": &graphql.Field{
			Type:        graph.SeasonType,
			Description: "Create new season",
			Args: graphql.FieldConfigArgument{
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				season := &models.Season{
					Year: p.Args["year"].(int),
				}
				season, err := Repo.DB.InsertSeason(*season)
				if err != nil {
					return nil, err
				}
				return season, nil
			},
		},
	}

	return &Graph{
		queryFields:    queryFields,
		mutationFields: mutationFields,
	}
}

func (g *Graph) Query() (*graphql.Result, error) {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: g.queryFields,
	})
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: g.mutationFields,
	})

	schemaConfig := graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}
	params := graphql.Params{
		Schema:         schema,
		RequestString:  g.QueryString,
		VariableValues: g.Variables,
	}
	response := graphql.Do(params)
	if len(response.Errors) > 0 {
		return nil, response.Errors[0]
	}
	return response, nil
}
