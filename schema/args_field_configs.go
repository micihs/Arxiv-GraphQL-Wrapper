package schema

import ( "github.com/graphql-go/graphql" )


var ArgsFieldCFG = gaphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{
		Type: grpahql.Int,
		Description: "Number of feed items to return in the results",
	},
	"offset": &graphql.ArgumentConfig{
		Type: grpahql.Int,
		Description: "Number of feed items to skip in the results",
	},
	"sortBy": &graphql.ArgumentConfig{
		Type: graphql.NewEnum(sortByEnumConfig),
		Description: "Sorting by the field of the feed",
	},
	"sortOrder": &graphql.ArgumentConfig{
		Type: graphql.NewEnum(sortOrderEnumConfig),
		Description: "Sorting the order of the feed",
	},
	"catergory": &graphql.ArgumentConfig{
		Type: grpahql.String,
		Description: "Catergory of the arXiv feed",
	},
}

var sortByEnumConfig = graphql.EnumConfig{
	Name: "sortBy",
	Values: graphql.EnumValueConfigMap{
		"relevance": &graphql.EnumValueConfig{Value: SortByRelevance},
		"lastUpdatedDate": &graphql.EnumValueConfig{Value: SortByLastUpdatedDate},
		"submittedDate": &graphql.EnumValueConfig{Value: SortBySubmittedDate},
	},
}

var sortOrderEnumConfig = graphql.EnumConfig{
	Name: "sortOrder",
	Values: graphql.EnumValueConfigMap{
		"ascending":  &graphql.EnumValueConfig{Value: SortOrderAscending},
		"descending": &graphql.EnumValueConfig{Value: SortOrderDescending},
	},
}

var FieldArgsFieldConfig = graphql.FieldConfigArgument{
	"tag": &graphql.ArgumentConfig{
		Type:        graphql.String,
		Description: "Tag of the specific field",
	},
}

var CategoryArgsFieldConfig = graphql.FieldConfigArgument{
	"field": &graphql.ArgumentConfig{
		Type:        graphql.String,
		Description: "Field Tag of the specific category",
	},
}

var PaperArgsFieldConfig = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type:        graphql.String,
		Description: "ID of the specific arXiv paper",
	},
}

