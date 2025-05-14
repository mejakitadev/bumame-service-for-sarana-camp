package controller

import (
	"math"
	"net/http"
	"sarana-dafa-ai-service/helper"
	"sarana-dafa-ai-service/model/web"
	"sarana-dafa-ai-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/**
* Interface design
**/
type BumameB2BDokterController interface {
	// Public
	FindAll(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	FindByName(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}

/**
* Object implementation creation
* And inject depedencies
**/
type BumameB2BDokterControllerImpl struct {
	BumameB2BDokterService service.BumameB2BDokterService
	Validate               *validator.Validate
}

func NewBumameB2BDokterController(buas service.BumameB2BDokterService, valid *validator.Validate) BumameB2BDokterController {
	return &BumameB2BDokterControllerImpl{
		BumameB2BDokterService: buas,
		Validate:               valid,
	}
}

// @Summary		Get all b2b dokter
// @Description	Get all b2b dokter
// @Tags			b2b-dokter
// @Param			page		query	int					false	"Page"
// @Param			limit		query	int					false	"Limit"
// @Param			sort		query	string				false	"Sort"
// @Param			queryList	query	web.QueryAdminB2B	false	"Query List"
// @Accept			json
// @Produce		json
// @Success		200	{object}	web.ResponseStandard
// @Failure		500	{object}	web.ResponseStandard
// @Router			/b2b-dokter [get]
func (cont *BumameB2BDokterControllerImpl) FindAll(c *fiber.Ctx) error {
	reqPagination := new(web.RequestPaginationNumber)
	if err := c.QueryParser(reqPagination); err != nil {
		return err
	}

	// Gunakan data statis langsung
	data := []map[string]interface{}{
		{"name": "Dr. Andi", "visits": 8},
		{"name": "Dr. Budi", "visits": 10},
		{"name": "Dr. Citra", "visits": 6},
		{"name": "Dr. Dewi", "visits": 9},
		{"name": "Dr. Eko", "visits": 7},
		{"name": "Dr. Gita", "visits": 12},
		{"name": "Dr. Hadi", "visits": 5},
	}

	// Buat pagination dari data statis
	pagination := helper.PaginationNumber{
		Type:       "number",
		Page:       reqPagination.Page,
		PerPage:    reqPagination.Limit,
		TotalRows:  int64(len(data)),
		TotalPages: int(math.Ceil(float64(len(data)) / float64(reqPagination.Limit))),
		Sort:       reqPagination.Sort,
	}

	// Langsung return data statis
	return helper.BuildSuccessResponsePagination(c, http.StatusOK, data, pagination)
}

// @Summary		Get b2b dokter by ID
// @Description	Get b2b dokter by ID
// @Tags			b2b-dokter
// @Param			b2b_dokter_id	path	string	true	"B2B dokter ID"
// @Accept			json
// @Produce		json
// @Success		200	{object}	web.ResponseStandard
// @Failure		404	{object}	web.ResponseStandard
// @Failure		500	{object}	web.ResponseStandard
// @Router			/b2b-dokter/{b2b_dokter_id} [get]
func (cont *BumameB2BDokterControllerImpl) FindById(c *fiber.Ctx) error {
	b2bDokterId := c.Params("b2b_dokter_id")

	data, err := cont.BumameB2BDokterService.FindById(c.Context(), b2bDokterId)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusNotFound, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// coba coba get produk by name
//
//	@Summary		Get b2b dokter by name
//	@Description	Get b2b dokter by name
//	@Tags			b2b-dokter
//	@Param			b2b_dokter_name	path	string	true	"B2B dokter name"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ResponseStandard
//	@Failure		404	{object}	web.ResponseStandard
//	@Failure		500	{object}	web.ResponseStandard
//	@Router			/b2b-dokter/name/{b2b_dokter_name} [get]
func (cont *BumameB2BDokterControllerImpl) FindByName(c *fiber.Ctx) error {
	b2bDokterName := c.Params("b2b_dokter_name")

	data, err := cont.BumameB2BDokterService.FindByName(c.Context(), b2bDokterName)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusNotFound, []web.ErrorResponse{{
			Title:    "dokter Not Found by Name",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary		Update b2b dokter
// @Description	Update b2b dokter
// @Tags			b2b-dokter
// @Accept			json
// @Produce		json
// @Param			b2b_dokter_id	path		string					true	"B2B dokter ID"
// @Param			request			body		web.RequestB2BDokter	true	"Request"  // Fix: Changed from web.RequestB2Bdokter to web.RequestB2Bdokter
// @Success		200				{object}	web.ResponseStandard
// @Failure		400				{object}	web.ResponseStandard
// @Failure		500				{object}	web.ResponseStandard
// @Router			/b2b-dokter/{b2b_dokter_id} [put]
func (cont *BumameB2BDokterControllerImpl) Update(c *fiber.Ctx) error {
	b2bDokterId := c.Params("b2b_dokter_id")

	request := new(web.RequestB2BDokter)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	data, err := cont.BumameB2BDokterService.Update(c.Context(), b2bDokterId, *request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data) // Changed from nil to data
}

// @Summary		Delete b2b dokter
// @Description	Delete b2b dokter
// @Tags			b2b-dokter
// @Accept			json
// @Produce		json
// @Param			b2b_dokter_id	path		string	true	"B2B dokter ID"
// @Success		200				{object}	web.ResponseStandard
// @Failure		400				{object}	web.ResponseStandard
// @Failure		500				{object}	web.ResponseStandard
// @Router			/b2b-dokter/{b2b_dokter_id} [delete]
func (cont *BumameB2BDokterControllerImpl) Delete(c *fiber.Ctx) error {
	b2bDokterId := c.Params("b2b_dokter_id")

	data, err := cont.BumameB2BDokterService.Delete(c.Context(), b2bDokterId)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary        Create b2b dokter
// @Description    Create b2b dokter
// @Tags           b2b-dokter
// @Accept         json
// @Produce        json
// @Param          request body        web.RequestB2BDokter    true    "Request"
// @Success        201     {object}    web.ResponseStandard
// @Failure        400     {object}    web.ResponseStandard
// @Failure        500     {object}    web.ResponseStandard
// @Router         /b2b-dokter [post]
// Fix the Create method to use the correct service method
func (cont *BumameB2BDokterControllerImpl) Create(c *fiber.Ctx) error {
	request := new(web.RequestB2BDokter)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := cont.BumameB2BDokterService.Create(c.Context(), *request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusCreated, data)
}
