package models

import (
	"time"
)

type GameLog struct {
	ID         string    `json:"id" gorm:"type:uuid;column:id;primaryKey;"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Start      time.Time `json:"start" gorm:"column:start;"`
	End        time.Time `json:"end" gorm:"column:end;"`
	NbTurn     int       `json:"nb_turn" gorm:"column:nb_turn"`
	NbSnake    int       `json:"nb_snake" gorm:"column:nb_snake"`
	Ruleset    string    `json:"ruleset" gorm:"column:ruleset"`
	Won        bool      `json:"won" gorm:"column:won"`
	Position   int       `json:"position" gorm:"column:position"`
	Width      int       `json:"width" gorm:"column:width"`
	Height     int       `json:"height" gorm:"column:height"`
	AlgVersion int       `json:"algorithm_version" gorm:"column:algorithm_version"`
}
