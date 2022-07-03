package action

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Action struct {
	Code    int         `json:"code,omitempty"`
	Total   int         `json:"total,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ActionMsg struct {
	Code    int    `json:"code,omitempty"`
	Total   int    `json:"total,omitempty"`
	Message string `json:"message"`
}

type ActionSum struct {
	Code    int         `json:"code,omitempty"`
	Total   int         `json:"total,omitempty"`
	SumQty  int         `json:"sum_qty"`
	Count   int64       `json:"count"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ActionPaginate struct {
	Code     int         `json:"code,omitempty"`
	Total    int         `json:"total_data,omitempty"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	PageMeta interface{} `json:"page_meta"`
}

// code 200
func AccepData(message string, data interface{}, res *gin.Context) {
	res.JSON(200, Action{
		Code:    200,
		Message: message,
		Data:    data,
	})
	return
}

// code 200
func AccepMsg(message string, res *gin.Context) {
	res.JSON(200, ActionMsg{
		Code:    200,
		Message: message,
	})
	return
}

// code 400
func BadRequest(data interface{}, res *gin.Context) {
	res.JSON(400, Action{
		Code:    400,
		Message: "Bad Request",
		Data:    data,
	})
	return
}

// code 403
func NoAccess(res *gin.Context) {
	res.JSON(403, ActionMsg{
		Code:    403,
		Message: "Akses Dilarang",
	})
	return
}

// code 500
func CustomError(message string, res *gin.Context) {
	res.JSON(500, gin.H{
		"code":    500,
		"message": message,
	})
	return
}

// code 500
func Abort(res *gin.Context) {
	res.JSON(500, ActionMsg{
		Code:    500,
		Message: "Internal Server Error",
	})
	return
}

// code 404
func NotFound(res *gin.Context) {
	res.JSON(404, ActionMsg{
		Code:    404,
		Message: "Data Not Found",
	})
	return
}

// code 201
func AccepCount(message string, total int, data interface{}, res *gin.Context) {
	res.JSON(201, Action{
		Code:    201,
		Total:   total,
		Message: message,
		Data:    data,
	})
	return
}

func AcceptPaginate(request *http.Request, res *gin.Context, message string, data interface{}, loadedItemsCount, page, page_size int, totalItemsCount int64) {
	page_meta := map[string]interface{}{}
	page_meta["offset"] = (page - 1) * page_size
	page_meta["requested_page_size"] = page_size
	page_meta["current_page_number"] = page
	page_meta["current_items_count"] = loadedItemsCount

	page_meta["prev_page_number"] = 1
	total_pages_count := int(math.Ceil(float64(totalItemsCount) / float64(page_size)))
	page_meta["total_pages_count"] = total_pages_count
	if page < total_pages_count {
		page_meta["has_next_page"] = true
		page_meta["next_page_number"] = page + 1
	} else {
		page_meta["has_next_page"] = false
		page_meta["next_page_number"] = 1
	}
	if page > 1 {
		page_meta["prev_page_number"] = page - 1
	} else {
		page_meta["has_prev_page"] = false
		page_meta["prev_page_number"] = 1
	}

	page_meta["next_page_url"] = fmt.Sprintf("%v?page=%d&page_size=%d", request.URL.Path, page_meta["next_page_number"], page_meta["requested_page_size"])
	page_meta["prev_page_url"] = fmt.Sprintf("%s?page=%d&page_size=%d", request.URL.Path, page_meta["prev_page_number"], page_meta["requested_page_size"])

	res.JSON(201, ActionPaginate{
		Code:     201,
		Total:    loadedItemsCount,
		Message:  message,
		Data:     data,
		PageMeta: page_meta,
	})
	return
}
