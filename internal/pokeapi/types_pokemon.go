package pokeapi

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type Type struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type Pokemon struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Exp    int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []Stat `json:"stats"`
	Types  []Type `json:"types"`
}
