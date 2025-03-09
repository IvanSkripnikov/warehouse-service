package models

const ServiceDatabase = "NotificationService"

type BookingItem struct {
	ItemID int `json:"id"`
	Volume int `json:"volume"`
}
