package v1

import (
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/labstack/echo/v4"
)

type ProductCategoryHandler struct {
	service port.ProductCategoryService
}

func NewProductCategoryHandler(sevice port.ProductCategoryService) *ProductCategoryHandler {
	return &ProductCategoryHandler{service: sevice}
}

func (h ProductCategoryHandler) CreateProductCategory(c echo.Context) error {
	var request dto.CreateProductCategoryRequest
	if err := c.Bind(&request); err != nil {
		logger.Error("Error while decoding create product category request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createData, appErr := h.service.Create(&request)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	return c.JSON(http.StatusOK, createData)
}

func (h ProductCategoryHandler) GetProductCategoryList(c echo.Context) error {
	productCategories, appErr := h.service.GetList()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	res := dto.NewGetProductCategoryListResponse("Successfully get data", productCategories)
	return c.JSON(http.StatusOK, res)
}

func (h ProductCategoryHandler) GetProductCategoryDetail(c echo.Context) error {
	productCategoryID, _ := strconv.Atoi(c.Param("product_category_id"))

	productCategory, appErr := h.service.GetDetail(int64(productCategoryID))
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	return c.JSON(http.StatusOK, productCategory)
}

func (h ProductCategoryHandler) UpdateProductCategory(c echo.Context) error {
	productCategoryID, _ := strconv.Atoi(c.Param("product_category_id"))
	var request dto.CreateProductCategoryRequest

	if err := c.Bind(&request); err != nil {
		logger.Error("Error while decoding update product category request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateData, appErr := h.service.Update(int64(productCategoryID), &request)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	return c.JSON(http.StatusOK, updateData)
}

func (h ProductCategoryHandler) Delete(c echo.Context) error {
	productCategoryID, _ := strconv.Atoi(c.Param("product_category_id"))

	deleteData, appErr := h.service.Delete(int64(productCategoryID))
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	return c.JSON(http.StatusOK, deleteData)
}
