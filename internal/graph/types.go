package graph

import "github.com/graphql-go/graphql"

var PlaceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Place",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"race_id": &graphql.Field{
				Type: graphql.Int,
			},
			"driver_id": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"place": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DriverType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Driver",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"surname": &graphql.Field{
				Type: graphql.String,
			},
			"country": &graphql.Field{
				Type: graphql.String,
			},
			"start_number": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"team": &graphql.Field{
				Type: graphql.String,
			},
			"places": &graphql.Field{
				Type: graphql.NewList(PlaceType),
			},
		},
	},
)

var RaceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:       "Race",
		Interfaces: nil,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"details": &graphql.Field{
				Type: graphql.String,
			},
			"season_id": &graphql.Field{
				Type: graphql.Int,
			},
			"track_id": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"places": &graphql.Field{
				Type: graphql.NewList(PlaceType),
			},
		},
	},
)

var SeasonType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:       "Season",
		Interfaces: nil,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"races": &graphql.Field{
				Type: graphql.NewList(RaceType),
			},
		},
	},
)

var TrackType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Track",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"details": &graphql.Field{
				Type: graphql.String,
			},
			"country": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"length": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"races": &graphql.Field{
				Type: graphql.NewList(RaceType),
			},
		},
	})
