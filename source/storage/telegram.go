package storage

import (
	"github.com/Habibullo-1999/notification-bot/source/entity"
)

func (s *storage) GetTelegramUserByUsername(username string) (*entity.TelegramUsers, error) {
	var telegramUser *entity.TelegramUsers
	err := s.db.Table("telegram_users").Where("username = ?", username).First(&telegramUser).Error
	if err != nil {
		return nil, err
	}
	return telegramUser, nil
}

func (s *storage) AddTgUser(tgUser *entity.TelegramUsers) error {
	err := s.db.Table("telegram_users").Save(&tgUser).Error
	if err != nil {
		return err
	}
	return nil
}
