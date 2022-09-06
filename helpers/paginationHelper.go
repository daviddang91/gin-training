package helpers

import (
	"fmt"
	"gin-training/database"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Param struct {
	Model   interface{}
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

type Pagination struct {
	Count    int64 `json:"count"`
	Pages    int   `json:"pages"`
	Offset   int   `json:"offset"`
	Limit    int   `json:"limit"`
	Page     int   `json:"page"`
	PrevPage int   `json:"prev_page"`
	NextPage int   `json:"next_page"`
}

func GeneratePagination(ctx *gin.Context, model interface{}) Param {
	orderBy := GetOrderParam(ctx)
	page, limit := GetPageParam(ctx)
	return Param{
		Model:   model,
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
	}
}

func GetOrderParam(c *gin.Context) []string {
	var orderBy []string
	orders := c.DefaultQuery("order", "-created_at")
	for _, order := range strings.Split(orders, ",") {
		order = strings.TrimSpace(order)
		if strings.Contains(order, "-") {
			order = strings.Replace(order, "-", "", -1)
			orderBy = append(orderBy, fmt.Sprintf("%s DESC", order))
		} else {
			orderBy = append(orderBy, fmt.Sprintf("%s ASC", order))
		}
	}
	return orderBy
}

func GetPageParam(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))
	return page, limit
}

func GetIntParam(key string, c *gin.Context) uint {
	val := c.Param(key)
	id, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return uint(id)
}

func Paginator(p *Param, data interface{}) (*Pagination, error) {
	queryBuider := database.Instance.Model(p.Model)
	offset := (p.Page - 1) * p.Limit

	if p.ShowSQL {
		queryBuider = queryBuider.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 30
	}
	if len(p.OrderBy) > 0 {
		for _, order := range p.OrderBy {
			queryBuider = queryBuider.Order(order)
		}
	}

	var pagination Pagination
	var count int64

	if err := queryBuider.Count(&count).Error; err != nil {
		return nil, err
	}

	pagination.Count = count
	pagination.Page = p.Page
	pagination.Offset = offset
	pagination.Limit = p.Limit
	pagination.Pages = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		pagination.PrevPage = p.Page - 1
	} else {
		pagination.PrevPage = 1
	}

	if p.Page >= pagination.Pages {
		pagination.NextPage = pagination.Pages
	} else {
		pagination.NextPage = p.Page + 1
	}

	if err := queryBuider.Limit(p.Limit).Offset(offset).Find(data).Error; err != nil {
		return nil, err
	}

	return &pagination, nil
}
