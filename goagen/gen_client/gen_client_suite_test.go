package genclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGenClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GenClient Suite")
}
