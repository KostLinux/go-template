package controller

import (
	"errors"
	"fmt"
	"strconv"

	"go-template/model"
	request "go-template/model/http/request"
	response "go-template/model/http/response"
	constants "go-template/pkg/const"
	httpstatus "go-template/pkg/http/status_codes"
	"go-template/service"

	"github.com/gin-gonic/gin"
)

// User interface defines the user controller operations
type User interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) User {
	return &userController{service: service}
}

// @Summary		Create a new user
// @Description	Create a new user with the provided input
// @Tags			Users API
// @Accept			json
// @Produce		json
// @Param			user	body		request.CreateUser	true	"User details"
// @Success		201		{object}	response.User
// @Failure		400		{object}	response.UserErrorBadRequest
// @Failure		409		{object}	response.UserErrorConflict
// @Failure		500		{object}	response.UserErrorInternalServer
// @Router			/users [post]
func (ctrl *userController) Create(ctx *gin.Context) {
	var req request.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httpstatus.BadRequest(ctx, "invalid request", err)
		return
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := ctrl.service.Create(ctx, user); err != nil {
		if err.Error() == "user already exists" {
			httpstatus.Conflict(ctx, "user already exists", err)
			return
		}
		httpstatus.InternalServerError(ctx, "failed to create user", err)
		return
	}

	resp := response.UserMapper(user)
	httpstatus.Created(ctx, resp)
}

// Update updates an existing user
//
//	@Summary		Update a user
//	@Description	Update a user's information by ID
//	@Tags			Users API
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"User ID"
//	@Param			user	body		request.UpdateUser	true	"User details"
//	@Success		200		{object}	response.User
//
//	@Failure		400		{object}	response.UserErrorBadRequest
//	@Failure		404		{object}	response.UserErrorNotFound
//	@Failure		500		{object}	response.UserErrorInternalServer
//
//	@Router			/users/{id} [put]
func (ctrl *userController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		httpstatus.BadRequest(ctx, "invalid id for user", err)
		return
	}

	var req request.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		httpstatus.BadRequest(ctx, "invalid request", err)
		return
	}

	user := &model.User{
		ID:       uint(id),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := ctrl.service.Update(ctx, user); err != nil {
		switch {
		case errors.Is(err, constants.ErrNotFound):
			httpstatus.NotFound(ctx, "user not found", err)
		default:
			httpstatus.InternalServerError(ctx, "failed to update user", err)
		}
		return
	}

	resp := response.UserMapper(user)
	httpstatus.OK(ctx, resp)
}

// Delete removes a user
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			Users API
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		204	{object}	response.DeleteUser
//
//	@Failure		400	{object}	response.UserErrorBadRequest
//	@Failure		404	{object}	response.UserErrorNotFound
//	@Failure		500	{object}	response.UserErrorInternalServer
//
//	@Router			/users/{id} [delete]
func (ctrl *userController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		httpstatus.BadRequest(ctx, "invalid id for user", err)
		return
	}

	if err := ctrl.service.Delete(ctx, uint(id)); err != nil {
		switch {
		case errors.Is(err, constants.ErrNotFound):
			httpstatus.NotFound(ctx, "user not found", err)
		default:
			httpstatus.InternalServerError(ctx, "failed to delete user", err)
		}
		return
	}

	resp := response.DeleteUser{
		Message: fmt.Sprintf("user %d deleted successfully", id),
	}
	httpstatus.OK(ctx, resp)
}

// @Summary		List all users
// @Description	Get all users in the system
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.UserList
// @Failure		500	{object}	response.UserErrorInternalServer
// @Router			/users [get]
func (ctrl *userController) Get(ctx *gin.Context) {
	users, err := ctrl.service.List(ctx)
	if err != nil {
		httpstatus.InternalServerError(ctx, "failed to get users", err)
		return
	}

	resp := response.UserListMapper(users)
	httpstatus.OK(ctx, resp)
}

// @Summary		Get a user by ID
// @Description	Get detailed information about a specific user
// @Tags			Users API
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	response.User
// @Failure		400	{object}	response.UserErrorBadRequest
// @Failure		404	{object}	response.UserErrorNotFound
// @Failure		500	{object}	response.UserErrorInternalServer
// @Router			/users/{id} [get]
func (ctrl *userController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		httpstatus.BadRequest(ctx, "invalid id for user", err)
		return
	}

	user, err := ctrl.service.GetByID(ctx, uint(id))
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrNotFound):
			httpstatus.NotFound(ctx, "user not found", err)
		default:
			httpstatus.InternalServerError(ctx, "failed to get user", err)
		}
		return
	}

	resp := response.UserMapper(user)
	httpstatus.OK(ctx, resp)
}
