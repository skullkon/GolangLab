package models

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"autor"`
	Content string `json:"content"`
}
