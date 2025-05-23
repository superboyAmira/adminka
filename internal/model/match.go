package model

import "github.com/google/uuid"

type MatchResult string

const (
	ResultWhiteWin MatchResult = "WhiteWin"
	ResultBlackWin MatchResult = "BlackWin"
	ResultDraw     MatchResult = "Draw"
)

type Match struct {
	ID      uuid.UUID   `json:"id"`
	Player1 uuid.UUID   `json:"player1"`
	Player2 uuid.UUID   `json:"player2"`
	Result  MatchResult `json:"result"`
}
