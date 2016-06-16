package rahva_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("rahva_test")

func TestRahvaFakt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rahva Suite")
}
