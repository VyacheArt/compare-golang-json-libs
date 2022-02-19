//go:generate easyjson -no_std_marshalers -all data.go
//go:generate codecgen -o data_codec.go data.go
package data

//easyjson:json
type (
	Simple struct {
		Title string `json:"title" faker:"word"`
		Text  string `json:"text" faker:"sentence"`

		Likes       int `json:"likes"`
		Subscribers int `json:"subscribers"`

		IsPublic bool `json:"is_public"`
		IsCool   bool `json:"is_cool"`
	}

	Simples []Simple
)

//easyjson:json
type (
	Medium struct {
		Title string   `json:"title" faker:"word"`
		Text  string   `json:"text" faker:"sentence"`
		Lines []string `json:"lines"`

		Likes       int   `json:"likes"`
		Subscribers int   `json:"subscribers"`
		Ids         []int `json:"ids"`

		IsPublic bool   `json:"is_public"`
		IsCool   bool   `json:"is_cool"`
		Options  []bool `json:"options"`
	}

	Mediums []Medium
)

//easyjson:json
type (
	Heavy struct {
		Title string   `json:"title" faker:"word"`
		Text  string   `json:"text" faker:"sentence"`
		Lines []string `json:"lines"`

		Likes       int   `json:"likes"`
		Subscribers int   `json:"subscribers"`
		Ids         []int `json:"ids"`

		IsPublic bool   `json:"is_public"`
		IsCool   bool   `json:"is_cool"`
		Options  []bool `json:"options"`

		MediumElements []Medium `json:"medium_elements"`
	}

	Heavies []Heavy
)
