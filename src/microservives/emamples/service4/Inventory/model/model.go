package model

type OrderStatus struct {
	OrderID    string `gorm:"column:orderid"`
	OrderName    string `gorm:"column:ordername"`
	OrderPrice int64  `gorm:"column:orderprice"`
	OrderUseName int64  `gorm:"column:orderusename"`
}

type Inventory struct {
	CommodityID    string `gorm:"column:commodityid"`
	CommodityName    string `gorm:"column:commodityname"`
	CommodityNumber int64  `gorm:"column:commoditynumber"`
}