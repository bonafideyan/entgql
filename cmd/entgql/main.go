package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var templates []string
	cmd := &cobra.Command{Use: "gentc"}
	gCmd := &cobra.Command{
		Use:     "generate",
		Short:   "generate go code and GraphQL schema for `ent` directory",
		Example: "entgql generate ./ent/schema",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, paths []string) {
			entSchemaPath := paths[0]
			if err := generate(
				entSchemaPath,
				defaultGeneratedGraphqlSchemaPath,
				defaultGeneratedResolversPath,
				defaultGeneratedModelsPath,
				templates,
			); err != nil {
				log.Fatalln(err)
			}
			// glog.Info("3rd entgql")
		},
	}
	gCmd.Flags().StringSliceVarP(&templates, "template", "", nil, "external templates to execute")
	cmd.AddCommand(
		gCmd,
	)

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

const (
	defaultGeneratedGraphqlSchemaPath = "./graph/"
	defaultGeneratedResolversPath     = "./graph/ent_schema.resolvers.go"
	defaultGeneratedModelsPath        = "./graph/model/ent_models_gen.go"
)
