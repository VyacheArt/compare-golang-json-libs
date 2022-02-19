package compare

import (
	"compare-json-libs/compare/data"
	"encoding/json"
	"testing"
)

func BenchmarkStdJsonMarshal(b *testing.B) {
	b.Run("simple", getStdJsonMarshalBenchmark(simpleData))
	b.Run("medium", getStdJsonMarshalBenchmark(mediumData))
	b.Run("heavy", getStdJsonMarshalBenchmark(heavyData))
}

func BenchmarkStdJsonUnmarshal(b *testing.B) {
	var (
		simple data.Simples
		medium data.Mediums
		heavy  data.Heavies
	)

	b.Run("simple", getStdJsonUnmarshalBenchmark(simpleContent, &simple))
	b.Run("medium", getStdJsonUnmarshalBenchmark(mediumContent, &medium))
	b.Run("heavy", getStdJsonUnmarshalBenchmark(heavyContent, &heavy))
}

func getStdJsonMarshalBenchmark(obj interface{}) func(b *testing.B) {
	return func(b *testing.B) {
		var (
			err    error
			target []byte
		)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			target, err = json.Marshal(obj)
			if err != nil {
				panic(err)
			}

			_ = target
		}
	}
}

func getStdJsonUnmarshalBenchmark(data []byte, pointer interface{}) func(b *testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if err := json.Unmarshal(data, pointer); err != nil {
				panic(err)
			}
		}
	}
}
