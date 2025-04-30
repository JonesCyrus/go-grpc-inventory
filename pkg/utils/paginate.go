package utils

import (
	"go-grpc-inventory/pkg/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GeneratePaginationModel(c *gin.Context) *models.Pagination {
	pagination := &models.Pagination{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		size = 5
	}

	pagination.Page = page
	pagination.Size = size

	return pagination
}

func Paginate(p *models.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var page, size int

		if p.Page > 0 {
			page = p.Page
		} else {
			// If Page is not provided or invalid, set default value
			page = 1
		}

		if p.Size > 0 && p.Size <= 100 {
			size = p.Size
		} else {
			// If PageSize is not provided or invalid, set default value
			size = 5
		}

		// Convert int64 to int for Offset
		offset := int((page - 1) * size)

		// Apply pagination
		return db.Offset(offset).Limit(size)
	}
}

func Populate(val string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(val) == 0 {
			return db
		}

		parts := strings.Split(val, ",")

		//separeate the strings with comma

		for _, str := range parts {
			trimmedStr := strings.TrimSpace(str)
			db = db.Preload(trimmedStr)
		}
		return db
	}
}
