package xxh3

import (
	"math/bits"
	"unsafe"
)

type (
	ptr = unsafe.Pointer
	ui  = uintptr

	u8  = uint8
	u32 = uint32
	u64 = uint64
)

const (
	_stripe = 64
	_block  = 1024

	prime32_1 = 2654435761
	prime32_2 = 2246822519
	prime32_3 = 3266489917

	prime64_1 = 11400714785074694791
	prime64_2 = 14029467366897019727
	prime64_3 = 1609587929392839161
	prime64_4 = 9650029242287828579
	prime64_5 = 2870177450012600261
)

var key = ptr(&[...]u8{
	0xb8, 0xfe, 0x6c, 0x39, 0x23, 0xa4, 0x4b, 0xbe, 0x7c, 0x01, 0x81, 0x2c, 0xf7, 0x21, 0xad, 0x1c,
	0xde, 0xd4, 0x6d, 0xe9, 0x83, 0x90, 0x97, 0xdb, 0x72, 0x40, 0xa4, 0xa4, 0xb7, 0xb3, 0x67, 0x1f,
	0xcb, 0x79, 0xe6, 0x4e, 0xcc, 0xc0, 0xe5, 0x78, 0x82, 0x5a, 0xd0, 0x7d, 0xcc, 0xff, 0x72, 0x21,
	0xb8, 0x08, 0x46, 0x74, 0xf7, 0x43, 0x24, 0x8e, 0xe0, 0x35, 0x90, 0xe6, 0x81, 0x3a, 0x26, 0x4c,
	0x3c, 0x28, 0x52, 0xbb, 0x91, 0xc3, 0x00, 0xcb, 0x88, 0xd0, 0x65, 0x8b, 0x1b, 0x53, 0x2e, 0xa3,
	0x71, 0x64, 0x48, 0x97, 0xa2, 0x0d, 0xf9, 0x4e, 0x38, 0x19, 0xef, 0x46, 0xa9, 0xde, 0xac, 0xd8,
	0xa8, 0xfa, 0x76, 0x3f, 0xe3, 0x9c, 0x34, 0x3f, 0xf9, 0xdc, 0xbb, 0xc7, 0xc7, 0x0b, 0x4f, 0x1d,
	0x8a, 0x51, 0xe0, 0x4b, 0xcd, 0xb4, 0x59, 0x31, 0xc8, 0x9f, 0x7e, 0xc9, 0xd9, 0x78, 0x73, 0x64,
	0xea, 0xc5, 0xac, 0x83, 0x34, 0xd3, 0xeb, 0xc3, 0xc5, 0x81, 0xa0, 0xff, 0xfa, 0x13, 0x63, 0xeb,
	0x17, 0x0d, 0xdd, 0x51, 0xb7, 0xf0, 0xda, 0x49, 0xd3, 0x16, 0x55, 0x26, 0x29, 0xd4, 0x68, 0x9e,
	0x2b, 0x16, 0xbe, 0x58, 0x7d, 0x47, 0xa1, 0xfc, 0x8f, 0xf8, 0xb8, 0xd1, 0x7a, 0xd0, 0x31, 0xce,
	0x45, 0xcb, 0x3a, 0x8f, 0x95, 0x16, 0x04, 0x28, 0xaf, 0xd7, 0xfb, 0xca, 0xbb, 0x4b, 0x40, 0x7e,
})

// HashString returns the hash of the byte slice.
func HashString(s string) uint64 {
	fn := hash
	if len(s) > 240 {
		fn = hashLarge
	}
	return fn(s)
}

// Hash returns the hash of the byte slice.
func Hash(b []byte) uint64 {
	fn := hash
	if len(b) > 240 {
		fn = hashLarge
	}
	return fn(*(*string)(ptr(&b)))
}

// hash handles all strings of length 0 to 240 inclusive.
func hash(s string) (acc u64) {
	p, l := *(*ptr)(ptr(&s)), u64(len(s))

	if l > 128 {
		var hi, lo u64
		acc = l * prime64_1

		// first 8 groups
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*0))^0xbe4ba423396cfeb8, *(*u64)(ptr(ui(p) + 8*1))^0x1cad21f72c81017c)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*2))^0xdb979083e96dd4de, *(*u64)(ptr(ui(p) + 8*3))^0x1f67b3b7a4a44072)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*4))^0x78e5c0cc4ee679cb, *(*u64)(ptr(ui(p) + 8*5))^0x2172ffcc7dd05a82)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*6))^0x8e2443f7744608b8, *(*u64)(ptr(ui(p) + 8*7))^0x4c263a81e69035e0)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*8))^0xcb00c391bb52283c, *(*u64)(ptr(ui(p) + 8*9))^0xa32e531b8b65d088)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*10))^0x4ef90da297486471, *(*u64)(ptr(ui(p) + 8*11))^0xd8acdea946ef1938)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*12))^0x3f349ce33f76faa8, *(*u64)(ptr(ui(p) + 8*13))^0x1d4f0bc7c7bbdcf9)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 8*14))^0x3159b4cd4be0518a, *(*u64)(ptr(ui(p) + 8*15))^0x647378d9c97e9fc8)
		acc += hi ^ lo

		// avalanche
		acc ^= acc >> 37
		acc *= prime64_3
		acc ^= acc >> 32

		// trailing groups
		top := ui(l) &^ 15
		for i := ui(8 * 16); i < top; i += 16 {
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + i))^*(*u64)(ptr(ui(key) + i - 125)), *(*u64)(ptr(ui(p) + i + 8))^*(*u64)(ptr(ui(key) + i - 117)))
			acc += hi ^ lo
		}

		// last bytes
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + ui(l) - 16))^*(*u64)(ptr(ui(key) + 127)), *(*u64)(ptr(ui(p) + ui(l) - 8))^*(*u64)(ptr(ui(key) + 135)))
		acc += hi ^ lo

		// avalanche
		acc ^= acc >> 37
		acc *= prime64_3
		acc ^= acc >> 32

		return acc
	}

	if l > 16 {
		var hi, lo u64
		acc = l * prime64_1

		if l > 96 {
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 48))^0x3f349ce33f76faa8, *(*u64)(ptr(ui(p) + 48 + 8))^0x1d4f0bc7c7bbdcf9)
			acc += hi ^ lo
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + ui(l) - 64))^0x3159b4cd4be0518a, *(*u64)(ptr(ui(p) + ui(l) - 64 + 8))^0x647378d9c97e9fc8)
			acc += hi ^ lo
		}

		if l > 64 {
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 32))^0xcb00c391bb52283c, *(*u64)(ptr(ui(p) + 32 + 8))^0xa32e531b8b65d088)
			acc += hi ^ lo
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + ui(l) - 48))^0x4ef90da297486471, *(*u64)(ptr(ui(p) + ui(l) - 48 + 8))^0xd8acdea946ef1938)
			acc += hi ^ lo
		}

		if l > 32 {
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 16))^0x78e5c0cc4ee679cb, *(*u64)(ptr(ui(p) + 16 + 8))^0x2172ffcc7dd05a82)
			acc += hi ^ lo
			hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + ui(l) - 32))^0x8e2443f7744608b8, *(*u64)(ptr(ui(p) + ui(l) - 32 + 8))^0x4c263a81e69035e0)
			acc += hi ^ lo
		}

		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + 0))^0xbe4ba423396cfeb8, *(*u64)(ptr(ui(p) + 0 + 8))^0x1cad21f72c81017c)
		acc += hi ^ lo
		hi, lo = bits.Mul64(*(*u64)(ptr(ui(p) + ui(l) - 16))^0xdb979083e96dd4de, *(*u64)(ptr(ui(p) + ui(l) - 16 + 8))^0x1f67b3b7a4a44072)
		acc += hi ^ lo

		// avalanche
		acc ^= acc >> 37
		acc *= prime64_3
		acc ^= acc >> 32

		return acc
	}

	if l > 8 {
		ll1 := *(*u64)(p) ^ 0xbe4ba423396cfeb8
		ll2 := *(*u64)(ptr(ui(p) + ui(l) - 8)) ^ 0x1cad21f72c81017c
		hi, lo := bits.Mul64(ll1, ll2)
		acc = l + ll1 + ll2 + (hi ^ lo)

		// avalanche
		acc ^= acc >> 37
		acc *= prime64_3
		acc ^= acc >> 32

		return acc
	}

	if l > 3 {
		in1 := *(*u32)(p)
		in2 := *(*u32)(ptr(ui(p) + ui(l) - 4))
		in64 := u64(in1) + u64(in2)<<32
		keyed := in64 ^ 0xbe4ba423396cfeb8
		acc = l + (keyed^(keyed>>51))*prime32_1
		acc = (acc ^ (acc >> 47)) * prime64_2

		// avalanche
		acc ^= acc >> 37
		acc *= prime64_3
		acc ^= acc >> 32

		return acc
	}

	if l > 0 {
		c1 := *(*u8)(p)
		c2 := *(*u8)(ptr(ui(p) + (ui(l) >> 1)))
		c3 := *(*u8)(ptr(ui(p) + ui(l) - 1))
		comb := u32(c1) + (u32(c2) << 8) + (u32(c3) << 16) + (u32(l) << 24)
		keyed := u64(comb ^ 0x396cfeb8)
		acc = keyed * prime64_1

		// avalanche
		acc ^= acc >> 37
		acc *= prime64_3
		acc ^= acc >> 32

		return acc
	}

	return 0
}

// hashLarge handles lengths greater than 240.
func hashLarge(s string) u64 {
	p, l := *(*ptr)(ptr(&s)), u64(len(s))

	if avx2 || sse2 {
		return hashVector(p, l)
	}

	acc := l * prime64_1
	accs := [8]u64{prime32_3, prime64_1, prime64_2, prime64_3, prime64_4, prime32_2, prime64_5, prime32_1}

	for l >= _block {
		k := key

		// accs
		for i := 0; i < 16; i++ {
			dv0, kv0 := *(*u64)(ptr(ui(p) + 8*0)), *(*u64)(ptr(ui(k) + 8*0))
			dk0 := dv0 ^ kv0
			accs[0] += dv0 + (dk0&0xffffffff)*(dk0>>32)

			dv1, kv1 := *(*u64)(ptr(ui(p) + 8*1)), *(*u64)(ptr(ui(k) + 8*1))
			dk1 := dv1 ^ kv1
			accs[1] += dv1 + (dk1&0xffffffff)*(dk1>>32)

			dv2, kv2 := *(*u64)(ptr(ui(p) + 8*2)), *(*u64)(ptr(ui(k) + 8*2))
			dk2 := dv2 ^ kv2
			accs[2] += dv2 + (dk2&0xffffffff)*(dk2>>32)

			dv3, kv3 := *(*u64)(ptr(ui(p) + 8*3)), *(*u64)(ptr(ui(k) + 8*3))
			dk3 := dv3 ^ kv3
			accs[3] += dv3 + (dk3&0xffffffff)*(dk3>>32)

			dv4, kv4 := *(*u64)(ptr(ui(p) + 8*4)), *(*u64)(ptr(ui(k) + 8*4))
			dk4 := dv4 ^ kv4
			accs[4] += dv4 + (dk4&0xffffffff)*(dk4>>32)

			dv5, kv5 := *(*u64)(ptr(ui(p) + 8*5)), *(*u64)(ptr(ui(k) + 8*5))
			dk5 := dv5 ^ kv5
			accs[5] += dv5 + (dk5&0xffffffff)*(dk5>>32)

			dv6, kv6 := *(*u64)(ptr(ui(p) + 8*6)), *(*u64)(ptr(ui(k) + 8*6))
			dk6 := dv6 ^ kv6
			accs[6] += dv6 + (dk6&0xffffffff)*(dk6>>32)

			dv7, kv7 := *(*u64)(ptr(ui(p) + 8*7)), *(*u64)(ptr(ui(k) + 8*7))
			dk7 := dv7 ^ kv7
			accs[7] += dv7 + (dk7&0xffffffff)*(dk7>>32)

			p, k = ptr(ui(p)+_stripe), ptr(ui(k)+8)
		}

		// scramble accs
		accs[0] ^= accs[0] >> 47
		accs[0] ^= *(*u64)(ptr(ui(k) + 0*8))
		accs[0] *= prime32_1

		accs[1] ^= accs[1] >> 47
		accs[1] ^= *(*u64)(ptr(ui(k) + 1*8))
		accs[1] *= prime32_1

		accs[2] ^= accs[2] >> 47
		accs[2] ^= *(*u64)(ptr(ui(k) + 2*8))
		accs[2] *= prime32_1

		accs[3] ^= accs[3] >> 47
		accs[3] ^= *(*u64)(ptr(ui(k) + 3*8))
		accs[3] *= prime32_1

		accs[4] ^= accs[4] >> 47
		accs[4] ^= *(*u64)(ptr(ui(k) + 4*8))
		accs[4] *= prime32_1

		accs[5] ^= accs[5] >> 47
		accs[5] ^= *(*u64)(ptr(ui(k) + 5*8))
		accs[5] *= prime32_1

		accs[6] ^= accs[6] >> 47
		accs[6] ^= *(*u64)(ptr(ui(k) + 6*8))
		accs[6] *= prime32_1

		accs[7] ^= accs[7] >> 47
		accs[7] ^= *(*u64)(ptr(ui(k) + 7*8))
		accs[7] *= prime32_1

		l -= 16 * _stripe
	}

	if l > 0 {
		t, k := (l%_block)/_stripe, key

		for i := u64(0); i < t; i++ {
			dv0, kv0 := *(*u64)(ptr(ui(p) + 8*0)), *(*u64)(ptr(ui(k) + 8*0))
			dk0 := dv0 ^ kv0
			accs[0] += dv0 + (dk0&0xffffffff)*(dk0>>32)

			dv1, kv1 := *(*u64)(ptr(ui(p) + 8*1)), *(*u64)(ptr(ui(k) + 8*1))
			dk1 := dv1 ^ kv1
			accs[1] += dv1 + (dk1&0xffffffff)*(dk1>>32)

			dv2, kv2 := *(*u64)(ptr(ui(p) + 8*2)), *(*u64)(ptr(ui(k) + 8*2))
			dk2 := dv2 ^ kv2
			accs[2] += dv2 + (dk2&0xffffffff)*(dk2>>32)

			dv3, kv3 := *(*u64)(ptr(ui(p) + 8*3)), *(*u64)(ptr(ui(k) + 8*3))
			dk3 := dv3 ^ kv3
			accs[3] += dv3 + (dk3&0xffffffff)*(dk3>>32)

			dv4, kv4 := *(*u64)(ptr(ui(p) + 8*4)), *(*u64)(ptr(ui(k) + 8*4))
			dk4 := dv4 ^ kv4
			accs[4] += dv4 + (dk4&0xffffffff)*(dk4>>32)

			dv5, kv5 := *(*u64)(ptr(ui(p) + 8*5)), *(*u64)(ptr(ui(k) + 8*5))
			dk5 := dv5 ^ kv5
			accs[5] += dv5 + (dk5&0xffffffff)*(dk5>>32)

			dv6, kv6 := *(*u64)(ptr(ui(p) + 8*6)), *(*u64)(ptr(ui(k) + 8*6))
			dk6 := dv6 ^ kv6
			accs[6] += dv6 + (dk6&0xffffffff)*(dk6>>32)

			dv7, kv7 := *(*u64)(ptr(ui(p) + 8*7)), *(*u64)(ptr(ui(k) + 8*7))
			dk7 := dv7 ^ kv7
			accs[7] += dv7 + (dk7&0xffffffff)*(dk7>>32)

			p, k, l = ptr(ui(p)+_stripe), ptr(ui(k)+8), l-_stripe
		}

		if l > 0 {
			p = ptr(ui(p) - uintptr(_stripe-l))

			dv0 := *(*u64)(ptr(ui(p) + 8*0))
			dk0 := dv0 ^ 0xea647378d9c97e9f
			accs[0] += dv0 + (dk0&0xffffffff)*(dk0>>32)

			dv1 := *(*u64)(ptr(ui(p) + 8*1))
			dk1 := dv1 ^ 0xc5c3ebd33483acc5
			accs[1] += dv1 + (dk1&0xffffffff)*(dk1>>32)

			dv2 := *(*u64)(ptr(ui(p) + 8*2))
			dk2 := dv2 ^ 0x17eb6313faffa081
			accs[2] += dv2 + (dk2&0xffffffff)*(dk2>>32)

			dv3 := *(*u64)(ptr(ui(p) + 8*3))
			dk3 := dv3 ^ 0xd349daf0b751dd0d
			accs[3] += dv3 + (dk3&0xffffffff)*(dk3>>32)

			dv4 := *(*u64)(ptr(ui(p) + 8*4))
			dk4 := dv4 ^ 0x2b9e68d429265516
			accs[4] += dv4 + (dk4&0xffffffff)*(dk4>>32)

			dv5 := *(*u64)(ptr(ui(p) + 8*5))
			dk5 := dv5 ^ 0x8ffca1477d58be16
			accs[5] += dv5 + (dk5&0xffffffff)*(dk5>>32)

			dv6 := *(*u64)(ptr(ui(p) + 8*6))
			dk6 := dv6 ^ 0x45ce31d07ad1b8f8
			accs[6] += dv6 + (dk6&0xffffffff)*(dk6>>32)

			dv7 := *(*u64)(ptr(ui(p) + 8*7))
			dk7 := dv7 ^ 0xaf280416958f3acb
			accs[7] += dv7 + (dk7&0xffffffff)*(dk7>>32)
		}
	}

	// fmt.Println(accs)

	// merge accs
	hi1, lo1 := bits.Mul64(accs[0]^0x6dd4de1cad21f72c, accs[1]^0xa44072db979083e9)
	acc += hi1 ^ lo1
	hi2, lo2 := bits.Mul64(accs[2]^0xe679cb1f67b3b7a4, accs[3]^0xd05a8278e5c0cc4e)
	acc += hi2 ^ lo2
	hi3, lo3 := bits.Mul64(accs[4]^0x4608b82172ffcc7d, accs[5]^0x9035e08e2443f774)
	acc += hi3 ^ lo3
	hi4, lo4 := bits.Mul64(accs[6]^0x52283c4c263a81e6, accs[7]^0x65d088cb00c391bb)
	acc += hi4 ^ lo4

	// avalanche
	acc ^= acc >> 37
	acc *= prime64_3
	acc ^= acc >> 32

	return acc
}
