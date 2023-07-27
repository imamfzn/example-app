package handler

import (
	"errors"
	"net/http"

	"telefun/usecase"

	"github.com/gin-gonic/gin"
)

type PhoneNumberHandler struct {
	usecase *usecase.PhoneNumberUsecase
}

func NewPhoneNumberHandler(usecase *usecase.PhoneNumberUsecase) *PhoneNumberHandler {
	return &PhoneNumberHandler{usecase}
}

func (h *PhoneNumberHandler) Create(c *gin.Context) {
	req := CreatePhoneNumberRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.handleError(c, errors.New("invalid request body"))
		return
	}

	pn, err := h.usecase.Create(req.PhoneNumber)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, BuildCreatePhoneNumberResponse(pn))
}

func (h *PhoneNumberHandler) Available(c *gin.Context) {
	available, err := h.usecase.Available(c.Query("phone_number"))
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccess(c, AvailablePhoneNumberResponse{Available: available})
}

func (h *PhoneNumberHandler) handleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, BuildErrorResponse(err))
}

func (h *PhoneNumberHandler) handleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BuildSuccessResponse(data))
}
