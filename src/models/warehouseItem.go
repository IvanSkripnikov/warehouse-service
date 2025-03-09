package models

const StatusNew = 1
const StatusBooked = 2

type WarehouseItem struct {
	ID          int `gorm:"index;type:int" json:"id"`
	WarehouseID int `gorm:"index;type:int" json:"warehouseId"`
	ItemID      int `gorm:"index;type:int" json:"itemId"`
	Volume      int `gorm:"type:int" json:"volume"`
	Created     int `gorm:"index;type:bigint" json:"created"`
	Updated     int `gorm:"index;type:bigint" json:"updated"`
	Status      int `gorm:"index;type:int" json:"status"`
}

func (s Warehouse) TableName() string { return "warehouse_items" }
