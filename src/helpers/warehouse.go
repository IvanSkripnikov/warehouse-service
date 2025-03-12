package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"warehouse-service/models"

	"github.com/IvanSkripnikov/go-gormdb"
	logger "github.com/IvanSkripnikov/go-logger"
)

func GetWarehousesList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/warehouses/list"
	var warehouses []models.Warehouse

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Find(&warehouses).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": warehouses,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetWarehouse(w http.ResponseWriter, r *http.Request) {
	category := "/v1/warehouses/items-get"
	var warehouseItems []models.WarehouseItem

	WarehouseID, _ := getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if WarehouseID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Where("warehouse_id = ?", WarehouseID).Find(&warehouseItems).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": warehouseItems,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func BookItem(w http.ResponseWriter, r *http.Request) {
	category := "/v1/warehouses/book-item"

	// получаем параметры
	var bookingParams models.BookingItem
	err := json.NewDecoder(r.Body).Decode(&bookingParams)

	if checkError(w, err, category) {
		return
	}

	// смотрим, есть ли нужное количество на складе
	var warehouseItems []models.WarehouseItem

	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Where("item_id = ? AND volume >= ? AND status = ?", bookingParams.ItemID, bookingParams.Volume, models.StatusNew).Find(&warehouseItems).Error
	if checkError(w, err, category) {
		return
	}

	var result string
	if len(warehouseItems) > 0 {
		result = "success"
		currentTimestamp := int(GetCurrentTimestamp())

		// смотрим, есть ли уже забронированый товар
		var bookedItems []models.WarehouseItem
		err = db.Where("item_id = ? AND status = ?", bookingParams.ItemID, models.StatusBooked).Find(&bookedItems).Error
		if err != nil {
			result = "failure"
			logger.Errorf("Query get booked item error: %v", err)
		}

		var bookedItem models.WarehouseItem
		if len(bookedItems) > 0 {
			bookedItem = bookedItems[0]
			err = db.Model(&bookedItem).Update("volume", bookedItem.Volume+bookingParams.Volume).Error
			if err != nil {
				result = "failure"
				logger.Errorf("Cant create new booking item: %v", err)
			}
		} else {
			bookedItem.WarehouseID = warehouseItems[0].ID
			bookedItem.ItemID = bookingParams.ItemID
			bookedItem.Volume = bookingParams.Volume
			bookedItem.Created = currentTimestamp
			bookedItem.Updated = currentTimestamp
			bookedItem.Status = models.StatusBooked
			err = db.Create(&bookedItem).Error
			if err != nil {
				result = "failure"
				logger.Errorf("Cant create new booking item: %v", err)
			}
		}

		// уменьшаем незабронированый товар
		var newItems []models.WarehouseItem
		err = db.Where("item_id = ? AND status = ?", bookingParams.ItemID, models.StatusNew).Find(&newItems).Error
		if err != nil {
			result = "failure"
			logger.Errorf("Query get new item error: %v", err)
		}
		var newItem models.WarehouseItem
		if len(newItems) > 0 {
			newItem = newItems[0]
			if newItem.Volume-bookingParams.Volume < 0 {
				result = "failure"
				logger.Errorf("Incorrect volume requested: %v", err)
			} else {
				err = db.Model(&newItem).Update("volume", newItem.Volume-bookingParams.Volume).Error
				if err != nil {
					result = "failure"
					logger.Errorf("Cant create new booking item: %v", err)
				}
			}
		} else {
			newItem.WarehouseID = warehouseItems[0].ID
			newItem.ItemID = bookingParams.ItemID
			newItem.Volume = bookingParams.Volume
			newItem.Created = currentTimestamp
			newItem.Updated = currentTimestamp
			newItem.Status = models.StatusBooked
			err = db.Create(&newItem).Error
			if err != nil {
				result = "failure"
				logger.Errorf("Cant create new booking item: %v", err)
			}
		}
	} else {
		logger.Error("No items in warehouse")
		result = "failure"
	}

	err = db.Where("volume = ?", 0).Delete(&warehouseItems).Error
	if err != nil {
		result = "failure"
		logger.Errorf("Cant delete empty items error: %v", err)
	}

	data := ResponseData{
		"response": result,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func RollbackBook(w http.ResponseWriter, r *http.Request) {
	category := "/v1/warehouses/rollback-book"

	// получаем параметры
	var bookingParams models.BookingItem
	err := json.NewDecoder(r.Body).Decode(&bookingParams)
	if checkError(w, err, category) {
		return
	}

	result := "success"
	db := gormdb.GetClient(models.ServiceDatabase)
	var bookedItems []models.WarehouseItem

	// смотрим, есть ли уже забронированый товар
	err = db.Where("item_id = ? AND status = ?", bookingParams.ItemID, models.StatusBooked).Find(&bookedItems).Error
	if err != nil {
		result = "failure"
		logger.Errorf("Query get booked item error: %v", err)
	} else {
		// уменьшаем забронироанное количество
		if len(bookedItems) > 0 {
			bookedItem := bookedItems[0]
			if bookedItem.Volume-bookingParams.Volume < 0 {
				result = "failure"
				logger.Errorf("Incorrect volume requested: %v", err)
			} else {
				err = db.Model(&bookedItem).Update("volume", bookedItem.Volume-bookingParams.Volume).Error
				if checkError(w, err, category) {
					return
				}

				var newItems []models.WarehouseItem
				err = db.Where("item_id = ? AND status = ?", bookingParams.ItemID, models.StatusNew).Find(&newItems).Error
				if checkError(w, err, category) {
					return
				}
				if len(newItems) > 0 {
					newItem := newItems[0]

					err = db.Model(&newItem).Update("volume", newItem.Volume+bookingParams.Volume).Error
					if checkError(w, err, category) {
						return
					}
				} else {
					currentTimestamp := int(GetCurrentTimestamp())
					var newItem models.WarehouseItem
					newItem.WarehouseID = bookedItem.ID
					newItem.ItemID = bookedItem.ItemID
					newItem.Volume = bookedItem.Volume
					newItem.Created = currentTimestamp
					newItem.Updated = currentTimestamp
					newItem.Status = models.StatusNew
					err = db.Create(&newItem).Error
					if err != nil {
						result = "failure"
						logger.Errorf("Cant create new free item: %v", err)
					}
				}
			}
		} else {
			result = "failure"
			logger.Errorf("Query get booked item error: %v", err)
		}

		err = db.Where("volume = ?", 0).Delete(&bookedItems).Error
		if err != nil {
			result = "failure"
			logger.Errorf("Cant delete empty items error: %v", err)
		}
	}

	data := ResponseData{
		"response": result,
	}
	SendResponse(w, data, category, http.StatusOK)
}
