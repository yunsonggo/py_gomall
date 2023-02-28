package user_params

import userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"

type LoginParam struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,phone" `
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

type SignupParam struct {
	Mobile   string `json:"mobile" binding:"required,phone"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Birthday int64  `json:"birthday"`
	Addr     string `json:"addr"`
	Desc     string `json:"desc"`
	Gender   string `json:"gender"`
	Role     int32  `json:"role"`
}

func SignupParamToReq(sp SignupParam) *userpb.UserRequest {
	ur := userpb.UserRequest{
		Mobile:   sp.Mobile,
		Password: sp.Password,
		Nickname: sp.Nickname,
		Icon:     sp.Icon,
		Birthday: sp.Birthday,
		Addr:     sp.Addr,
		Desc:     sp.Desc,
		Gender:   sp.Gender,
		Role:     sp.Role,
	}
	return &ur
}
