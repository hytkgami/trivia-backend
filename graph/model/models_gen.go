// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Connection interface {
	IsConnection()
	GetPageInfo() *PageInfo
	GetEdges() []Edge
}

type Edge interface {
	IsEdge()
	GetCursor() string
	GetNode() Node
}

type Node interface {
	IsNode()
	GetID() string
}

type Answer struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (Answer) IsNode()            {}
func (this Answer) GetID() string { return this.ID }

type AnswerPayload struct {
	Answer *Answer `json:"answer"`
}

type CreateLobbyPayload struct {
	Lobby *Lobby `json:"lobby"`
}

type CreateQuestionPayload struct {
	Questions []*Question `json:"questions"`
}

type Lobby struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

func (Lobby) IsNode()            {}
func (this Lobby) GetID() string { return this.ID }

type LobbyConnection struct {
	Edges    []*LobbyEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

func (LobbyConnection) IsConnection()               {}
func (this LobbyConnection) GetPageInfo() *PageInfo { return this.PageInfo }
func (this LobbyConnection) GetEdges() []Edge {
	if this.Edges == nil {
		return nil
	}
	interfaceSlice := make([]Edge, 0, len(this.Edges))
	for _, concrete := range this.Edges {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type LobbyEdge struct {
	Cursor string `json:"cursor"`
	Node   *Lobby `json:"node"`
}

func (LobbyEdge) IsEdge()                {}
func (this LobbyEdge) GetCursor() string { return this.Cursor }
func (this LobbyEdge) GetNode() Node     { return *this.Node }

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	Cursor      string `json:"cursor"`
}

type Question struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	OrderNumber int    `json:"orderNumber"`
	Score       int    `json:"score"`
}

func (Question) IsNode()            {}
func (this Question) GetID() string { return this.ID }

type QuestionInput struct {
	Title       string `json:"title"`
	OrderNumber int    `json:"orderNumber"`
	Score       int    `json:"score"`
}

type SigninPayload struct {
	User *User `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type LobbiesQueryOrderBy string

const (
	LobbiesQueryOrderByLobbyID LobbiesQueryOrderBy = "LOBBY_ID"
)

var AllLobbiesQueryOrderBy = []LobbiesQueryOrderBy{
	LobbiesQueryOrderByLobbyID,
}

func (e LobbiesQueryOrderBy) IsValid() bool {
	switch e {
	case LobbiesQueryOrderByLobbyID:
		return true
	}
	return false
}

func (e LobbiesQueryOrderBy) String() string {
	return string(e)
}

func (e *LobbiesQueryOrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LobbiesQueryOrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LobbiesQueryOrderBy", str)
	}
	return nil
}

func (e LobbiesQueryOrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
