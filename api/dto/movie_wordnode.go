package dto

type MovieWordNode struct {
	MovieID int `json:"movie_id"`
}

type MovieWordNodeOrderBy struct {
	MovieID           int  `json:"movie_id"`
	SortByPreposition bool `json:"sort_by_preposition"`
	SortByOccurence   bool `json:"sort_by_occurence"`
	SortByID          bool `json:"sort_by_id"`
}
