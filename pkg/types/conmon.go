package types

import "bnb-48-ins-indexer/dao"

type CommonListCond struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int8  `json:"page_size" binding:"required" form:"page_size"`
}

type CommonListRsp struct {
	Count    int64  `json:"count"`
	Page     uint64 `json:"page"`
	PageSize uint8  `json:"page_size"`
}

type ListAccountWalletReq struct {
	CommonListCond
	TickHash string `json:"tick_hash" binding:"required" form:"tick_hash"`
}

type ListAccountWalletRsp struct {
	CommonListRsp
	List []*dao.AccountWalletModel `json:"list"`
}

type AccountBalanceReq struct {
	TickHash []string `json:"tick_hash" form:"tick_hash"`
	Address  string   `json:"address" binding:"required" form:"address"`
}

type AccountBalanceRsp struct {
	Wallet []*dao.AccountWalletModel `json:"wallet"`
}

type ListRecordRsp struct {
	CommonListRsp
	List []*dao.AccountRecordsModel `json:"list"`
}

type ListInscriptionWalletReq struct {
	CommonListCond
	Protocol string `json:"protocol" form:"protocol"`
	TickHash string `json:"tick_hash" form:"tick_hash"`
	Status   uint64 `json:"status" form:"status"`
	Tick     string `json:"tick" form:"tick"`
	DeployBy string `json:"deploy_by" form:"deploy_by"`
}

type ListRecordReq struct {
	CommonListCond
	TickHash string `json:"tick_hash" form:"tick_hash"`
}

type ListInscriptionRsp struct {
	CommonListRsp
	List []*dao.InscriptionModel `json:"list"`
}
