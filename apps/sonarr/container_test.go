package main

import (
	"testing"

	helpers "github.com/heavybullets8/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/heavybullets8/sonarr:rolling")
	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{Port: "8989"}, nil)
}
