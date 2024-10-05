// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package coldata

import (
	"os"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

func TestMain(m *testing.M) {
	randutil.SeedForTests()
	os.Exit(m.Run())
}
