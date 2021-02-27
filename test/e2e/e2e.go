package e2e

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"k8s.io/component-base/logs"

	"github.com/kubeedge/sedna/test/e2e/framework"
)

var _ = ginkgo.BeforeSuite(func() {
})

var _ = ginkgo.AfterSuite(func() {
})

func RunE2ETests(t *testing.T) {
	logs.InitLogs()
	defer logs.FlushLogs()

	gomega.RegisterFailHandler(framework.Fail)

	ginkgo.RunSpecs(t, "Sedna e2e suite")

}
