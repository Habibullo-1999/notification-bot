package service

import (
	"errors"
	"github.com/Habibullo-1999/notification-bot/source/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func (s *service) WorkerAddSendMessage() {

	reservations, err := s.storage.GetAllReservationLastFIveMinutes()
	if err != nil {
		s.logger.Log(err.Error(), "Service", "WorkerAddSendMessage", "GetAllReservationLastFIveMinutes")
	}
	for _, reservation := range reservations {
		user, err := s.storage.GetUserByID(reservation.UserID)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				s.logger.Log(err.Error(), "Service", "WorkerAddSendMessage", "GetUserByID")

			default:
				s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "GetAllReservationLastFIveMinutes")
			}
			continue
		}

		tgUser, err := s.storage.GetTelegramUserByUsername(user.TgAccount)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				s.logger.Log(err.Error(), "Service", "WorkerAddSendMessage", "GetTelegramUserByUsername")
			default:
				s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "GetTelegramUserByUsername")
			}
			continue
		}

		// ******************* Create struct by send notification Thirteen ( 30 ) Minutes To Start ***********************//
		t := reservation.StartTime.Add(time.Minute * (-30))
		sendM := entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("THIRTEEN_MINUTES_TO_START"), &t, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}

		// ******************* Create struct by send notification Five ( 5 ) Minutes To Start ***********************//
		t = reservation.StartTime.Add(time.Minute * (-5))
		sendM = entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("FIVE_MINUTES_TO_START"), &t, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}

		// ******************* Create struct by send notification Start ***********************//
		sendM = entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("START"), &reservation.StartTime, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}

		//******************* Create struct by send notification Half The Time Has Passed ***********************//
		t2 := reservation.EndTime
		t = reservation.StartTime.Add(t2.Sub(reservation.StartTime) / 2)
		sendM = entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("HALF_THE_TIME_HAS_PASSED"), &t, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}

		////******************* Create struct by send notification Left A Little  ***********************//
		t2 = reservation.EndTime
		t = reservation.StartTime.Add(t2.Sub(reservation.StartTime) * 80 / 100)
		sendM = entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("LEFT_A_LITTLE"), &t, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}

		// ******************* Create struct by send notification Left A Five Minutes  ***********************//
		t = reservation.EndTime.Add(time.Minute * (-5))
		sendM = entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("LEFT_A_FIVE_MINUTES"), &t, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}

		// ******************* Create struct by send notification Time Is Over   ***********************//
		sendM = entity.NewSendMessageTg(uuid.New().String(), s.config.GetString("TIME_IS_OVER"), &reservation.EndTime, reservation.ID, tgUser.ID, true, reservation.RepeatID)
		err = s.storage.AddSendMessageTg(sendM)
		if err != nil {
			s.logger.PrintError(err, "Service", "WorkerAddSendMessage", "AddSendMessageTg")
			continue
		}
	}
}
