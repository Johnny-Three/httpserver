package process

import (
	"fmt"
	"net/http"
	. "wbproject/httpserver/src/envbuild"

	"github.com/gin-gonic/gin"
)

//参数正确性检查
func CheckParaValid(sid string) error {

	s := []byte(sid)
	if len(s) != 5 {
		return fmt.Errorf("student id 长度错误")
	}

	for _, v := range s {
		if v < '0' || v > '9' {
			return fmt.Errorf("student id 格式错误")
		}
	}
	return nil
}

func GetScore(cfg *Config, c *gin.Context) {

	sid := c.Param("sid")
	cid, totalscores := -1, 0
	err := CheckParaValid(sid)
	if err != nil {
		DealError(err, c)
		return
	}
	//处理班级编号不存在的情况
	err = cfg.Db.QueryRow("select class from student where sid = ?", sid).Scan(&cid)
	//数据库查询出错且出错标志不为空
	if err != nil && err.Error() != "sql: no rows in result set" {
		DealError(err, c)

	} else if err != nil && err.Error() == "sql: no rows in result set" {
		//没有找到学号为sid的学生
		err = fmt.Errorf("student-not-found")
		DealError(err, c)

	} else {
		//找到class，利用mysql做sum
		err = cfg.Db.QueryRow("select sum(score) from student where class = ?", cid).Scan(&totalscores)
		if err != nil {
			DealError(err, c)
		}
		c.JSON(http.StatusOK, gin.H{
			"total": totalscores,
		})
	}
	return
}
