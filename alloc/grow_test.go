package alloc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkGrowAppend_0_0_10(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 10)
}

func BenchmarkGrowRemake_0_0_10(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 10)
}

func BenchmarkGrowAppend_0_0_20(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 20)
}

func BenchmarkGrowRemake_0_0_20(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 20)
}

func BenchmarkGrowAppend_0_0_29(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 29)
}

func BenchmarkGrowRemake_0_0_29(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 29)
}

func BenchmarkGrowAppend_0_0_30(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 30)
}

func BenchmarkGrowRemake_0_0_30(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 30)
}

func BenchmarkGrowAppend_0_0_31(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 31)
}

func BenchmarkGrowRemake_0_0_31(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 31)
}

func BenchmarkGrowAppend_0_0_32(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 32)
}

func BenchmarkGrowRemake_0_0_32(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 32)
}

func BenchmarkGrowAppend_0_0_33(b *testing.B) {
	benchmarkGrowAppend(b, 0, 0, 33)
}

func BenchmarkGrowRemake_0_0_33(b *testing.B) {
	benchmarkGrowRemake(b, 0, 0, 33)
}

func BenchmarkGrowAppend_8_0_10(b *testing.B) {
	benchmarkGrowAppend(b, 8, 0, 10)
}

func BenchmarkGrowRemake_8_0_10(b *testing.B) {
	benchmarkGrowRemake(b, 8, 0, 10)
}

func BenchmarkGrowAppend_18_0_20(b *testing.B) {
	benchmarkGrowAppend(b, 18, 0, 20)
}

func BenchmarkGrowRemake_18_0_20(b *testing.B) {
	benchmarkGrowRemake(b, 18, 0, 20)
}

func BenchmarkGrowAppend_27_0_29(b *testing.B) {
	benchmarkGrowAppend(b, 27, 0, 29)
}

func BenchmarkGrowRemake_27_0_29(b *testing.B) {
	benchmarkGrowRemake(b, 27, 0, 29)
}

func BenchmarkGrowAppend_28_0_30(b *testing.B) {
	benchmarkGrowAppend(b, 28, 0, 30)
}

func BenchmarkGrowRemake_28_0_30(b *testing.B) {
	benchmarkGrowRemake(b, 28, 0, 30)
}

func BenchmarkGrowAppend_29_0_31(b *testing.B) {
	benchmarkGrowAppend(b, 29, 0, 31)
}

func BenchmarkGrowRemake_29_0_31(b *testing.B) {
	benchmarkGrowRemake(b, 29, 0, 31)
}

func BenchmarkGrowAppend_30_0_32(b *testing.B) {
	benchmarkGrowAppend(b, 30, 0, 32)
}

func BenchmarkGrowRemake_30_0_32(b *testing.B) {
	benchmarkGrowRemake(b, 30, 0, 32)
}

func BenchmarkGrowAppend_31_0_33(b *testing.B) {
	benchmarkGrowAppend(b, 31, 0, 33)
}

func BenchmarkGrowRemake_31_0_33(b *testing.B) {
	benchmarkGrowRemake(b, 31, 0, 33)
}

func BenchmarkGrowAppend_9_0_10(b *testing.B) {
	benchmarkGrowAppend(b, 9, 0, 10)
}

func BenchmarkGrowRemake_9_0_10(b *testing.B) {
	benchmarkGrowRemake(b, 9, 0, 10)
}

func BenchmarkGrowAppend_19_0_20(b *testing.B) {
	benchmarkGrowAppend(b, 19, 0, 20)
}

func BenchmarkGrowRemake_19_0_20(b *testing.B) {
	benchmarkGrowRemake(b, 19, 0, 20)
}

func BenchmarkGrowAppend_28_0_29(b *testing.B) {
	benchmarkGrowAppend(b, 28, 0, 29)
}

func BenchmarkGrowRemake_28_0_29(b *testing.B) {
	benchmarkGrowRemake(b, 28, 0, 29)
}

func BenchmarkGrowAppend_29_0_30(b *testing.B) {
	benchmarkGrowAppend(b, 29, 0, 30)
}

func BenchmarkGrowRemake_29_0_30(b *testing.B) {
	benchmarkGrowRemake(b, 29, 0, 30)
}

func BenchmarkGrowAppend_30_0_31(b *testing.B) {
	benchmarkGrowAppend(b, 30, 0, 31)
}

func BenchmarkGrowRemake_30_0_31(b *testing.B) {
	benchmarkGrowRemake(b, 30, 0, 31)
}

func BenchmarkGrowAppend_31_0_32(b *testing.B) {
	benchmarkGrowAppend(b, 31, 0, 32)
}

func BenchmarkGrowRemake_31_0_32(b *testing.B) {
	benchmarkGrowRemake(b, 31, 0, 32)
}

func BenchmarkGrowAppend_32_0_33(b *testing.B) {
	benchmarkGrowAppend(b, 32, 0, 33)
}

func BenchmarkGrowRemake_32_0_33(b *testing.B) {
	benchmarkGrowRemake(b, 32, 0, 33)
}

// 128 = 2^7.
func BenchmarkGrowAppend_9_128_10(b *testing.B) {
	benchmarkGrowAppend(b, 9, 128, 10)
}

// 128 = 2^7.
func BenchmarkGrowRemake_9_128_10(b *testing.B) {
	benchmarkGrowRemake(b, 9, 128, 10)
}

// 129 = 2^7+1.
func BenchmarkGrowAppend_9_129_10(b *testing.B) {
	benchmarkGrowAppend(b, 9, 129, 10)
}

// 129 = 2^7+1.
func BenchmarkGrowRemake_9_129_10(b *testing.B) {
	benchmarkGrowRemake(b, 9, 129, 10)
}

// 516096 = 2^18+2^17+2^16+2^15+2^14+2^13.
func BenchmarkGrowAppend_19_516096_20(b *testing.B) {
	benchmarkGrowAppend(b, 19, 516096, 20)
}

// 516096 = 2^18+2^17+2^16+2^15+2^14+2^13.
func BenchmarkGrowRemake_19_516096_20(b *testing.B) {
	benchmarkGrowRemake(b, 19, 516096, 20)
}

// 516097 = 2^18+2^17+2^16+2^15+2^14+2^13+1.
func BenchmarkGrowAppend_19_516097_20(b *testing.B) {
	benchmarkGrowAppend(b, 19, 516097, 20)
}

// 516097 = 2^18+2^17+2^16+2^15+2^14+2^13+1.
func BenchmarkGrowRemake_19_516097_20(b *testing.B) {
	benchmarkGrowRemake(b, 19, 516097, 20)
}

// 268427264 = 2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+2^14+
// 2^13.
func BenchmarkGrowAppend_28_268427264_29(b *testing.B) {
	benchmarkGrowAppend(b, 28, 268427264, 29)
}

// 268427264 = 2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+2^14+
// 2^13.
func BenchmarkGrowRemake_28_268427264_29(b *testing.B) {
	benchmarkGrowRemake(b, 28, 268427264, 29)
}

// 268427265 = 2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+2^14+
// 2^13+1.
func BenchmarkGrowAppend_28_268427265_29(b *testing.B) {
	benchmarkGrowAppend(b, 28, 268427265, 29)
}

// 268427265 = 2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+2^14+
// 2^13+1.
func BenchmarkGrowRemake_28_268427265_29(b *testing.B) {
	benchmarkGrowRemake(b, 28, 268427265, 29)
}

// 536862720 = 2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+
// 2^14+2^13.
func BenchmarkGrowAppend_29_536862720_30(b *testing.B) {
	benchmarkGrowAppend(b, 29, 536862720, 30)
}

// 536862720 = 2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+
// 2^14+2^13.
func BenchmarkGrowRemake_29_536862720_30(b *testing.B) {
	benchmarkGrowRemake(b, 29, 536862720, 30)
}

// 536862721 = 2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+
// 2^14+2^13+1.
func BenchmarkGrowAppend_29_536862721_30(b *testing.B) {
	benchmarkGrowAppend(b, 29, 536862721, 30)
}

// 536862721 = 2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+2^15+
// 2^14+2^13+1.
func BenchmarkGrowRemake_29_536862721_30(b *testing.B) {
	benchmarkGrowRemake(b, 29, 536862721, 30)
}

// 1073733632 = 2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+
// 2^15+2^14+2^13.
func BenchmarkGrowAppend_30_1073733632_31(b *testing.B) {
	benchmarkGrowAppend(b, 30, 1073733632, 31)
}

// 1073733632 = 2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+
// 2^15+2^14+2^13.
func BenchmarkGrowRemake_30_1073733632_31(b *testing.B) {
	benchmarkGrowRemake(b, 30, 1073733632, 31)
}

// 1073733633 = 2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+
// 2^15+2^14+2^13+1.
func BenchmarkGrowAppend_30_1073733633_31(b *testing.B) {
	benchmarkGrowAppend(b, 30, 1073733633, 31)
}

// 1073733633 = 2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+2^16+
// 2^15+2^14+2^13+1.
func BenchmarkGrowRemake_30_1073733633_31(b *testing.B) {
	benchmarkGrowRemake(b, 30, 1073733633, 31)
}

// 2147475456 = 2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+
// 2^16+2^15+2^14+2^13.
func BenchmarkGrowAppend_31_2147475456_32(b *testing.B) {
	benchmarkGrowAppend(b, 31, 2147475456, 32)
}

// 2147475456 = 2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+
// 2^16+2^15+2^14+2^13.
func BenchmarkGrowRemake_31_2147475456_32(b *testing.B) {
	benchmarkGrowRemake(b, 31, 2147475456, 32)
}

// 2147475457 = 2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+
// 2^16+2^15+2^14+2^13+1.
func BenchmarkGrowAppend_31_2147475457_32(b *testing.B) {
	benchmarkGrowAppend(b, 31, 2147475457, 32)
}

// 2147475457 = 2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+2^17+
// 2^16+2^15+2^14+2^13+1.
func BenchmarkGrowRemake_31_2147475457_32(b *testing.B) {
	benchmarkGrowRemake(b, 31, 2147475457, 32)
}

// 4294959104 = 2^31+2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+
// 2^17+2^16+2^15+2^14+2^13.
func BenchmarkGrowAppend_32_4294959104_33(b *testing.B) {
	benchmarkGrowAppend(b, 32, 4294959104, 33)
}

// 4294959104 = 2^31+2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+
// 2^17+2^16+2^15+2^14+2^13.
func BenchmarkGrowRemake_32_4294959104_33(b *testing.B) {
	benchmarkGrowRemake(b, 32, 4294959104, 33)
}

// 4294959105 = 2^31+2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+
// 2^17+2^16+2^15+2^14+2^13+1.
func BenchmarkGrowAppend_32_4294959105_33(b *testing.B) {
	benchmarkGrowAppend(b, 32, 4294959105, 33)
}

// 4294959105 = 2^31+2^30+2^29+2^28+2^27+2^26+2^25+2^24+2^23+2^22+2^21+2^20+2^19+2^18+
// 2^17+2^16+2^15+2^14+2^13+1.
func BenchmarkGrowRemake_32_4294959105_33(b *testing.B) {
	benchmarkGrowRemake(b, 32, 4294959105, 33)
}

// from - power of two of initial slice length.
// addition - addition to result of raising two to power from.
// to - power of two of finite slice length.
func benchmarkGrowAppend(b *testing.B, from, addition, to int) {
	begin := 1<<from + addition
	end := 1 << to

	slice := append([]byte(nil), make([]byte, begin)...)

	// prevent compiler optimizations
	slice[len(slice)-1] = 1

	b.ResetTimer()

	slice = append(slice, make([]byte, end-len(slice))...)

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}

// from - power of two of initial slice length.
// addition - addition to result of raising two to power from.
// to - power of two of finite slice length.
func benchmarkGrowRemake(b *testing.B, from, addition, to int) {
	begin := 1<<from + addition
	end := 1 << to

	slice := make([]byte, begin)

	// prevent compiler optimizations
	slice[len(slice)-1] = 1

	b.ResetTimer()

	slice = make([]byte, end)

	b.StopTimer()

	// meaningless check, used to prevent compiler optimizations
	require.NotNil(b, slice)
}
