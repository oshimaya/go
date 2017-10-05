// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sym

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"debug/elf"
)

// Reloc is a relocation.
//
// The typical Reloc rewrites part of a symbol at offset Off to address Sym.
// A Reloc is stored in a slice on the Symbol it rewrites.
//
// Relocations are generated by the compiler as the type
// cmd/internal/obj.Reloc, which is encoded into the object file wire
// format and decoded by the linker into this type. A separate type is
// used to hold linker-specific state about the relocation.
//
// Some relocations are created by cmd/link.
type Reloc struct {
	Off     int32            // offset to rewrite
	Siz     uint8            // number of bytes to rewrite, 1, 2, or 4
	Done    bool             // set to true when relocation is complete
	Variant RelocVariant     // variation on Type
	Type    objabi.RelocType // the relocation type
	Add     int64            // addend
	Xadd    int64            // addend passed to external linker
	Sym     *Symbol          // symbol the relocation addresses
	Xsym    *Symbol          // symbol passed to external linker
}

// RelocVariant is a linker-internal variation on a relocation.
type RelocVariant uint8

const (
	RV_NONE RelocVariant = iota
	RV_POWER_LO
	RV_POWER_HI
	RV_POWER_HA
	RV_POWER_DS

	// RV_390_DBL is a s390x-specific relocation variant that indicates that
	// the value to be placed into the relocatable field should first be
	// divided by 2.
	RV_390_DBL

	RV_CHECK_OVERFLOW RelocVariant = 1 << 7
	RV_TYPE_MASK      RelocVariant = RV_CHECK_OVERFLOW - 1
)

func RelocName(arch *sys.Arch, r objabi.RelocType) string {
	// We didn't have some relocation types at Go1.4.
	// Uncomment code when we include those in bootstrap code.

	switch {
	case r >= 512: // Mach-O
		// nr := (r - 512)>>1
		// switch ctxt.Arch.Family {
		// case sys.AMD64:
		// 	return macho.RelocTypeX86_64(nr).String()
		// case sys.ARM:
		// 	return macho.RelocTypeARM(nr).String()
		// case sys.ARM64:
		// 	return macho.RelocTypeARM64(nr).String()
		// case sys.I386:
		// 	return macho.RelocTypeGeneric(nr).String()
		// default:
		// 	panic("unreachable")
		// }
	case r >= 256: // ELF
		nr := r - 256
		switch arch.Family {
		case sys.AMD64:
			return elf.R_X86_64(nr).String()
		case sys.ARM:
			return elf.R_ARM(nr).String()
		case sys.ARM64:
			return elf.R_AARCH64(nr).String()
		case sys.I386:
			return elf.R_386(nr).String()
		case sys.MIPS, sys.MIPS64:
			// return elf.R_MIPS(nr).String()
		case sys.PPC64:
			// return elf.R_PPC64(nr).String()
		case sys.S390X:
			// return elf.R_390(nr).String()
		default:
			panic("unreachable")
		}
	}

	return r.String()
}
