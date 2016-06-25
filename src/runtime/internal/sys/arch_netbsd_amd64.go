// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build netbsd,amd64

package sys

const (
	ArchFamily    = AMD64
	BigEndian     = 0
	CacheLineSize = 64
	PCQuantum     = 1
	Int64Align    = 8
	HugePageSize  = 1 << 21
	MinFrameSize  = 0
)

type Uintreg uint64
var (
	PhysPageSize  uintptr
)
