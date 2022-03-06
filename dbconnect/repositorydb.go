package dbconnect

import (
	"fmt"
	"go-be-app/model"
	"strconv"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Phone, error)
	FindByDisplay(display string) ([]model.Phone, error)
	PostPhone(phone model.Phone) error
	UpdatePhone(phone model.Phone, id string) (model.Phone, error)
	DeletePhone(id string) error
}

type repository struct {
	db_t *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Phone, error) {
	var phone_list []model.Phone

	sql_db, _ := r.db_t.DB()

	data := r.db_t.Find(&phone_list)
	if data.Error != nil {
		fmt.Println("Error: ", data.Error)
		sql_db.Close()
		return phone_list, data.Error
	}

	sql_db.Close()

	return phone_list, data.Error

}

func (r *repository) FindByDisplay(display string) ([]model.Phone, error) {
	var phone_data []model.Phone

	sql_db, _ := r.db_t.DB()

	data := r.db_t.Where("display = ?", display).Find(&phone_data)
	if data.Error != nil {
		fmt.Println("Error: ", data.Error)
		sql_db.Close()
		return phone_data, data.Error
	}

	sql_db.Close()

	return phone_data, data.Error

}

func (r *repository) PostPhone(phone model.Phone) error {
	var phone_entity model.Phone
	phone_entity.Brand = phone.Brand
	phone_entity.Display = phone.Display
	phone_entity.Discount = phone.Discount
	phone_entity.Rating = phone.Rating
	phone_entity.Price = phone.Price

	sql_db, _ := r.db_t.DB()

	data := r.db_t.Create(&phone_entity)
	if data.Error != nil {
		fmt.Println(data.Error)
		sql_db.Close()
		return data.Error
	}

	sql_db.Close()

	return data.Error

}

func (r *repository) UpdatePhone(phone_update model.Phone, id string) (model.Phone, error) {
	var phone_entity model.Phone

	id_phone, _ := strconv.Atoi(id)
	phone_update.Id = id_phone

	sql_db, _ := r.db_t.DB()

	err_search_id := r.db_t.Where("id = ?", id_phone).First(&phone_entity)
	if err_search_id.Error != nil || err_search_id.RowsAffected <= 0 {
		fmt.Println("Failed! msg: ", err_search_id)
		sql_db.Close()
		return model.Phone{}, err_search_id.Error
	}

	phone_entity.Id = phone_update.Id
	phone_entity.Brand = phone_update.Brand
	phone_entity.Display = phone_update.Display
	phone_entity.Discount = phone_update.Discount
	phone_entity.Rating = phone_update.Rating
	phone_entity.Price = phone_update.Price

	// get created data
	err_db_created := r.db_t.Select("created_at").Where("id = ?", id_phone).Find(&phone_update)

	if err_db_created.Error != nil {
		fmt.Println("Failed to find created_at!", err_db_created)
		sql_db.Close()
		return model.Phone{}, err_db_created.Error
	}

	phone_entity.CreatedAt = phone_update.CreatedAt

	// put to db
	err_db := r.db_t.Where("id = ?", id_phone).Save(&phone_entity)

	if err_db.Error != nil {
		fmt.Println("Failed to update!", err_db)
		sql_db.Close()
		return model.Phone{}, err_db.Error
	}

	sql_db.Close()

	return phone_entity, err_db.Error

}

func (r *repository) DeletePhone(id string) error {
	var phone_delete model.Phone

	sql_db, _ := r.db_t.DB()

	phone_id, _ := strconv.Atoi(id)

	// find id
	data_id := r.db_t.Where("id = ?", phone_id).First(&phone_delete)
	if data_id.Error != nil || data_id.RowsAffected <= 0 {
		fmt.Println("Error data find: ", data_id.Error)
		sql_db.Close()
		return data_id.Error
	}

	data_delete := r.db_t.Where("id = ?", phone_id).Delete(&phone_delete)
	if data_delete.Error != nil {
		fmt.Println("Error data delete: ", data_delete.Error)
		sql_db.Close()
		return data_delete.Error
	}

	return data_delete.Error

}
