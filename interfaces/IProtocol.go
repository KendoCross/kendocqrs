package interfaces

import (
	"time"
)

type IProtocolRepo interface {
	Add(p *ProtocolOrm) (int64, error)
}

type ProtocolOrm struct {
	ID int `xorm:"autoincr pk ID" `
	//协议号
	ProtocolNo string `xorm:"ProtocolNo"`
	//支付平台类型
	AisleType string `xorm:"AisleType"`
	//同一平台细分的子类
	AisleBranch int `xorm:"AisleBranch"`
	//支付平台生成的协议号
	AisleProtocol string `xorm:"AisleProtocol" json:"-"`
	BankId        string `xorm:"BankId"`
	//银行code，各平台不一样
	BankCode string `xorm:"BankCode"`
	//账户类型
	AccountType int `xorm:"AccountType"`
	//账户属性
	AccountProp int `xorm:"AccountProp"`
	//银行卡号
	BankCardNo string `xorm:"BankCardNo"`
	//预留手机号
	ReservePhone string `xorm:"ReservePhone"`
	//银行卡预留的证件类型
	IDCardType int `xorm:"IDCardType"`
	//银行卡开户姓名
	CardName string `xorm:"CardName"`
	//银行卡预留的证件号码
	IDCardNo string `xorm:"IDCardNo"`
	//银行所在城市代号
	CityNo string `xorm:"CityNo"`
	//银行支行名称
	BranchNm string `xorm:"BranchNm"`
	//协议状态
	Status int `xorm:"Status"`
	//冻结
	Freeze bool `xorm:"Freeze"`
	//商户预留信息
	Merrem string `xorm:"Merrem"`
	//备注
	Remark string `xorm:"Remark"`
	//提示标识
	Notes string `xorm:"Notes"`
	//删除
	DelMark bool `xorm:"DelMark"`
	//新增时间
	AddTime time.Time `xorm:"AddTime"`
	//更新时间
	UpdTime time.Time `xorm:"UpdTime"`
}

func (p ProtocolOrm) TableName() string {
	return "Protocol"
}
