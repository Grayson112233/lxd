package api

type Restriction struct {
	ProjectName 	string `json:"project_name" yaml:"project_name"`
	RestrictionKey 	string `json:"restriction_key" yaml:"restriction_key"`
	Enabled			bool `json:"enabled" yaml:"enabled"`
}

type Restrictions struct {
	Restrictions []restriction `json:"restrictions" yaml:"restrictions"`
}

type RestrictionsList struct {
	RestrictionsList []string `json:"restrictions_list" yaml:"restrictions_list"`
}

