// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build netbsd,arm

package sys

const (
	ArchFamily    = ARM
	BigEndian     = 0
	CacheLineSize = 32
	PCQuantum     = 4
	Int64Align    = 4
	HugePageSize  = 0
	MinFrameSize  = 4
)

type Uintreg uint32

var (
	PhysPageSize	uintptr
)
