package generator

import "github.com/gophergala2016/gogen"

var (
	// Model is global registration of the generator
	Model = &ModelGenerator{}
)

// ModelGenerator encapsulates the logic behind
// generating of models
type ModelGenerator struct {
	gogen.GeneratorContext
}

// Generate will call the generator to generate
// results
func (g *ModelGenerator) Generate() error {
	return nil
}