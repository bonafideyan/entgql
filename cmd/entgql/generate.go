package main

import (
	"fmt"
	"log"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/pkg/errors"
	"github.com/sijad/entgql/generator"
)

func generate(entSchemaPath, generatedGraphqlPath, generatedResolversPath, generatedModelsPath string, tmplAdditionals []string) error {
	graph, err := entc.LoadGraph(entSchemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	g := &generator.Generator{Graph: graph}

	if err := g.SchemaDefinition(generatedGraphqlPath, tmplAdditionals); err != nil {
		return fmt.Errorf("generate schema: %w", err)
	}

	err = entc.Generate(entSchemaPath, &gen.Config{
		Features: []gen.Feature{
			gen.FeatureLock,
			gen.FeatureUpsert,
			gen.FeatureModifier,
			gen.FeatureEntQL,
			gen.FeaturePrivacy,
			gen.FeatureExecQuery,
			gen.FeatureNamedEdges,
			gen.FeatureSnapshot,
		},
		Templates: entgql.AllTemplates,
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	return gqlgenGenerate()
}

func gqlgenGenerate() error {
	var cfg *config.Config
	var err error

	cfg, err = config.LoadConfigFromDefaultLocations()
	if os.IsNotExist(errors.Cause(err)) {
		cfg, err = config.LoadDefaultConfig()
	}

	if err != nil {
		return err
	}

	if err = api.Generate(cfg); err != nil {
		return err
	}
	return nil
}

// func generatedModels(g *generator.Generator, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	return g.Models(file)
// }

// func generatedResolvers(g *generator.Generator, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	return g.Resolvers(file)
// }
