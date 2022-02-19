package compare

import (
	"compare-json-libs/compare/data"
	_ "embed"
	"encoding/json"
)

//Initial entire content
var (
	//go:embed data/simple.json
	simpleContent []byte

	//go:embed data/medium.json
	mediumContent []byte

	//go:embed data/heavy.json
	heavyContent []byte
)

//Initial parsed data
var (
	simpleData data.Simples
	mediumData data.Mediums
	heavyData  data.Heavies
)

func init() {
	if err := json.Unmarshal(simpleContent, &simpleData); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(mediumContent, &mediumData); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(heavyContent, &heavyData); err != nil {
		panic(err)
	}
}
