package models

type Type string

const (
	Movie             Type = "movie"
	Serie             Type = "tv_series"
	People            Type = "person"
	Collection        Type = "collection"
	TVNetwork         Type = "tv_network"
	Keyword           Type = "keyword"
	ProductionCompany Type = "production_company"
)
