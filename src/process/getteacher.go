package process

import (
	"fmt"
	"net/http"
	. "wbproject/httpserver/src/envbuild"

	"github.com/gin-gonic/gin"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		Logger.Critical(err)
	}
}

func DealError(err error, c *gin.Context) {

	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"error": err.Error(),
	})
}

func GetTeacher(cfg *Config, c *gin.Context) {

	//查看班级号
	cid := -1
	err := cfg.Db.QueryRow(`select class from student where score = (select max(score) from student) order by sid asc limit 1
`).Scan(&cid)
	//数据库查询出错且出错标志不为空
	if err != nil {
		DealError(err, c)
	} else {
		//通过class，找teacher名字
		teacher := ""
		result := ""
		rows, err := cfg.Db.Query("select teacher from class where class = ? ", cid)
		defer rows.Close()
		if err != nil {
			DealError(err, c)
		}
		for rows.Next() {
			err := rows.Scan(&teacher)
			result = result + teacher + ","
			if err != nil {
				DealError(err, c)
			}
		}
		err = rows.Err()
		if err != nil {
			DealError(err, c)
		}
		result = result[0 : len(result)-1]

		c.JSON(http.StatusOK, gin.H{
			"teacher": result,
		})
	}
	return
}
