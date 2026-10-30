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
lftp --version
pdftotext -v
pg_dump --version | grep -E 'PostgreSQL\) 17\.'
pg_restore --version | grep -E 'PostgreSQL\) 17\.'
pg_isready --version
aws --version
age --version`,
	)
}

func TestCProgram(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/heavybullets8/business-runner:rolling")

	helpers.RequireCommandSucceeds(t, image, nil, "/bin/sh", "-ec", `
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
cat > "$work/main.c" <<'C'
#include <stdio.h>
int main(void) {
    return puts("hello") < 0;
}
C
cc -Wall -Wextra -Werror "$work/main.c" -o "$work/hello"
test "$("$work/hello")" = hello
`)
}
