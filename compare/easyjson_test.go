package compare

import (
	"compare-json-libs/compare/data"
	"github.com/mailru/easyjson"
	"testing"
)

func BenchmarkEasyJsonMarshal(b *testing.B) {
	b.Run("simple", getEasyJsonMarshalBenchmark(simpleData))
	b.Run("medium", getEasyJsonMarshalBenchmark(mediumData))
	b.Run("heavy", getEasyJsonMarshalBenchmark(heavyData))
}

func BenchmarkEasyJsonUnmarshal(b *testing.B) {
	var (
		simple data.Simples
		medium data.Mediums
		heavy  data.Heavies
	)

	b.Run("simple", getEasyJsonUnmarshalBenchmark(simpleContent, &simple))
	b.Run("medium", getEasyJsonUnmarshalBenchmark(mediumContent, &medium))
	b.Run("heavy", getEasyJsonUnmarshalBenchmark(heavyContent, &heavy))
}

func getEasyJsonMarshalBenchmark(obj easyjson.Marshaler) func(b *testing.B) {
	return func(b *testing.B) {
		var (
			err    error
			target []byte
		)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			target, err = easyjson.Marshal(obj)
			if err != nil {
				panic(err)
			}

			_ = target
		}
	}
}

func getEasyJsonUnmarshalBenchmark(data []byte, pointer easyjson.Unmarshaler) func(b *testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if err := easyjson.Unmarshal(data, pointer); err != nil {
				panic(err)
			}
		}
	}
}
