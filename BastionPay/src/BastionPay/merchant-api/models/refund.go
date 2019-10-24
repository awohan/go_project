package models

import (
	"BastionPay/merchant-api/api"
	"BastionPay/merchant-api/common"
	"BastionPay/merchant-api/db"
)

type (
	ReFund struct {
		MerchantId              *string `json:"merchant_id,omitempty"                   gorm:"column:merchant_id;type:varchar(10)"`
		MerchantRefundNo        *string `json:"merchant_refund_no"                      gorm:"column:merchant_refund_no;type:varchar(64)"`
		NotifyUrl               *string `json:"notify_url,omitempty"                    gorm:"column:notify_url;type:varchar(64)"`
		OriginalMerchantTradeNo *string `json:"original_merchant_trade_no,omitempty"    gorm:"column:original_merchant_trade_no;type:varchar(64)"`
		Remark                  *string `json:"remark,omitempty"                        gorm:"column:remark;type:varchar(120)"`
		Table
	}

	ReFundWithInfo struct {
		MerchantId              *string `json:"merchant_id,omitempty"                   gorm:"column:merchant_id;type:varchar(10)"`
		MerchantRefundNo        *string `json:"merchant_refund_no"                      gorm:"column:merchant_refund_no;type:varchar(64)"`
		NotifyUrl               *string `json:"notify_url,omitempty"                    gorm:"column:notify_url;type:varchar(64)"`
		OriginalMerchantTradeNo *string `json:"original_merchant_trade_no,omitempty"    gorm:"column:original_merchant_trade_no;type:varchar(64)"`
		Remark                  *string `json:"remark,omitempty"                        gorm:"column:remark;type:varchar(120)"`
		PayeeId                 *string `json:"payee_id,omitempty"                      gorm:"column:payee_id;type:varchar(64)"`
		Assets                  *string `json:"assets,omitempty"                        gorm:"column:assets;type:varchar(16)"`
		Amount                  *string `json:"amount,omitempty"                        gorm:"column:amount;type:varchar(20)"`
		Table
	}
)

//状态， 0初始状态， 1成功， 2失败， 3关闭

func (this *ReFund) TableName() string {
	return "refund"
}

func (this *ReFund) Parse(p *api.RefundTrade) *ReFund {
	return &ReFund{
		MerchantId:              p.MerchantId,
		MerchantRefundNo:        p.MerchantRefundNo,
		NotifyUrl:               p.NotifyUrl,
		OriginalMerchantTradeNo: p.OriginalMerchantTradeNo,
		Remark:                  p.Remark,
	}
}

func (this *ReFundWithInfo) Parse(p *ReFund) *ReFundWithInfo {
	return &ReFundWithInfo{
		MerchantId:              p.MerchantId,
		MerchantRefundNo:        p.MerchantRefundNo,
		NotifyUrl:               p.NotifyUrl,
		OriginalMerchantTradeNo: p.OriginalMerchantTradeNo,
		Remark:                  p.Remark,
	}
}

func (this *ReFund) ParseList(p *api.RefundTradeList) *ReFund {
	return &ReFund{
		MerchantId: p.MerchantId,
	}
}

func (this *ReFund) Add() error {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *ReFund) List(merchantId string, page, size int64) ([]*ReFund, []string, error) {
	var list []*ReFund
	condition := &ReFund{MerchantId: &merchantId}

	query := db.GDbMgr.Get().Where(condition)

	//return new(common.Result).PageQuery(query, &Trade{}, &list, page, size, nil, "")
	_, err := new(common.Result).PageQuery(query, &ReFund{}, &list, page, size, nil, "")
	if err != nil {
		return nil, nil, err
	}

	originOrders := make([]string, 0)
	for i := 0; i < len(list); i++ {
		originOrders = append(originOrders, *list[i].OriginalMerchantTradeNo)
	}

	return list, originOrders, nil
}

//func (this *ReFund) UpdateStatusByTradeNo(merchant_refund_no string, status int, trade_no string) ( error) {
//
//	if err := db.GDbMgr.Get().Model(&ReFund{}).Where("merchant_refund_no = ?", merchant_refund_no).Updates(ReFund{ReFundStatus: &status, TradeNo: &trade_no}).Error; err != nil {
//		return err
//	}
//	//mod := new(table.Trade)
//	//err := db.GDbMgr.Get().Where("merchant_trade_no = ?", MerchantTradeNo).Last(mod).Error
//	//if err != nil {
//	//	return nil, err
//	//}
//	return nil
//}

//部分退款用
//func (this *ReFund) GetFrozenAmount(tx *gorm.DB, OriginalMerchantTradeNo string) ( error) {
//
//	if err := db.GDbMgr.Get().Model(&ReFund{}).Select("sum()").Where("merchant_refund_no = ?", merchant_refund_no).Updates(ReFund{ReFundStatus: &status, TradeNo: &trade_no}).Error; err != nil {
//		return err
//	}
//	//mod := new(table.Trade)
//	//err := db.GDbMgr.Get().Where("merchant_trade_no = ?", MerchantTradeNo).Last(mod).Error
//	//if err != nil {
//	//	return nil, err
//	//}
//	return nil
//}