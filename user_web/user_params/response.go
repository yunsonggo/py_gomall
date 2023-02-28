package user_params

import (
	"fmt"
	userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"
	"time"
)

type JsonTime time.Time

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	stmp := fmt.Sprintf("\"%s\"", time.Time(jt).Format("2006-01-02"))
	return []byte(stmp), nil
}

type Users struct {
	Total int64  `json:"total"`
	Data  []User `json:"data"`
}

type User struct {
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	//Birthday JsonTime `json:"birthday"`
	Birthday string `json:"string"`
	Addr     string `json:"addr"`
	Desc     string `json:"desc"`
	Gender   string `json:"gender"`
}

func UserRespToParam(info *userpb.UserResponse) User {
	u := User{
		Mobile:   info.Mobile,
		Nickname: info.Nickname,
		Icon:     info.Icon,
		//Birthday: time.Unix(info.Birthday, 0).Format("2006-01-02"),
		//Birthday: JsonTime(time.Unix(info.Birthday, 0)),
		Addr:   info.Addr,
		Desc:   info.Desc,
		Gender: info.Gender,
	}
	if info.Birthday > 0 {
		u.Birthday = time.Unix(info.Birthday, 0).Format("2006-01-02")
	}
	return u
}
