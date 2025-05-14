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
type BumameB2BDokterService interface {
	FindAllPagination(ctx context.Context, reqPage web.RequestPaginationNumber) ([]map[string]interface{}, helper.PaginationNumber)
	FindById(ctx context.Context, DokterId string) (map[string]interface{}, error)
	FindByName(ctx context.Context, DokterName string) (map[string]interface{}, error)
	Create(ctx context.Context, request web.RequestB2BDokter) (map[string]interface{}, error)
	Update(ctx context.Context, DokterId string, request web.RequestB2BDokter) (map[string]interface{}, error)
	Delete(ctx context.Context, DokterId string) (map[string]interface{}, error)
}

/**
* Object implementation creation
**/
type BumameB2BDokterServiceImpl struct {
	DB *gorm.DB
}

// Add this function at the top level of your service file
func NewBumameB2BDokterService(db *gorm.DB) BumameB2BDokterService {
	return &BumameB2BDokterServiceImpl{
		DB: db,
	}
}

func (service BumameB2BDokterServiceImpl) FindAllPagination(
	ctx context.Context, reqPage web.RequestPaginationNumber) (
	[]map[string]interface{}, helper.PaginationNumber) {
	listData := []model.B2BDokter{}
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
			"Dokter_name": v.DokterName,
			"created_at":  v.CreatedAt,
			"updated_at":  v.UpdatedAt,
		}
		responseData = append(responseData, response)
	}
	return responseData, pagination
}

func (service BumameB2BDokterServiceImpl) FindById(ctx context.Context, DokterId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Dokter := model.B2BDokter{}
	err := tx.Where("id = ? AND is_deleted = 0", DokterId).First(&Dokter).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Dokter not found")
		}
		helper.PanicIfError(err)
	}

	tx.Commit()

	return map[string]interface{}{
		"id":          Dokter.Id,
		"dokter_name": Dokter.DokterName,
		"created_at":  Dokter.CreatedAt,
		"updated_at":  Dokter.UpdatedAt,
	}, nil
}

func (service BumameB2BDokterServiceImpl) FindByName(ctx context.Context, DokterName string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Dokter := model.B2BDokter{}
	err := tx.Where("dokter_name LIKE ? AND is_deleted = 0", "%"+DokterName+"%").First(&Dokter).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Dokter not found")
		}
		helper.PanicIfError(err)
	}

	tx.Commit()

	return map[string]interface{}{
		"id":          Dokter.Id,
		"dokter_name": Dokter.DokterName,
		"created_at":  Dokter.CreatedAt,
		"updated_at":  Dokter.UpdatedAt,
	}, nil
}

func (service *BumameB2BDokterServiceImpl) Create(ctx context.Context, request web.RequestB2BDokter) (map[string]interface{}, error) {
    tx := service.DB.Begin()
    defer tx.Rollback()

    dokter := model.B2BDokter{
        DokterName: request.DokterName,
    }

    // Changed: Make sure it's using b2b_dokter table
    err := tx.Create(&dokter).Error
    if err != nil {
        return nil, err
    }

    tx.Commit()

    return map[string]interface{}{
        "id":          dokter.Id,
        "dokter_name": dokter.DokterName,
        "created_at":  dokter.CreatedAt,
    }, nil
}

func (service BumameB2BDokterServiceImpl) Update(ctx context.Context, DokterId string, request web.RequestB2BDokter) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Dokter := model.B2BDokter{}
	err := tx.Where("id = ? AND is_deleted = 0", DokterId).First(&Dokter).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Dokter not found")
		}
		helper.PanicIfError(err)
	}

	Dokter.DokterName = request.DokterName
	err = tx.Save(&Dokter).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"id":          Dokter.Id,
		"dokter_name": Dokter.DokterName,
		"created_at":  Dokter.CreatedAt,
		"updated_at":  Dokter.UpdatedAt,
	}, nil
}

func (service BumameB2BDokterServiceImpl) Delete(ctx context.Context, DokterId string) (map[string]interface{}, error) {
	tx := service.DB.Begin()
	defer tx.Rollback()

	Dokter := model.B2BDokter{}
	err := tx.Where("id = ? AND is_deleted = 0", DokterId).First(&Dokter).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Dokter not found")
		}
		helper.PanicIfError(err)
	}

	Dokter.IsDeleted = 1
	err = tx.Save(&Dokter).Error
	helper.PanicIfError(err)

	tx.Commit()

	return map[string]interface{}{
		"message": "Dokter deleted successfully",
	}, nil
}
