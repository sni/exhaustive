package exhaustive

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEnum(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "enumvariants/...")
}

func TestSwitch(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "switch/...")
}
