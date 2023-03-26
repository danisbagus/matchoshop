package v1

import (
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/auth"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service port.UserService
}

func NewUserhandler(service port.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h UserHandler) Login(c echo.Context) error {
	var loginRequest dto.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, appErr := h.service.Login(loginRequest)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	return c.JSON(http.StatusOK, *token)
}

func (h UserHandler) Refresh(c echo.Context) error {
	var refreshRequest dto.RefreshTokenRequest
	if err := c.Bind(&refreshRequest); err != nil {
		logger.Error("Error while decoding refresh token request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, appErr := h.service.Refresh(refreshRequest)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	return c.JSON(http.StatusOK, *token)
}

func (h UserHandler) RegisterCustomer(c echo.Context) error {
	var registerRequest dto.RegisterCustomerRequest
	if err := c.Bind(&registerRequest); err != nil {
		logger.Error("Error while decoding register customer request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, appErr := h.service.RegisterCustomer(&registerRequest)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())

	}

	return c.JSON(http.StatusOK, *token)
}

func (h UserHandler) GetUserDetail(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	userData, appErr := h.service.GetDetail(userInfo.UserID)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())

	}

	return c.JSON(http.StatusOK, *userData)
}

func (h UserHandler) UpdateUser(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	req := new(dto.UpdateUserRequest)

	if err := c.Bind(&req); err != nil {
		logger.Error("Error while decoding update user request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.User)
	form.UserID = userInfo.UserID
	form.Name = req.Name

	appErr = h.service.Update(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	res := dto.NewGetUserDetailResponse("Sucessfully update data", form)
	return c.JSON(http.StatusOK, res)

}

func (h UserHandler) GetUserList(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	roleID := userInfo.RoleID

	users, appErr := h.service.GetList(roleID)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())

	}

	resData := dto.NewGetUserListResponse("Successfully get data", users)
	return c.JSON(http.StatusOK, resData)
}

func (h UserHandler) DeleteUser(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	userID, _ := strconv.Atoi(c.Param("user_id"))
	roleID := userInfo.RoleID

	appErr := h.service.Delete(int64(userID), roleID)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())

	}

	res := dto.GenerateResponseData("Sucessfully delete data", nil)
	return c.JSON(http.StatusOK, res)
}

func (h UserHandler) UpdateUserAdmin(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	req := new(dto.UpdateUserRequest)

	if err := c.Bind(&req); err != nil {
		logger.Error("Error while decoding update user request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.User)
	form.UserID = int64(userID)
	form.Name = req.Name

	appErr = h.service.Update(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())

	}

	res := dto.NewGetUserDetailResponse("Sucessfully update data", form)
	return c.JSON(http.StatusOK, res)
}

func (h UserHandler) GetUserDetailAdmin(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	userData, appErr := h.service.GetDetail(int64(userID))
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())

	}

	return c.JSON(http.StatusOK, userData)
}
