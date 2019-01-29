package protocol

import (
	"context"
	"fmt"

	"github.com/KendoCross/kendocqrs/domain/bus"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func init() {
	prdctAgg := &ProtocolAggregate{
		AggregateBase: events.NewAggregateBase("ProtocolAggregate", uuid.New()),
	}
	bus.RegisterHandler(SignCommand, prdctAgg)
}

type ProtocolAggregate struct {
	*events.AggregateBase
}

var _ = eh.Aggregate(&ProtocolAggregate{})

func (a *ProtocolAggregate) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd := cmd.(type) {
	case *SignProtocol:
		err := fmt.Errorf(cmd.CardName)
		if err != nil {
			return err
		}

		bus.RaiseEvents(ctx, a.Events(), 0)
		return nil
	}
	return fmt.Errorf("couldn't handle command")
}

func (p *ProtocolAggregate) Sign() (string, error) {

	// smsContent := &domainser.QuickPaySignSmsContent{
	// 	BankCode:     p.SignVm.BankCode,
	// 	ACCOUNT_TYPE: p.SignVm.AccountType,
	// 	ACCOUNT_NO:   strings.Replace(p.SignVm.BankCardNo, " ", "", -1),
	// 	ACCOUNT_NAME: p.SignVm.CardName,
	// 	ACCOUNT_PROP: p.SignVm.AccountProp,
	// 	ID_TYPE:      p.SignVm.IDCardType,
	// 	ID:           p.SignVm.IDCardNo,
	// 	TEL:          p.SignVm.ReservePhone,
	// 	MERREM:       p.SignVm.Merrem,
	// 	REMARK:       p.SignVm.Remark,
	// 	AisleType:    p.SignVm.AisleType,
	// }
	// result := domainser.Allinpay.SignGetCode(smsContent)
	// jsonByte, err := json.Marshal(result)
	// if err != nil {
	// 	return "", err
	// }

	// return string(jsonByte), nil

	return "", nil
}

func (p *ProtocolAggregate) CfrmSign(reqSn string, verCode string) (string, error) {

	// result := domainser.ApiResp{}

	// if !infra.MemoryCache.IsExist(reqSn) {
	// 	result.ErrMsg = "无此请求号的记录！"
	// 	jsonByte, err := json.Marshal(result)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return string(jsonByte), nil
	// }

	// smsContent := infra.MemoryCache.Get(reqSn).(domainser.QuickPaySignSmsContent)
	// smsCfrm := &domainser.SmsCfrmContent{
	// 	VERCODE:  verCode,
	// 	SRCREQSN: reqSn,
	// }

	// result = domainser.Allinpay.SignCfrm(smsCfrm)
	// if result.IsSuccess {

	// 	accountType, _ := strconv.Atoi(smsContent.ACCOUNT_TYPE)
	// 	accountProp, _ := strconv.Atoi(smsContent.ACCOUNT_PROP)
	// 	idTp, _ := strconv.Atoi(smsContent.ID_TYPE)

	// 	protocolOrm := &ddd_interfaces.ProtocolOrm{
	// 		ProtocolNo:    uuid.New().String(),
	// 		AisleType:     smsContent.AisleType,
	// 		AisleProtocol: result.RetInfo,
	// 		BankCode:      smsContent.BankCode,
	// 		AccountType:   accountType,
	// 		AccountProp:   accountProp,
	// 		BankCardNo:    smsContent.ACCOUNT_NO,
	// 		ReservePhone:  smsContent.TEL,
	// 		IDCardType:    idTp,
	// 		CardName:      smsContent.ACCOUNT_NAME,
	// 		IDCardNo:      smsContent.ID,
	// 		Merrem:        smsContent.MERREM,
	// 		Remark:        smsContent.REMARK,
	// 		Status:        1,
	// 	}

	// 	protocolRepo := crossutting.RepoFac.GetProtocolRepo()
	// 	_, err := protocolRepo.Add(protocolOrm)

	// 	if err != nil {
	// 		result.ErrMsg = "保存数据库失败，严重异常！"
	// 	}

	// 	result.RetInfo = protocolOrm.ProtocolNo
	// }

	// jsonByte, err := json.Marshal(result)
	// if err != nil {
	// 	return "", err
	// }

	// return string(jsonByte), nil

	// eventEnvelop, err := dddcore.NewEvent(p.id, fmt.Sprintf("%T", p), p.version, protocolOrm)
	// if err != nil {
	// 	return "-1", err
	// }
	// p.changes = append(p.changes, eventEnvelop)
	return "", nil
}
