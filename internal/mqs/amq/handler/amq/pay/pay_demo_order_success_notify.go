package pay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/agui-coder/simple-admin-pay-rpc/payclient"
	"github.com/hibiken/asynq"
	"github.com/suyuan32/simple-admin-job/internal/mqs/amq/types/payload"
	"github.com/suyuan32/simple-admin-job/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type DemoOrderSuccessNotifyHandler struct {
	svcCtx *svc.ServiceContext
}

func NewDemoOrderSuccessNotifyHandler(svcCtx *svc.ServiceContext) *DemoOrderSuccessNotifyHandler {
	// 需要使用taskid，可以从数据库中查询任务日志,参照hello_world.go
	//task, err := svcCtx.DB.Task.Query().Where(task.PatternEQ(pattern.PayOrderSuccessNotify)).First(context.Background())
	//if err != nil || task == nil {
	//	logx.Error("failed to load task pattern: %s", pattern.PayOrderSuccessNotify)
	//	return nil
	//}
	return &DemoOrderSuccessNotifyHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 支付结果通知
func (l *DemoOrderSuccessNotifyHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p payload.PayOrderNotifyReq
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Join(err, fmt.Errorf("failed to umarshal the payload :%s", string(t.Payload())))
	}
	id, err := strconv.ParseUint(p.MerchantOrderId, 10, 64)
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to int MerchantOrderId  :%s", p.MerchantOrderId))
	}
	resp, err := l.svcCtx.PayRpc.UpdateDemoOrderPaid(ctx, &payclient.UpdateDemoOrderPaidReq{Id: id, PayOrderId: p.PayOrderId})
	if err != nil {
		return err
	}
	logx.Infof("DemoOrderSuccessNotifyTradeHandler success resp: %s", resp.Msg)
	return nil
}
