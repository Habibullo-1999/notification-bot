package entity

import "time"

// Structure ReservationMeetingRoom
type ReservationMeetingRoom struct {
	ID            string     `json:"id" db:"id" `
	UserID        string     `json:"user_id" db:"user_id"`
	MeetingRoomID string     `json:"meeting_room_id" db:"meeting_room_id"`
	StartTime     time.Time  `json:"start_time" db:"start_time"`
	EndTime       time.Time  `json:"end_time" db:"end_time"`
	Status        bool       `json:"status,omitempty" db:"status"`
	Purpose       string     `json:"purpose,omitempty" db:"purpose"`
	RepeatID      string     `json:"-" db:"repeat_id"`
	RepeatDays    *[]int64   `json:"repeat_days" db:"repeat_days" gorm:"type:int[]"`
	CreatedAt     *time.Time `json:"created_at,omitempty" db:"created_at"`
}
