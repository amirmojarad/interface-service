package dto

type GetWordSentencesRequest struct {
	Title  string `json:"title"`
	FileID int    `json:"file_id"`
}
