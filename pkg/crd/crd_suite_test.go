package crd

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CRD Suite")
}
