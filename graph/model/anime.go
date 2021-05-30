package model

type Anime struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	TitleJp  string `json:"titleJp"`
	Tid      int    `json:"tid"`
	Summary  string `json:"summary"`
	Source   string `json:"source"`
	Studio   string `json:"studio"`
	ImageURL string `json:"imageUrl"`
}
