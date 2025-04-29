package services

import (
    "strconv"

    "github.com/gin-gonic/gin"
)

func PaginationHandler(c *gin.Context) (int, int, int) {
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "5")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit < 1 {
        limit = 5
    }
    offset := (page - 1) * limit

    return offset, limit, page
}