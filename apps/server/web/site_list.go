package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *SiteController) HandleList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	//TODO: maybe auto sign uid params for normal user?

	uid, err := strconv.Atoi(ctx.Query("uid"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	sites, err := c.usecase.List(ctx, GetUserOp(ctx), page, pageSize, uid)
	if err != nil {
		c.log.Error("Failed to retrived site data", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.Int("uid", uid), zap.Error(err))
		switch {
		case errors.Is(err, usecase.ItemNotExistedError{}):
			ctx.Status(http.StatusNotFound)
		case errors.Is(err, perm.PermissionError{}):
			ctx.Status(http.StatusForbidden)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		return
	}

	result := mappingFunc(sites, c.ToDto)
	// var result []Site
	// for _, a := range sites {
	// 	result = append(result, *c.ToDto(a))
	// }

	ctx.JSON(http.StatusOK, result)
}
