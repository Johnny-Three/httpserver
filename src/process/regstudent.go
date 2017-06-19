package process

import (
	"fmt"
	"net/http"
	. "wbproject/httpserver/src/envbuild"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

func CheckStuValid(sid string, class, score int) error {

	s := []byte(sid)
	if len(s) != 5 {
		return fmt.Errorf("student id 长度错误")
	}
	if class < 0 || class > 99 {
		return fmt.Errorf("class 数值错误")
	}
	if score < 0 || score > 100 {
		return fmt.Errorf("score 数值错误")
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			return fmt.Errorf("student id 格式错误")
		}
	}
	return nil
}

func RegStudent(cfg *Config, c *gin.Context) {

	//解析命令
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	js, err := simplejson.NewJson(buf[0:n])
	Logger.Debug("RegStudent request", string(buf[0:n]))

	if err != nil {
		err := fmt.Errorf("err:%v", err)
		DealError(err, c)
		return
	}
	count := 0
	sid := js.Get("id").MustString()
	class := js.Get("classNumber").MustInt()
	score := js.Get("score").MustInt()
	err = CheckStuValid(sid, class, score)
	if err != nil {
		DealError(err, c)
		return
	}
	//处理班级编号不存在的情况
	err = cfg.Db.QueryRow("select count(*) from class where class = ?", class).Scan(&count)
	if err != nil {
		DealError(err, c)
		return
	}
	if count == 0 {
		err = fmt.Errorf("班级编号不存在")
		DealError(err, c)
	} else {
		//正常流程，用insert into on dulicate 覆盖
		sql := fmt.Sprintf(`insert into student 
		(sid, class,score) values ('%s', %d, %d) 
		on duplicate key update class = values(class),score = values(score)`,
			sid, class, score)

		_, err = cfg.Db.Exec(sql)
		if err != nil {
			err = fmt.Errorf("insert sql err: %s", err.Error())
			DealError(err, c)
		} else {
			info := fmt.Sprintf("Successful register id:[%s],class:[%d],score:[%d] into student table", sid, class, score)
			Logger.Infof(info)
			c.JSON(http.StatusOK, gin.H{
				"success": info,
			})
		}
	}
	return
}
