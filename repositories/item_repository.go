package repositories

import (
	"gin-api/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updatedItem models.Item) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemRepository struct {
	db *gorm.DB
}

// Create implements IItemRepository.
func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// Delete implements IItemRepository.
func (r *ItemRepository) Delete(itemId uint, userId uint) error {
	deleteItem, err := r.FindById(itemId, userId)
	if err != nil {
		return err
	}

	result := r.db.Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements IItemRepository.
func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// FindById implements IItemRepository.
func (r *ItemRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, "id = ? AND user_id = ?", itemId, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

// Update implements IItemRepository.
func (r *ItemRepository) Update(updatedItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updatedItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedItem, nil
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}
