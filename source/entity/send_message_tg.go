package entity

import "time"

// SendMessageTg ...
type SendMessageTg struct {
	ID       string     `json:"id"`
	Text     string     `json:"text"`
	Time     *time.Time `json:"time" gorm:"timestamp"`
	ReservId string     `json:"reserv_id"`
	TgUserId int        `json:"tg_user_id"`
	Status   bool       `json:"status"`
	RepeatId string     `json:"repeat_id"`
}

func NewSendMessageTg(id string, text string, time *time.Time, reservId string, tgUserId int, status bool, repeatId string) *SendMessageTg {
	return &SendMessageTg{ID: id, Text: text, Time: time, ReservId: reservId, TgUserId: tgUserId, Status: status, RepeatId: repeatId}
}
