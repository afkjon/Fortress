package models

import (
	"strconv"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name       string         `json:"name"`
	Project    string         `json:"project"`
	Ticket     string         `json:"ticket"`
	Hours      uint           `json:"hours"`
	TargetDate datatypes.Date `json:"target_date"`
}

func (t Task) GetDateString() string {
	val, _ := t.TargetDate.Value()
	s := val.(time.Time)
	return s.Format("2006-01-02")
}

func (t Task) GetHoursString() string {
	return strconv.FormatUint(uint64(t.Hours), 10)
}

func (t Task) GetIdString() string {
	return strconv.FormatUint(uint64(t.ID), 10)
}
