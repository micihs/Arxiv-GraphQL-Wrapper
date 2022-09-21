package schema

import ( "github.com/graphql-go/graphql" )

var Papertype = grpahql.NewObject(
	grpahql.ObjectConfig{
		Name: "paper",
		Fields: graphql.Fields{
			"id": &graphql.Field{},
			"updated": &graphql.Field{},
			"published": &graphql.Field{},
			"title": &graphql.Field{},
			"summary": &graphql.Field{},
			"author": &graphql.Field{},
			"link": &graphql.Field{},
			"primaryCatergory": &graphql.Field{},
			"catergory": &graphql.Field{},
		},
	},
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "author",
		Fields: graphql.Fields{
			"name": &grpahql.Field{
				Type: graphql.String,
				Description: "Name of the author",
			},
		},
	},
)

var linkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "link",
		Fields: graphql.Fields{
			"href": &graphql.Field{
				Type:        graphql.String,
				Description: "Hyperlink of the arXiv entry",
			},
			"rel": &graphql.Field{
				Type:        graphql.String,
				Description: "Relationship between link and arXiv entry",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of the link",
			},
		},
	},
)

var catergoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "category",
		Fields: graphql.Fields{
			"term": &graphql.Field{
				Type:        graphql.String,
				Description: "Category Term of the arXiv entry",
			},
			"scheme": &graphql.Field{
				Type:        graphql.String,
				Description: "Schema Definition of category",
			},
		},
	},
)