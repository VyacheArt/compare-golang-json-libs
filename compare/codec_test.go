package compare

import (
	"compare-json-libs/compare/data"
	"github.com/ugorji/go/codec"
	"testing"
)

func BenchmarkCodecMarshal(b *testing.B) {
	b.Run("simple", getCodecMarshalBenchmark(simpleData))
	b.Run("medium", getCodecMarshalBenchmark(mediumData))
	b.Run("heavy", getCodecMarshalBenchmark(heavyData))
}

func BenchmarkCodecUnmarshal(b *testing.B) {
	var (
		simple data.Simples
		medium data.Mediums
		heavy  data.Heavies
	)

	b.Run("simple", getCodecUnmarshalBenchmark(simpleContent, &simple))
	b.Run("medium", getCodecUnmarshalBenchmark(mediumContent, &medium))
	b.Run("heavy", getCodecUnmarshalBenchmark(heavyContent, &heavy))
}

func getCodecMarshalBenchmark(obj interface{}) func(b *testing.B) {
	return func(b *testing.B) {
		var target []byte
		handle := new(codec.JsonHandle)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			encoder := codec.NewEncoderBytes(&target, handle)

			if err := encoder.Encode(obj); err != nil {
				panic(err)
			}
		}
	}
}

func getCodecUnmarshalBenchmark(data []byte, pointer interface{}) func(b *testing.B) {
	return func(b *testing.B) {
		handle := new(codec.JsonHandle)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			decoder := codec.NewDecoderBytes(data, handle)
			if err := decoder.Decode(pointer); err != nil {
				panic(err)
			}
		}
	}
}
