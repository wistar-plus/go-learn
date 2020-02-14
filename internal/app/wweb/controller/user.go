package controller

import (
	"go-learn/internal/app/wweb/model"
	"go-learn/internal/app/wweb/persistence/orm"
	"go-learn/internal/app/wweb/service"
	"go-learn/internal/pkg/e"
	"go-learn/internal/pkg/gintool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController {
	repo := orm.NewUserRepository()
	return &UserController{
		service: service.NewUserService(repo),
	}
}

func (uc *UserController) Create(c *gin.Context) {
	var user *model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	err = uc.service.Create(user)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}

	gintool.ResSuccess(c, nil)
}

func (uc *UserController) Update(c *gin.Context) {
	var user *model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	err = uc.service.Update(user)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}

	gintool.ResSuccess(c, nil)
}

func (uc *UserController) FindByPage(c *gin.Context) {
	page, size := gintool.GetPaginationParam(c)
	list, err := uc.service.FindByPage(page, size)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}
	gintool.ResSuccess(c, list)
}

func (uc *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}
	err = uc.service.Delete(uint32(id))
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}
	gintool.ResSuccess(c, nil)
}

func (uc *UserController) Login(c *gin.Context) {

	var param *model.UserLoginParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	user, err := uc.service.Login(param)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	gintool.ResSuccess(c, map[string]interface{}{
		"loginUser": user,
	})

}
