// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func main() {
	// Lay out the prime constant in memory
	primeData := GLOBL("prime_sse", RODATA|NOPTR)
	DATA(0, U32(2654435761))
	DATA(4, U32(2654435761))
	DATA(8, U32(2654435761))
	DATA(12, U32(2654435761))

	TEXT("accum_sse", NOSPLIT, "func(acc *[8]uint64, data, key *byte, len uint64)")
	// %rdi, %rsi, %rdx, %rcx

	acc := Mem{Base: Load(Param("acc"), GP64())}
	data := Mem{Base: Load(Param("data"), GP64())}
	key := Mem{Base: Load(Param("key"), GP64())}
	len := Load(Param("len"), GP64())
	prime := XMM()
	a := [4]VecVirtual{XMM(), XMM(), XMM(), XMM()}

	advance := func(n int) {
		ADDQ(U32(n*64), data.Base)
		ADDQ(U32(n*8), key.Base)
		SUBQ(U32(n*64), len)
	}

	accum := func(n int) {
		doff, koff := 64*n, 8*n

		for n, offset := range []int{0x00, 0x10, 0x20, 0x30} {
			x0, x1, x2 := XMM(), XMM(), XMM()

			MOVOU(data.Offset(doff+offset), x0)
			MOVOU(key.Offset(koff+offset), x1)
			PXOR(x0, x1)
			PSHUFD(Imm(0xf5), x1, x2)
			PMULULQ(x1, x2)
			PADDQ(x0, x2)

			PADDQ(x2, a[n])
		}
	}

	scramble := func() {
		for n, offset := range []int{0x00, 0x10, 0x20, 0x30} {
			x0, x1 := XMM(), XMM()

			MOVOU(a[n], x0)
			PSRLQ(Imm(0x2f), x0)
			PXOR(x0, a[n])
			MOVOU(key.Offset(offset), x0)
			PXOR(x0, a[n])
			PSHUFD(Imm(0xf5), a[n], x1) // 3 3 1 1
			PMULULQ(prime, a[n])
			PMULULQ(prime, x1)
			PSLLQ(Imm(0x20), x1)

			PADDQ(x1, a[n])
		}
	}

	Label("load")
	{
		MOVOU(acc.Offset(0x00), a[0])
		MOVOU(acc.Offset(0x10), a[1])
		MOVOU(acc.Offset(0x20), a[2])
		MOVOU(acc.Offset(0x30), a[3])
		MOVOU(primeData, prime)
	}

	Label("accum_large")
	{
		CMPQ(len, U32(1024))
		JLT(LabelRef("accum"))

		for i := 0; i < 16; i++ {
			accum(i)
		}
		advance(16)

		scramble()
		Load(Param("key"), key.Base)

		JMP(LabelRef("accum_large"))
	}

	Label("accum")
	{
		CMPQ(len, Imm(64))
		JLT(LabelRef("finalize"))

		accum(0)
		advance(1)

		JMP(LabelRef("accum"))
	}

	Label("finalize")
	{
		CMPQ(len, Imm(0))
		JE(LabelRef("return"))

		SUBQ(Imm(64), data.Base)
		ADDQ(len, data.Base)

		Load(Param("key"), key.Base)
		ADDQ(U8(121), key.Base)

		accum(0)
	}

	Label("return")
	{
		MOVOU(a[0], acc.Offset(0x00))
		MOVOU(a[1], acc.Offset(0x10))
		MOVOU(a[2], acc.Offset(0x20))
		MOVOU(a[3], acc.Offset(0x30))
		RET()
	}

	Generate()
}
