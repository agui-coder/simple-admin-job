// Package payload defines all the payload structures used in tasks
package payload

type HelloWorldPayload struct {
	Name string `json:"name"`
}

type PayOrderNotifyReq struct {
	MerchantOrderId string `json:"merchantOrderId"`
	PayOrderId      uint64 `json:"payOrderId"`
}

type PayRefundNotifyReq struct {
	MerchantOrderId string `json:"merchantOrderId"`
	PayRefundId     uint64 `json:"payRefundId"`
}
