package compare

import (
	"compare-json-libs/compare/data"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func BenchmarkJsonIterMarshal(b *testing.B) {
	b.Run("simple", getJsonIterMarshalBenchmark(simpleData))
	b.Run("medium", getJsonIterMarshalBenchmark(mediumData))
	b.Run("heavy", getJsonIterMarshalBenchmark(heavyData))
}

func BenchmarkJsonIterUnmarshal(b *testing.B) {
	var (
		simple data.Simples
		medium data.Mediums
		heavy  data.Heavies
	)

	b.Run("simple", getJsonIterUnmarshalBenchmark(simpleContent, &simple))
	b.Run("medium", getJsonIterUnmarshalBenchmark(mediumContent, &medium))
	b.Run("heavy", getJsonIterUnmarshalBenchmark(heavyContent, &heavy))
}

func getJsonIterMarshalBenchmark(obj interface{}) func(b *testing.B) {
	return func(b *testing.B) {
		var (
			err    error
			target []byte
		)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			target, err = jsoniter.Marshal(obj)
			if err != nil {
				panic(err)
			}

			_ = target
		}
	}
}

func getJsonIterUnmarshalBenchmark(data []byte, pointer interface{}) func(b *testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if err := jsoniter.Unmarshal(data, pointer); err != nil {
				panic(err)
			}
		}
	}
}
