// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Code generated by "stringer"; DO NOT EDIT.

package scop

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MutationType-1]
	_ = x[BackfillType-2]
	_ = x[ValidationType-3]
}

func (i Type) String() string {
	switch i {
	case MutationType:
		return "MutationType"
	case BackfillType:
		return "BackfillType"
	case ValidationType:
		return "ValidationType"
	default:
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
