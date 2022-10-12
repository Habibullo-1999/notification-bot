package storage

import (
	"github.com/Habibullo-1999/notification-bot/source/entity"
	"time"
)

func (s *storage) GetAllReservationLastFIveMinutes() ([]*entity.ReservationMeetingRoom, error) {
	var reservation []*entity.ReservationMeetingRoom
	t := time.Now().Add(time.Minute * (-5)).Add(time.Hour * (5))
	err := s.db.Table("reservation_meeting_rooms").Where("created_at > ?", t).Scan(&reservation).Error
	if err != nil {
		return nil, err
	}
	return reservation, nil
}

func (s *storage) GetStatusReservationByID(id string) (bool, error) {
	var active bool
	err := s.db.Table("reservation_meeting_rooms").Where("id = ?", id).Select("status").Find(&active).Error
	if err != nil {
		return false, err
	}
	return active, nil
}
