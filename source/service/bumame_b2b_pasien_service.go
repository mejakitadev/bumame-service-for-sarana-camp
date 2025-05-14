package service

import (
	"context"
	"errors"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model"
	"sarana-dafa-ai-service/model/web"

	"gorm.io/gorm"
)

/**
* Interface
**/
type BumameB2BPasienService interface {
	FindAllPagination(ctx context.Context, reqPage web.RequestPaginationNumber) ([]map[string]interface{}, helper.PaginationNumber)
	FindById(ctx context.Context, pasienId string) (map[string]interface{}, error)
	FindByName(ctx context.Context, pasienName string) (map[string]interface{}, error)
	Create(ctx context.Context, request web.RequestB2BPasien) (map[string]interface{}, error)
	Update(ctx context.Context, pasienId string, request web.RequestB2BPasien) (map[string]interface{}, error)
	Delete(ctx context.Context, pasienId string) (map[string]interface{}, error)
}

/**
* Object implementation creation
**/
type BumameB2BPasienServiceImpl struct {
	DB *gorm.DB
}

func NewBumameB2BPasienService(db *gorm.DB) BumameB2BPasienService {
	return &BumameB2BPasienServiceImpl{
		DB: db,
	}
}

func (service BumameB2BPasienServiceImpl) FindAllPagination(
	ctx context.Context, reqPage web.RequestPaginationNumber) (
	[]map[string]interface{}, helper.PaginationNumber) {
	listData := []model.B2BPasien{}
	tx := service.DB.Begin()
	defer tx.Rollback()

	pagination := helper.PaginationNumber{}
	pagination.Page = reqPage.Page
	pagination.PerPage = reqPage.Limit
	pagination.Sort = reqPage.Sort

	tx = tx.Where("is_deleted = 0")
	err := tx.Scopes(helper.Paginate(listData, &pagination, tx)).Find(&listData).Error
	helper.PanicIfError(err)

	tx.Commit()

	responseData := []map[string]interface{}{}
	for _, v := range listData {
		response := map[string]interface{}{
			"id":          v.Id,
			"pasien_name": v.PasienName,
			"created_at":  v.CreatedAt,
			"updated_at":  v.UpdatedAt,
		}
		responseData = append(responseData, response)
	}
	return responseData, pagination
}

func (service BumameB2BPasienServiceImpl) FindById(ctx context.Context, PasienId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Pasien := model.B2BPasien{}
	err := tx.Where("id = ? AND is_deleted = 0", PasienId).First(&Pasien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Pasien not found")
		}
		helper.PanicIfError(err)
	}

	tx.Commit()

	return map[string]interface{}{
		"id":          Pasien.Id,
		"pasien_name": Pasien.PasienName,
		"created_at":  Pasien.CreatedAt,
		"updated_at":  Pasien.UpdatedAt,
	}, nil
}

// Update the FindByName method signature
func (service BumameB2BPasienServiceImpl) FindByName(ctx context.Context, PasienName string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Pasien := model.B2BPasien{}
	// Change the query to search by name instead of id
	err := tx.Where("pasien_name LIKE ? AND is_deleted = 0", "%"+PasienName+"%").First(&Pasien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Pasien not found")
		}
		helper.PanicIfError(err)
	}

	tx.Commit()

	return map[string]interface{}{
		"id":          Pasien.Id,
		"pasien_name": Pasien.PasienName,
		"created_at":  Pasien.CreatedAt,
		"updated_at":  Pasien.UpdatedAt,
	}, nil
}

func (service BumameB2BPasienServiceImpl) Create(ctx context.Context, request web.RequestB2BPasien) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Pasien := model.B2BPasien{
		PasienName: request.PasienName,
	}

	err := tx.Create(&Pasien).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":         Pasien.Id,
		"name":       Pasien.PasienName,
		"created_at": Pasien.CreatedAt,
	}, nil
}

func (service BumameB2BPasienServiceImpl) Update(ctx context.Context, PasienId string, request web.RequestB2BPasien) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Pasien := model.B2BPasien{}
	err := tx.Where("id = ? AND is_deleted = 0", PasienId).First(&Pasien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Pasien not found")
		}
		helper.PanicIfError(err)
	}

	Pasien.PasienName = request.PasienName
	err = tx.Save(&Pasien).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":          Pasien.Id,
		"pasien_name": Pasien.PasienName,
		"created_at":  Pasien.CreatedAt,
		"updated_at":  Pasien.UpdatedAt,
	}, nil
}

func (service BumameB2BPasienServiceImpl) Delete(ctx context.Context, PasienId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Pasien := model.B2BPasien{}
	err := tx.Where("id = ? AND is_deleted = 0", PasienId).First(&Pasien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Pasien not found")
		}
		helper.PanicIfError(err)
	}

	Pasien.IsDeleted = 1
	err = tx.Save(&Pasien).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"message": "Pasien deleted successfully",
	}, nil
}
