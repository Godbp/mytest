package sopdelay

import (
	"context"
	"encoding/json"
	context2 "git.dustess.com/mk-base/context"
	"git.dustess.com/mk-biz/mk-plan-center/application/event/model"
	"git.dustess.com/shared-golib/mesher/log/logrusx"
	"git.dustess.com/shared-golib/pulsar-driver/constant"
	"git.dustess.com/shared-golib/pulsar-driver/consumer"
	"github.com/apache/pulsar-client-go/pulsar"
	xerrors "github.com/pkg/errors"
	"sync"
)

// 单例创建通知的消费者
var nodeDelayOnce sync.Once

// InitNodeDelayConsumer 延迟消费
func InitNodeDelayConsumer() {
	defer func() {
		err := recover()
		if err != nil {
			logrusx.NoTrace.Infof("node.InitNodeDelayConsumer fetch is panic,err=[%+v]", err)
		}
	}()

	nodeDelayOnce.Do(func() {
		var sub = pulsar.SubscriptionPositionLatest
		var keyShare = pulsar.KeyShared
		c := consumer.Config{
			Addrs:                       []string{"http://pulsar-vgp2awkmr4ne.tdmq-pulsar.ap-sh.public.tencenttdmq.com:8080"},
			Topic:                       "pulsar-vgp2awkmr4ne/mk/sop_node_delay_test3",
			Name:                        "pb",
			Token:                       "eyJrZXlJZCI6InB1bHNhci12Z3AyYXdrbXI0bmUiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItdmdwMmF3a21yNG5lX2FsaS10ZXN0In0.yWu-zTsJtgx_bSBzYrOcYgX-FmCQeKyuiYo_HrcmwF8",
			SubscriptionName:            "pb_g",
			SubscriptionType:            &keyShare,
			RetryEnable:                 false,
			SubscriptionInitialPosition: &sub,
		}
		con := consumer.NewConsumer(c, &nodeDelayHandler{})

		go func() {
			ctx := context.Background()
			err := con.Consume(ctx)
			if err != nil {
				panic(err)
			}
		}()
	})
}

type nodeDelayHandler struct{}

func (h *nodeDelayHandler) Close() {}

// ConsumeClaim 消费通知任务
func (h *nodeDelayHandler) ConsumeClaim(messages chan *constant.GroupSession) error {
	funcName := "nodeDelayHandler.ConsumeClaim"
	ctx := context.TODO()
	logger := logrusx.WithContext(ctx)

	for message := range messages {
		msg := message.GetMessage()
		if msg == nil {
			logger.Infof("%s msg is nil", funcName)
			message.Ack()
			continue
		}
		payload := msg.Payload()
		if payload == nil {
			logger.Infof("%s payload is nil", funcName)
			message.Ack()
			continue
		}

		if err := h.handle(ctx, payload); err != nil {
			logger.Errorf(err, "%s handle error payload=%+s", funcName, string(payload))
			// // 判断重试次数
			// if message.GetMessage().RedeliveryCount() <= model.PulsarRetryCount {
			// 	message.ReconsumeLater(message.GetMessage(), model.PulsarRetryDelayTime)
			// } else { // 超过重试次数
			// 	logger.Errorf(err, "%s handle error and over retry payload=%+s", funcName, string(payload))
			// 	message.Ack()
			// }
			// continue
		}
		message.Ack()
	}
	return nil
}

// handle 数据处理
func (h *nodeDelayHandler) handle(ctx context.Context, message []byte) error {
	funcName := "消费延迟队列"

	// 参数解析
	data := &model.NodeDelay{}
	err := json.Unmarshal(message, data)
	if err != nil {
		return xerrors.Wrapf(err, "%s unmarshal error", funcName)
	}

	// todo 彭博 处理metrics埋点
	//startTime := time.Now()
	fields := logrusx.NewFields().Set("mq", "delayMQ").SetWithMap(data.GetTraceMap())
	ctx = logrusx.ContextWithFields(context2.NewDetachContextWithCID(data.CompanyID, ""), fields)
	logger := logrusx.WithContext(ctx)

	//logger.Infof("%s received data[%+v] nowTime[%+v]", funcName, data, startTime)
	//defer func() {
	//	logger.Infof("%s processed timecost[%v]", funcName, time.Since(startTime))
	//}()
	logger.Infof("消费延迟队列成功 ---------")
	//err = engine.NewDelayer(ctx).SubDelayMQ(data)
	//if err != nil {
	//	logger.Errorf(err, "%s error[%+v] data[%+v]", funcName, err, data)
	//}
	return nil
}
