// A basic module written in Go to build a base image for a Node.js application.
// The module is written in Go and uses Dagger functions to define the build steps.

package main

import (
	"context"
)

type GoCi struct{}

// create a production build
func (m *GoCi) Build(source *Directory) *Directory {
	return dag.Node().WithContainer(m.buildBaseImage(source)).
		Build().
		Container().
		Directory("./dist")
}

// run unit tests
func (m *GoCi) Test(ctx context.Context, source *Directory) (string, error) {
	return dag.Node().WithContainer(m.buildBaseImage(source)).
		Run([]string{"run", "test:unit", "run"}).
		Stdout(ctx)
}

// build base image
func (m *GoCi) buildBaseImage(source *Directory) *Container {
	return dag.Node().
		WithVersion("21").
		WithNpm().
		WithSource(source).
		Install(nil).
		Container()
}
