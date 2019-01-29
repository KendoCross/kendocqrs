package application

import (
	"context"
	"encoding/json"

	"github.com/KendoCross/kendocqrs/domain/bus"

	eh "github.com/looplab/eventhorizon"
)

//api2Cmd  路由到领域CMD的映射
var api2Cmd map[string]eh.CommandType

type Result struct {
	Succ bool        `json:"success"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

func HandCommand(postBody []byte, commandKey string) (result Result) {
	cmd, err := eh.CreateCommand(eh.CommandType(commandKey))
	if err != nil {
		result.Msg = "could not create command: " + err.Error()
		return
	}
	if err := json.Unmarshal(postBody, &cmd); err != nil {
		result.Msg = "could not decode Json" + err.Error()
		return
	}
	ctx := context.Background()
	if err := bus.HandleCommand(ctx, cmd); err != nil {
		result.Msg = "could not handle command: " + err.Error()
		return
	}

	result.Succ = true
	result.Msg = "ok"

	return
}
