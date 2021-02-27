package e2e

import (
	"flag"
	"os"
	"testing"

	"github.com/kubeedge/sedna/test/e2e/framework"
)

func TestMain(m *testing.M) {
	framework.RegisterFlags(flag.CommandLine)
	flag.Parse()
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}
