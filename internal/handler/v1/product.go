package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service port.ProductService
}

func (rc ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req domain.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Error while decoding create product request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product")
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
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

	appErr = rc.Service.Create(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.GenerateResponseData("Successfully create data", nil)
	response.Write(w, http.StatusCreated, res)
}

func (rc ProductHandler) GetTopProduct(w http.ResponseWriter, r *http.Request) {
	criteria := new(domain.ProductListCriteria)
	criteria.Limit = 3
	criteria.Sort = "numb_reviews"
	criteria.Order = "DESC"

	products, appErr := rc.Service.GetList(criteria)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetProductListResponse("Successfully get data", products, nil)
	response.Write(w, http.StatusOK, res)
}

func (rc ProductHandler) GetProductListPaginate(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	criteria := new(domain.ProductListCriteria)
	criteria.Keyword = keyword
	criteria.Page, criteria.Limit = helper.SetPaginationParameter(int64(page), int64(limit))
	products, total, appErr := rc.Service.GetListPaginate(criteria)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	meta := new(helper.Meta)
	meta.SetPaginationData(criteria.Page, criteria.Limit, total)

	res := domain.NewGetProductListResponse("Successfully get data", products, meta)
	response.Write(w, http.StatusOK, res)
}

func (rc ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])

	product, appErr := rc.Service.GetDetail(int64(productID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetProductDetailResponse("Successfully get data", product)
	response.Write(w, http.StatusOK, res)
}

func (rc ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])
	var req domain.ProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Error while decoding update product request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product category")
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
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

	appErr = rc.Service.Update(int64(productID), form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.GenerateResponseData("Successfully update data", nil)
	response.Write(w, http.StatusOK, res)
}

func (rc ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])
	appErr := rc.Service.Delete(int64(productID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	res := domain.GenerateResponseData("Successfully delete data", nil)
	response.Write(w, http.StatusOK, res)
}
