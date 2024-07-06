package api_utils

import "time"

type Entry struct {
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	Username  string     `json:"twitch_username" db:"twitch_username"`
	UserId    string     `json:"twitch_user_id" db:"twitch_user_id"`
	Notes     string     `json:"notes" db:"notes"`
}
