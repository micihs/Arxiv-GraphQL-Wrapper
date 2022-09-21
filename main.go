package main

const (
	cfgfilepath = "arxiv_graphql_cfg.json"
)

func Main()  {
	err := config.Load(cfgFilePath)
	if err != nil {
		panic(err)
	}

	schema, err := gql.GetSchema()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	log.Printf("Starting %s", config.Config.Service.Name)

	http.Handle("/graphql", h)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Printf("Cannot start %s", config.Config.Service.Name)
		panic(err)
	}
}