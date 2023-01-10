package serializer

import "gin_mall_tmp/model"

type Address struct {
	Id       uint   `json:"id"`
	CreateAt int64  `json:"created_at"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func BuildAddress(address *model.Address) Address {
	return Address{
		Id:       address.ID,
		CreateAt: address.CreatedAt.Unix(),
		UserId:   address.UserID,
		Name:     address.Name,
		Phone:    address.Phone,
		Address:  address.Address,
	}
}

func BuildAddresses(items []*model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return
}
