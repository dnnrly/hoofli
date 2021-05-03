package test_test

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

// nolint
type testContext struct {
	err       error
	cmdResult struct {
		Output string
		Err    error
	}
}

// Errorf is used by the called assertion to report an error and is required to
// make testify assertions work
func (c *testContext) Errorf(format string, args ...interface{}) {
	c.err = fmt.Errorf(format, args...)
}

func (c *testContext) theAppRunsWithParameters(args string) error {
	cmd := exec.Command("../hoofli", strings.Split(args, " ")...)
	output, err := cmd.CombinedOutput()
	c.cmdResult.Output = string(output)
	c.cmdResult.Err = err

	return nil
}

func (c *testContext) theAppExitsWithoutError() error {
	assert.NoError(c, c.cmdResult.Err)
	return c.err
}

func (c *testContext) theAppExitsWithAnError() error {
	assert.Error(c, c.cmdResult.Err)
	return c.err
}

func (c *testContext) theAppOutputContains(expected string) error {
	assert.Contains(c, c.cmdResult.Output, expected)
	return c.err
}

// nolint
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}

// nolint
func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := testContext{}
	ctx.BeforeScenario(func(*godog.Scenario) {})
	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		if err != nil {
			fmt.Printf("Command line output for \"%s\"\n%s", s.GetName(), tc.cmdResult.Output)
		}
	})
	ctx.Step(`^the app runs with parameters "([^"]*)"$`, tc.theAppRunsWithParameters)
	ctx.Step(`^the app exits without error$`, tc.theAppExitsWithoutError)
	ctx.Step(`^the app exits with an error$`, tc.theAppExitsWithAnError)
	ctx.Step(`^the app output contains "([^"]*)"$`, tc.theAppOutputContains)
}
