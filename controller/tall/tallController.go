package tall

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanghongc/blockChainStudy/modules"
)

type TallControllerInit struct {
}

func (tall TallControllerInit) AddTall(c *gin.Context) {
	var com modules.Comment
	if err := c.ShouldBind(&com); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := modules.DB.Create(&com) // 通过数据的指针来创建
	if result.RowsAffected > 1 {
		fmt.Print(com)
	}
	c.String(http.StatusOK, "add 成功")

}

func (tall TallControllerInit) SelTall(c *gin.Context) {
	var com = []modules.Comment{}

	modules.DB.Where("post_id = ?", c.PostForm("post_id")).Find(&com)

	c.JSON(http.StatusOK, gin.H{
		"result": com,
	})

}
