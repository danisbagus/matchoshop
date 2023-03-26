package v1

import (
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service port.ProductService
}

func NewProductHandler(sevice port.ProductService) *ProductHandler {
	return &ProductHandler{service: sevice}
}

func (h ProductHandler) CreateProduct(c echo.Context) error {
	var req dto.ProductRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("Error while decoding create product request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.Product)
	form.Name = req.Name
	form.Sku = req.Sku
	form.Brand = req.Brand
	form.Image = req.Image
	form.Description = req.Description
	form.Price = req.Price
	form.Stock = req.Stock
	form.ProductCategoryIDs = req.ProductCategoryIDs

	appErr = h.service.Create(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	res := dto.GenerateResponseData("Successfully create data", nil)
	return c.JSON(http.StatusOK, res)
}

func (h ProductHandler) GetTopProduct(c echo.Context) error {
	criteria := new(domain.ProductListCriteria)
	criteria.Limit = 3
	criteria.Sort = "numb_reviews"
	criteria.Order = "DESC"

	products, appErr := h.service.GetList(criteria)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	res := dto.NewGetProductListResponse("Successfully get data", products, nil)
	return c.JSON(http.StatusOK, res)
}

func (h ProductHandler) GetProductListPaginate(c echo.Context) error {
	req := new(dto.ProductListRequest)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	criteria := new(domain.ProductListCriteria)
	criteria.Keyword = req.Keyword
	criteria.Page, criteria.Limit = helper.SetPaginationParameter(req.Page, req.Limit)

	products, total, appErr := h.service.GetListPaginate(criteria)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	meta := new(helper.Meta)
	meta.SetPaginationData(criteria.Page, criteria.Limit, total)

	res := dto.NewGetProductListResponse("Successfully get data", products, meta)
	return c.JSON(http.StatusOK, res)
}

func (h ProductHandler) GetProductDetail(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	product, appErr := h.service.GetDetail(int64(productID))
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	res := dto.NewGetProductDetailResponse("Successfully get data", product)
	return c.JSON(http.StatusOK, res)
}

func (h ProductHandler) UpdateProduct(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("product_id"))
	var req dto.ProductRequest

	if err := c.Bind(&req); err != nil {
		logger.Error("Error while decoding update product request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.Product)
	form.Name = req.Name
	form.Sku = req.Sku
	form.Price = req.Price
	form.Brand = req.Brand
	form.Image = req.Image
	form.Stock = req.Stock
	form.Description = req.Description
	form.ProductCategoryIDs = req.ProductCategoryIDs

	appErr = h.service.Update(int64(productID), form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	res := dto.GenerateResponseData("Successfully update data", nil)
	return c.JSON(http.StatusOK, res)
}

func (h ProductHandler) Delete(c echo.Context) error {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	appErr := h.service.Delete(int64(productID))
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}
	res := dto.GenerateResponseData("Successfully delete data", nil)
	return c.JSON(http.StatusOK, res)
}
