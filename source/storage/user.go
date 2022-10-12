package storage

import "github.com/Habibullo-1999/notification-bot/source/entity"

func (s *storage) GetUserByID(id string) (*entity.User, error) {
	var user *entity.User
	err := s.db.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *storage) GetUserByTgUsername(username string) (*entity.User, error) {
	var user *entity.User
	err := s.db.Table("users").Where("tg_account = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
