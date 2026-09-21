package main

import (
	"testing"

	helpers "github.com/heavybullets8/containers/tests"
)

func TestRunnerContract(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/heavybullets8/business-runner:rolling")
	helpers.RequireCommandSucceeds(
		t,
		image,
		nil,
		"/bin/sh",
		"-lc",
		`set -eu
test "$(id -u)" = 1001
test "$(id -g)" = 1001
getent group docker | grep -q ':123:'
test -x /home/runner/run.sh
test -d /home/runner/externals
command -v sudo
pdftotext -v`,
	)
}
