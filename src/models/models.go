package models

const ServiceDatabase = "WarehouseService"

type BookingItem struct {
	ItemID int `json:"id"`
	Volume int `json:"volume"`
}
