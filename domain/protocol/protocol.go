package protocol

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KendoCross/kendocqrs/domain/bus"
	"github.com/KendoCross/kendocqrs/infrastructure"
	"github.com/KendoCross/kendocqrs/interfaces"

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
		//命令只需要确定输入参数满足业务校验即可
		err := a.CheckSign()
		if err != nil {
			return err
		}
		//实际的业务可以异步进行处理
		go a.CfrmSign(cmd)

		return nil
	}
	return fmt.Errorf("couldn't handle command")
}

func (a *ProtocolAggregate) CheckSign() error {
	//校验输入参数是否有效，
	return nil
}

func (a *ProtocolAggregate) CfrmSign(cmd *SignProtocol) error {

	protocolOrm := &interfaces.ProtocolOrm{
		ProtocolNo:   uuid.New().String(),
		AisleType:    cmd.AisleType,
		BankCode:     cmd.BankCode,
		BankCardNo:   cmd.BankCardNo,
		ReservePhone: cmd.ReservePhone,
		CardName:     cmd.CardName,
		IDCardNo:     cmd.IDCardNo,
		Merrem:       cmd.Merrem,
		Remark:       cmd.Remark,
		Status:       1,
	}
	protocolOrm.AccountType, _ = strconv.Atoi(cmd.AccountType)
	protocolOrm.AccountProp, _ = strconv.Atoi(cmd.AccountProp)
	protocolOrm.IDCardType, _ = strconv.Atoi(cmd.IDCardType)

	//这里本应该的业务逻辑请求，通过领域服务
	//result := domainser.Allinpay.SignGetCode(protocolOrm)

	protocolRepo := infrastructure.RepoFac.ProtocolRepo
	_, err := protocolRepo.Add(protocolOrm)

	if err != nil {
		return err
	}
	ctx := context.Background()
	//业务处理成功后，激活领域事件
	bus.RaiseEvents(ctx, a.Events(), 0)

	return nil
}
