package main

import (
	"testing"

	helpers "github.com/heavybullets8/containers/tests"
)

func TestRunnerContract(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/heavybullets8/business-nix-runner:rolling")
	helpers.RequireCommandSucceeds(
		t,
		image,
		nil,
		"/bin/sh",
		"-lc",
		`set -eu
test "$(id -u)" = 1001
test "$(id -g)" = 1001
test -x /home/runner/run.sh
test -d /home/runner/externals
test "$(nix --version)" = "nix (Nix) 2.35.2"
nix config show experimental-features | grep -q nix-command
nix config show experimental-features | grep -q flakes
test -z "$(nix config show build-users-group)"
nix eval --expr '1 + 1' | grep -q '^2$'
test -w /nix/store`,
	)
}
