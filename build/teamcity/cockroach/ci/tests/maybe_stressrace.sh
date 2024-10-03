#!/usr/bin/env bash

# Copyright 2021 The Cockroach Authors.
#
# Use of this software is governed by the CockroachDB Software License
# included in the /LICENSE file.


set -euo pipefail

dir="$(dirname $(dirname $(dirname $(dirname $(dirname "${0}")))))"

source "$dir/teamcity-support.sh"  # For $root, would_stress
source "$dir/teamcity-bazel-support.sh"  # For run_bazel

if would_stress; then
    tc_start_block "Run stress tests"
    run_bazel env BUILD_VCS_NUMBER="$BUILD_VCS_NUMBER" build/teamcity/cockroach/ci/tests/maybe_stress_impl.sh stressrace
    tc_end_block "Run stress tests"
fi
