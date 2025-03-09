package models

type Warehouse struct {
	ID      int    `gorm:"index;type:int" json:"id"`
	Title   string `gorm:"type:text" json:"title"`
	Volume  int    `gorm:"type:int" json:"volume"`
	Created int    `gorm:"index;type:bigint" json:"created"`
	Updated int    `gorm:"index;type:bigint" json:"updated"`
}
