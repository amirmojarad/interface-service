package dto

import "interface_project/ent"

type GetWordSentencesRequest struct {
	Title  string `json:"title"`
	FileID int    `json:"file_id"`
}

type WordSentencesResponse struct {
	Page       int                    `json:"page"`
	PerPage    int                    `json:"per_page"`
	PageCount  int                    `json:"page_count"`
	TotalCount int                    `json:"total_count"`
	Links      map[string]interface{} `json:"links"`
	Records    []*ent.Word            `json:"records"`
}

type FavoriteWord struct {
	Title string `json:"title"`
}
