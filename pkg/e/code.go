package e

//标准化区分，user模块的错误均以30开头，product模块以40开头,favorite模块以50开头

const (
	Success            = 200
	Error              = 500
	InvalidParams      = 400
	ErrorJwtMiddleWare = 401

	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUploadFail            = 30007
	ErrorSendEmail             = 30008
	ErrorUpdateUser            = 30009
	ErrorGetUser               = 30010

	ErrorProductImgUpload = 40001
	ErrorProductUpDate    = 40002
	ErrorProductDelete    = 40003
	ErrorProductExist     = 40004
	ErrorProductGet       = 40005
	ErrorProductCreate    = 40006

	ErrorFavoriteExist  = 50001
	ErrorFavoriteCreate = 50002
	ErrorFavoriteDelete = 50003

	ErrorAddressCreate = 60001
	ErrorAddressGet    = 60002
	ErrorAddressList   = 60003
	ErrorAddressUpdate = 60004
	ErrorAddressDelete = 60005

	ErrorCartExist  = 70001
	ErrorCartCreate = 70002
	ErrorCartDelete = 70003
	ErrorCartGet    = 70004
	ErrorCartUpdate = 70005

	ErrorOrderExist  = 80001
	ErrorOrderCreate = 80002
	ErrorOrderDelete = 80003
	ErrorOrderList   = 80004
	ErrorOrderGet    = 80005
)
