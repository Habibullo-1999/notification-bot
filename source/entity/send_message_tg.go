package entity

import "time"

// SendMessageTg ...
type SendMessageTg struct {
	ID       string     `json:"id"`
	Text     string     `json:"text"`
	Time     *time.Time `json:"time" gorm:"timestamp"`
	ReservID string     `json:"reserv_id"`
	TgUserID int        `json:"tg_user_id"`
	Status   bool       `json:"status"`
	RepeatID string     `json:"repeat_id"`
}

// NewSendMessageTg constructor for SendMessageTg
func NewSendMessageTg(id string, text string, time *time.Time, reservID string, tgUserID int, status bool, repeatID string) *SendMessageTg {
	return &SendMessageTg{ID: id, Text: text, Time: time, ReservID: reservID, TgUserID: tgUserID, Status: status, RepeatID: repeatID}
}
