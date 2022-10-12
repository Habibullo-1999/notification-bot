package storage

import (
	"github.com/Habibullo-1999/notification-bot/source/entity"
	"time"
)

func (s *storage) AddSendMessageTg(sendMessage *entity.SendMessageTg) error {
	err := s.db.Table("send_message_tgs").Create(&sendMessage).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) GetAllSendMessagesByActive() ([]*entity.SendMessageTg, error) {
	var SendMessageTgs []*entity.SendMessageTg
	t := time.Now().Add(time.Minute * (-1)).Add(time.Hour * 5)
	err := s.db.Table("send_message_tgs").Where("status=? AND time BETWEEN  ? and ?", true, t, time.Now().Add(time.Hour*5)).Find(&SendMessageTgs).Error
	if err != nil {
		return nil, err
	}

	return SendMessageTgs, nil
}

func (s *storage) UpdateStatusDeActiveSendMessageByID(id string) error {
	err := s.db.Table("send_message_tgs").Where("id=?", id).Update("status", false).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *storage) DeleteSendMessageByID(id string) error {
	err := s.db.Table("send_message_tgs").Where("id=? ", id).Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}
