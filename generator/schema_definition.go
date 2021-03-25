package generator

import (
	"io"
	"os"

	"entgo.io/ent/entc/gen"
)

func (g *Generator) SchemaDefinition(filePath string, tmplAdditionals []string) error {
	initTempl(tmplAdditionals)
	os.MkdirAll(filePath, 0755)
	file, err := os.Create(filePath + "schema.ent.base.graphql")
	if err != nil {
		println(err.Error())
		return err
	}
	defer file.Close()

	if err := templates.ExecuteTemplate(file, "template/base_sdl.tmpl", nil); err != nil {
		return err
	}

	for _, t := range g.Graph.Nodes {
		file, err := os.Create(filePath + "schema.ent." + t.Name + ".graphql")
		if err != nil {
			println(err.Error())
			return err
		}
		defer file.Close()

		if err := nodeSchemaDefinition(file, t); err != nil {
			return err
		}
	}
	return nil
}

func nodeSchemaDefinition(wr io.Writer, t *gen.Type) error {
	return templates.ExecuteTemplate(wr, "template/node_sdl.tmpl", t)
}
