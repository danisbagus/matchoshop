package v1

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/danisbagus/matchoshop/utils/constants"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	service port.ReviewService
}

func NewReviewHandler(sevice port.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: sevice}
}

func (h ReviewHandler) Create(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	var req dto.ReviewRequest

	err := c.Bind(&req)
	if err != nil {
		logger.Error("Error while decoding create review request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.Review)
	form.UserID = userInfo.UserID
	form.ProductID = req.ProductID
	form.Rating = req.Rating
	form.Comment = req.Comment
	appErr = h.service.Create(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.GenerateResponseData(constants.SuccessCreate, nil)
	return c.JSON(http.StatusOK, resData)
}

func (h ReviewHandler) GetDetail(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	productID := helper.StringToInt64(c.Param("product_id"), 0)

	review, appErr := h.service.GetDetail(userInfo.UserID, productID)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.NewReviewResponse(constants.SuccesGet, review)
	return c.JSON(http.StatusOK, resData)
}

func (h ReviewHandler) Update(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	var req dto.ReviewRequest

	err := c.Bind(&req)
	if err != nil {
		logger.Error("Error while decoding update review request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.Review)
	form.UserID = userInfo.UserID
	form.ProductID = req.ProductID
	form.Rating = req.Rating
	form.Comment = req.Comment

	appErr = h.service.Create(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}
	resData := dto.GenerateResponseData(constants.SuccessUpdate, nil)
	return c.JSON(http.StatusOK, resData)
}
