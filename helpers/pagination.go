package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		if page <= 0 {
			page = 1
		}

		limit, _:= strconv.Atoi(ctx.DefaultQuery("limit", "10")) 
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}