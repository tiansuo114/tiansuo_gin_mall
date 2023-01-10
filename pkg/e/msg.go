package e

var MsgFlags = map[int]string{
	Success:            "ok",
	Error:              "fail",
	InvalidParams:      "参数错误",
	ErrorJwtMiddleWare: "中间件认证错误",

	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthToken:             "token验证失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorUploadFail:            "文件上传失败",
	ErrorSendEmail:             "邮件发送失败",
	ErrorUpdateUser:            "更新用户信息失败",
	ErrorGetUser:               "获取用户信息失败",

	ErrorProductImgUpload: "图片上传错误",
	ErrorProductUpDate:    "商品更新失败",
	ErrorProductDelete:    "商品删除失败",
	ErrorProductExist:     "商品不存在",
	ErrorProductGet:       "获取商品失败",
	ErrorProductCreate:    "商品创建失败",

	ErrorFavoriteExist:  "收藏夹已存在",
	ErrorFavoriteCreate: "创建收藏失败",
	ErrorFavoriteDelete: "删除收藏失败",

	ErrorAddressCreate: "地址创建失败",
	ErrorAddressGet:    "获取地址失败",
	ErrorAddressList:   "获取地址清单失败",
	ErrorAddressUpdate: "更新地址失败",
	ErrorAddressDelete: "删除地址失败",

	ErrorCartExist:  "购物车中已存在",
	ErrorCartCreate: "购物车创建失败",
	ErrorCartDelete: "购物车删除失败",
	ErrorCartGet:    "购物车获取失败",
	ErrorCartUpdate: "购物车更新失败",

	ErrorOrderExist:  "订单已存在",
	ErrorOrderCreate: "创建订单失败",
	ErrorOrderDelete: "删除订单失败",
	ErrorOrderList:   "获取订单列表失败",
	ErrorOrderGet:    "获取订单失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
