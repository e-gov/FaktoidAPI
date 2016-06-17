package ehak_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("ehak_test")

func TestEHAK(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EHAK test")
}
