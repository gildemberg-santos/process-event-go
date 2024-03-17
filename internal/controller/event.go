package controller

import (
	"strconv"
	"time"

	"github.com/gildemberg-santos/process-event-go/internal/model"
	"github.com/gin-gonic/gin"
)

type Events struct {
	CompanyID       int32  `json:"company"`
	ScriptID        int32  `json:"script"`
	ScriptVersionID int32  `json:"scriptVersion"`
	ClientVersion   string `json:"clientVersion"`
	UserID          string `json:"userId"`
	SessionID       string `json:"sessionId"`
	Event           string `json:"event"`
	Time            int32  `json:"time"`
	Mobile          bool   `json:"isMobile"`
	Referrer        string `json:"referrer"`
	Location        string `json:"location"`
	ScriptName      string `json:"script_name"`
	AnalyticChannel string `json:"analytic_channel"`
	Score           int32  `json:"score"`
	DevMode         bool   `json:"devMode"`
}

func EventController(c *gin.Context) {
	var event = Events{}
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		c.Abort()
		return
	}

	eventModel := model.NewEvents()
	eventModel.CompanyID = event.CompanyID
	eventModel.ScriptID = event.ScriptID
	eventModel.ScriptVersionID = event.ScriptVersionID
	eventModel.ClientVersion = event.ClientVersion
	eventModel.UserID = event.UserID
	eventModel.SessionID = event.SessionID
	eventModel.Event = event.Event
	eventModel.Time = time.Unix(int64(event.Time), 0)
	eventModel.Mobile = strconv.FormatBool(event.Mobile)
	eventModel.Referrer = event.Referrer
	eventModel.Location = event.Location
	eventModel.ScriptName = event.ScriptName
	eventModel.AnalyticChannel = event.AnalyticChannel
	eventModel.Score = event.Score
	err := eventModel.Save()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		c.Abort()
		return
	}

	c.JSON(200, eventModel)
}
