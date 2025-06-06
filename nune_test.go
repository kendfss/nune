// Copyright © The Nune Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nune_test

import (
	"testing"
	"time"

	"github.com/mattn/nune"
)

func benchmarkMilli(b *testing.B, f func()) {
	b.ResetTimer()

	start := time.Now()
	for i := 0; i < b.N; i++ {
		f()
	}
	execTime := time.Since(start)

	b.ReportMetric(0, "ns/op")
	b.ReportMetric((1e3*execTime.Seconds())/float64(b.N), "ms/op")
}

func newTensor() nune.Tensor[float64] {
	return nune.Range[float64](0, 1e7, 1)
}

func benchmarkOp(b *testing.B, f func()) {
	b.Run("1e7Float64Procs1", func(b *testing.B) {
		nune.EnvConfig.NumCPU = 1

		benchmarkMilli(b, func() {
			f()
		})
	})

	b.Run("1e7Float64ProcsN", func(b *testing.B) {
		nune.EnvConfig.NumCPU = 0

		benchmarkMilli(b, func() {
			f()
		})
	})
}
