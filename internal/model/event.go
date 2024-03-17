package model

import (
	"time"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	CompanyID       int32     `json:"company_id"`
	ScriptID        int32     `json:"script_id"`
	ScriptVersionID int32     `json:"script_version_id"`
	ClientVersion   string    `json:"client_version"`
	UserID          string    `json:"user_id"`
	SessionID       string    `json:"session_id"`
	Event           string    `json:"event"`
	Time            time.Time `json:"time"`
	Mobile          string    `json:"mobile"`
	Referrer        string    `json:"referrer"`
	Location        string    `json:"location"`
	ScriptName      string    `json:"script_name"`
	AnalyticChannel string    `json:"analytic_channel"`
	Score           int32     `json:"score"`
}

func NewEvents() *Events {
	return &Events{}
}

func (e *Events) Save() error {
	conection := ConectionDB{}
	conection.Open()
	defer conection.Close()
	return conection.DB.Create(e).Error
}
