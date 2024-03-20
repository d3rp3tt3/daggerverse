// A basic module written in Go to build a base image for a Node.js application.
// The module is written in Go and uses Dagger functions to define the build steps.

package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
)

type GoCi struct{}

// create a service from the production image
func (m *GoCi) Serve(source *Directory) *Service {
	return m.Package(source).AsService()
}

// publish an image
func (m *GoCi) Publish(ctx context.Context, source *Directory) (string, error) {
	return m.Package(source).
		Publish(ctx, fmt.Sprintf("ttl.sh/myapp-%.0f:10m", math.Floor(rand.Float64()*10000000)))
}

// create a production image
func (m *GoCi) Package(source *Directory) *Container {
	return dag.Container().From("nginx:1.25-alpine").
		WithDirectory("/usr/share/nginx/html", m.Build(source)).
		WithExposedPort(80)
}

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
