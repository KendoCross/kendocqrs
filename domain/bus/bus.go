package bus

import (
	"context"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
)

//commandBus 命令总线
var commandBus = bus.NewCommandHandler()

//eventBus 事件总线
var eventBus = eventbus.NewEventBus(nil)

//
var eventStore eh.EventStore

//aggregateStore 领域事件存储与发布
//var AggregateStore *events.AggregateStore

func InitBus() {
	eventStore, _ = eventstore.NewEventStore("127.0.0.1:27017", "EventStore")
	//AggregateStore, _ = events.NewAggregateStore(eventStore, eventBus)
}

//RegisterHandler 注册命令的处理
func RegisterHandler(cmd eh.CommandType, cmdHandler eh.Aggregate) {
	err := commandBus.SetHandler(cmdHandler, cmd)
	if err != nil {
		panic(err)
	}
}

//HandleCommand 命令的执行
func HandleCommand(ctx context.Context, cmd eh.Command) error {
	return commandBus.HandleCommand(ctx, cmd)
}

//RegisterEventHandler 注册事件的处理
func RegisterEventHandler(evtMatcher eh.EventMatcher, evtHandler eh.EventHandler) {
	eventBus.AddHandler(evtMatcher, evtHandler)
}

//RaiseEvents 异步进行事件的存储 和 发布
func RaiseEvents(ctx context.Context, events []eh.Event, originalVersion int) error {
	go eventStore.Save(ctx, events, originalVersion)
	for _, event := range events {
		err := eventBus.PublishEvent(ctx, event)
		if err != nil {
			return err
		}
	}

	return nil
}
