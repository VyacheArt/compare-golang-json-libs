package main

import (
	"compare-json-libs/compare/data"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"io/fs"
	"io/ioutil"
	"log"
)

func main() {
	generate(make(data.Simples, 10), "simple.json")
	generate(make(data.Mediums, 10), "medium.json")
	generate(make(data.Heavies, 10), "heavy.json")
}

func generate(object interface{}, name string) {
	_ = faker.SetRandomMapAndSliceSize(50)
	if err := faker.FakeData(&object); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(&object)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(name, jsonData, fs.ModePerm); err != nil {
		log.Fatal(err)
	}
}
