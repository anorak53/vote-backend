// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateVote struct {
	Name    string `json:"name"`
	Number  int    `json:"number"`
	Details string `json:"details"`
	LogoURL string `json:"logoUrl"`
}

type DeleteVote struct {
	ID int `json:"id"`
}

type EditVote struct {
	Name    string `json:"name"`
	Number  int    `json:"number"`
	Details string `json:"details"`
	LogoURL string `json:"logoUrl"`
}

type Mutation struct {
}

type Query struct {
}

type Result struct {
	Success bool `json:"success"`
}

type VoteList struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Number  int    `json:"number"`
	Details string `json:"details"`
	LogoURL string `json:"logoUrl"`
	Score   int    `json:"Score"`
}

type VoteSelect struct {
	ID            int `json:"id"`
	IDCardNumber  int `json:"ID_CARD_NUMBER"`
	StudentNumber int `json:"STUDENT_NUMBER"`
}
