package runmode

import (
	"github.com/golib/assert"
	"testing"
)

func TestRunModeList(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(Production, RunMode("production"))
	assertion.Equal(Development, RunMode("development"))
	assertion.Equal(Test, RunMode("test"))
}

func TestIsValid(t *testing.T) {
	assertion := assert.New(t)

	assertion.True(Production.IsValid())
	assertion.True(Development.IsValid())
	assertion.True(Test.IsValid())
	assertion.False(RunMode("invalid mode").IsValid())
}

func TestIsProduction(t *testing.T) {
	assertion := assert.New(t)

	assertion.True(Production.IsProduction())
	assertion.False(Development.IsProduction())
	assertion.False(Test.IsProduction())
	assertion.False(RunMode("invalid").IsProduction())
}

func TestIsDevelopment(t *testing.T) {
	assertion := assert.New(t)

	assertion.False(Production.IsDevelopment())
	assertion.True(Development.IsDevelopment())
	assertion.False(Test.IsDevelopment())
	assertion.False(RunMode("invalid").IsDevelopment())
}

func TestIsTest(t *testing.T) {
	assertion := assert.New(t)

	assertion.False(Production.IsTest())
	assertion.False(Development.IsTest())
	assertion.True(Test.IsTest())
	assertion.False(RunMode("invalid").IsTest())
}