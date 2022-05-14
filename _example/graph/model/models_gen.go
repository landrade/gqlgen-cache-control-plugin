// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	Post *Post  `json:"post"`
	Text string `json:"text"`
}

type Post struct {
	ID                int        `json:"id"`
	Votes             int        `json:"votes"`
	Comments          []*Comment `json:"comments"`
	ReadByCurrentUser bool       `json:"readByCurrentUser"`
}