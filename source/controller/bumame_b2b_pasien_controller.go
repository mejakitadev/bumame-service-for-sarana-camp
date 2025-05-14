package controller

import (
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
type BumameB2BPasienController interface {
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
type BumameB2BPasienControllerImpl struct {
	BumameB2BPasienService service.BumameB2BPasienService
	Validate               *validator.Validate
}

func NewBumameB2BPasienController(buas service.BumameB2BPasienService, valid *validator.Validate) BumameB2BPasienController {
	return &BumameB2BPasienControllerImpl{
		BumameB2BPasienService: buas,
		Validate:               valid,
	}
}

// @Summary		Get all b2b pasien
// @Description	Get all b2b pasien
// @Tags			b2b-pasien
// @Param			page		query	int					false	"Page"
// @Param			limit		query	int					false	"Limit"
// @Param			sort		query	string				false	"Sort"
// @Param			queryList	query	web.QueryAdminB2B	false	"Query List"
// @Accept			json
// @Produce		json
// @Success		200	{object}	web.ResponseStandard
// @Failure		500	{object}	web.ResponseStandard
// @Router			/b2b-pasien [get]
func (cont *BumameB2BPasienControllerImpl) FindAll(c *fiber.Ctx) error {
	reqPagination := new(web.RequestPaginationNumber)
	if err := c.QueryParser(reqPagination); err != nil {
		return err
	}

	data, pagination := cont.BumameB2BPasienService.FindAllPagination(c.Context(), *reqPagination)
	return helper.BuildSuccessResponsePagination(c, http.StatusOK, data, pagination)
}

// @Summary		Get b2b pasien by ID
// @Description	Get b2b pasien by ID
// @Tags			b2b-pasien
// @Param			b2b_pasien_id	path	string	true	"B2B pasien ID"
// @Accept			json
// @Produce		json
// @Success		200	{object}	web.ResponseStandard
// @Failure		404	{object}	web.ResponseStandard
// @Failure		500	{object}	web.ResponseStandard
// @Router			/b2b-pasien/{b2b_pasien_id} [get]
func (cont *BumameB2BPasienControllerImpl) FindById(c *fiber.Ctx) error {
	b2bpasienId := c.Params("b2b_pasien_id")

	data, err := cont.BumameB2BPasienService.FindById(c.Context(), b2bpasienId)
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
//	@Summary		Get b2b pasien by name
//	@Description	Get b2b pasien by name
//	@Tags			b2b-pasien
//	@Param			b2b_pasien_name	path	string	true	"B2B pasien name"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ResponseStandard
//	@Failure		404	{object}	web.ResponseStandard
//	@Failure		500	{object}	web.ResponseStandard
//	@Router			/b2b-pasien/name/{b2b_pasien_name} [get]
func (cont *BumameB2BPasienControllerImpl) FindByName(c *fiber.Ctx) error {
	b2bPasienName := c.Params("b2b_pasien_name")

	data, err := cont.BumameB2BPasienService.FindByName(c.Context(), b2bPasienName)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusNotFound, []web.ErrorResponse{{
			Title:    "pasien Not Found by Name",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary		Update b2b pasien
// @Description	Update b2b pasien
// @Tags			b2b-pasien
// @Accept			json
// @Produce		json
// @Param			b2b_pasien_id	path		string					true	"B2B pasien ID"
// @Param			request			body		web.RequestB2BPasien	true	"Request"  // Fix: Changed from web.RequestB2Bpasien to web.RequestB2BPasien
// @Success		200				{object}	web.ResponseStandard
// @Failure		400				{object}	web.ResponseStandard
// @Failure		500				{object}	web.ResponseStandard
// @Router			/b2b-pasien/{b2b_pasien_id} [put]
func (cont *BumameB2BPasienControllerImpl) Update(c *fiber.Ctx) error {
	b2bPasienId := c.Params("b2b_pasien_id")

	request := new(web.RequestB2BPasien)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	data, err := cont.BumameB2BPasienService.Update(c.Context(), b2bPasienId, *request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary		Delete b2b pasien
// @Description	Delete b2b pasien
// @Tags			b2b-pasien
// @Accept			json
// @Produce		json
// @Param			b2b_pasien_id	path		string	true	"B2B pasien ID"
// @Success		200				{object}	web.ResponseStandard
// @Failure		400				{object}	web.ResponseStandard
// @Failure		500				{object}	web.ResponseStandard
// @Router			/b2b-pasien/{b2b_pasien_id} [delete]
func (cont *BumameB2BPasienControllerImpl) Delete(c *fiber.Ctx) error {
	b2bPasienId := c.Params("b2b_pasien_id")

	data, err := cont.BumameB2BPasienService.Delete(c.Context(), b2bPasienId)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusOK, data)
}

// @Summary        Create b2b pasien
// @Description    Create b2b pasien
// @Tags           b2b-pasien
// @Accept         json
// @Produce        json
// @Param          request body        web.RequestB2BPasien    true    "Request"
// @Success        201     {object}    web.ResponseStandard
// @Failure        400     {object}    web.ResponseStandard
// @Failure        500     {object}    web.ResponseStandard
// @Router         /b2b-pasien [post]
// Fix the Create method to use the correct service method
func (cont *BumameB2BPasienControllerImpl) Create(c *fiber.Ctx) error {
	request := new(web.RequestB2BPasien)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	data, err := cont.BumameB2BPasienService.Create(c.Context(), *request)
	if err != nil {
		return helper.BuildErrorResponse(c, http.StatusInternalServerError, []web.ErrorResponse{{
			Title:    "Internal Server Error",
			Messages: []string{err.Error()},
		}})
	}
	return helper.BuildSuccessResponse(c, http.StatusCreated, data)
}
