// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

// TestQuery creates a stack and runs a query over the stack's resource ouptputs.
func TestQuery(t *testing.T) {
	t.Parallel()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		// Create Pulumi resources.
		Dir:          "query/step1",
		Dependencies: []string{"@pulumi/pulumi"},
		EditDirs: []integration.EditDir{
			// Try to create resources during `pulumi query`. This should fail.
			{
				Dir:           "query/step2",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: true,
			},
			// Run a query during `pulumi query`. Should succeed.
			{
				Dir:           "query/step3",
				Additive:      true,
				QueryMode:     true,
				ExpectFailure: false,
			},
		},
	})
}