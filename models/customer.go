package models

import "time"

type Customer struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	SignupDate    time.Time `json:"signup_date"`
	Location      string    `json:"location"`
	LifetimeValue float64   `json:"lifetime_value"`
}
