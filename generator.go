package gogen

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// GeneratorContext is base class that should be used
// as a composite to any other created generator context.
// It supports basic data flow and provides helpers.
//
// This type should encapsulate all widely used methods
// that are needed by the generators, thus may be extended
// by the time.
type GeneratorContext struct {
	// directory to which should all outputs go
	OutputDir string
}

// SetOutputDir will set the output dir of the generator
// to the specified value, which should result in code
// generated to the destination
func (g *GeneratorContext) SetOutputDir(dir string) {
	g.OutputDir = dir
}

// Name is virtual method that should return the
// name of the generator. This is used for the debugging
// purpose
func (g *GeneratorContext) Name() string {
	return "Generator"
}

// PackageName returns the name of the package based on the
// last directory from the OutputDir
func (g *GeneratorContext) PackageName() string {
	// get package chain from the output dir
	packChain := strings.Split(g.OutputDir, "/")
	// get the package (last in the chain)
	return packChain[len(packChain)-1]
}

// Prepare will ensure, that output directory exists
// and all needed values are correctly set
func (g *GeneratorContext) Prepare() error {
	var err error

	// create directories that are needed
	err = os.MkdirAll(g.OutputDir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile will save provided content into the
// specified file with extension .gen.go and output
// directory previously set
func (g *GeneratorContext) SaveFile(name string, content bytes.Buffer) error {
	// calculate path to the file
	filePath := path.Join(g.OutputDir, name+".gen.go")
	// save file
	return ioutil.WriteFile(filePath, content.Bytes(), os.ModePerm)
}