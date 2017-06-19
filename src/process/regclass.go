package process

import (
	"fmt"
	"net/http"
	. "wbproject/httpserver/src/envbuild"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

//检查teacher合法性
func CheckTeacherValid(name string, class int) error {

	if name == "" {
		return fmt.Errorf("teacher 不能为空")
	}

	s := []byte(name)
	if len(s) > 20 {
		return fmt.Errorf("teacher 长度超限")
	}

	if class < 0 || class > 99 {
		return fmt.Errorf("class 数值错误")
	}

	return nil
}

func RegClass(cfg *Config, c *gin.Context) {

	//解析命令
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	js, err := simplejson.NewJson(buf[0:n])
	Logger.Debug("create request", string(buf[0:n]))

	if err != nil {
		err := fmt.Errorf("err:%v", err)
		DealError(err, c)
		return
	}
	cid := js.Get("classNumber").MustInt()
	name := js.Get("teacher").MustString()
	err = CheckTeacherValid(name, cid)
	if err != nil {
		DealError(err, c)
	} else {
		//正常流程，用insert into on dulicate 覆盖
		sql := fmt.Sprintf(`insert into class 
		(class,teacher) values (%d, '%s') 
		on duplicate key update teacher = values(teacher)`,
			cid, name)

		_, err = cfg.Db.Exec(sql)
		if err != nil {
			err = fmt.Errorf("insert sql err: %s", err.Error())
			DealError(err, c)
		} else {
			info := fmt.Sprintf("Successful register class:[%d],teacher:[%s] into class table", cid, name)
			Logger.Infof(info)
			c.JSON(http.StatusOK, gin.H{
				"success": info,
			})
		}
	}
	return
}
