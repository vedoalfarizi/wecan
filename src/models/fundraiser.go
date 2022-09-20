package models

import "time"

type Fundraiser struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SheetID     string    `json:"sheet_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateFundraiserPayload struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateFundraiserPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateFundraiserSheetPayload struct {
	SheetID string `json:"sheet_id"`
}
