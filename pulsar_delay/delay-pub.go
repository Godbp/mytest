package sopdelay

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"git.dustess.com/mk-base/gin-ext/config"
	"git.dustess.com/mk-biz/mk-plan-center/application/event/model"
	"git.dustess.com/shared-golib/mesher/log/logrusx"
	"git.dustess.com/shared-golib/pulsar-driver/producer"
	xerrors "github.com/pkg/errors"
)

// delayProducer 延迟等待生产者
var delayProducer *producer.Producer

var delayOnce sync.Once

// newDelayProducer 获取生产者
func newDelayProducer() *producer.Producer {
	if delayProducer == nil {
		delayOnce.Do(func() {
			initDelayProducer()
		})
	}
	return delayProducer
}

// initDelayProducer 初始化
func initDelayProducer() {
	conf := config.NewMKMongoConfig().Pulsar
	c := producer.Config{
		Addrs: conf.Addrs,
		Topic: conf.Topic.SOPNodeDelay,
		Name:  conf.SubscriptionName.SOPNodeDelay,
		Token: conf.Token,
	}
	pulsar, err := producer.NewProducer(c, producer.SetBatchingMaxSize(model.PulsarMegSize))
	if err != nil {
		logrusx.NoTrace.Errorf(err, "init delayProducer failed，producer config[%+v]", c)
		panic(err)
	}
	delayProducer = pulsar
	logrusx.NoTrace.Infof("init delayProducer success, topic[%s], name[%s]",
		conf.Topic.SOPNodeDelay,
		conf.SubscriptionName.SOPNodeDelay,
	)
}

// INodeDelayer 延迟异步服务
type INodeDelayer interface {
	ProducerOne(data model.NodeDelay) error
}

// NewNodeDelayer 创建延迟异步服务
func NewNodeDelayer(ctx context.Context) INodeDelayer {
	return &nodeDelayer{ctx: ctx}
}

type nodeDelayer struct {
	ctx context.Context
}

// ProducerOne 生产一条消息
func (u *nodeDelayer) ProducerOne(data model.NodeDelay) error {
	funcName := "推送延迟队列"
	logger := logrusx.WithContext(u.ctx)
	b, err := json.Marshal(data)
	if err != nil {
		return xerrors.Wrapf(err, "%s marshal failed data[%+v],", funcName, data)
	}
	key := fmt.Sprintf("%s%s", data.CompanyID, data.CustomerID) // 保证分区
	msg := producer.NewMessage(nil, b, key)
	// 添加延迟时间
	msg.DeliverAfter = data.DelayTime
	p := newDelayProducer()
	messageID, err := p.SendOneWithMessage(producer.WithAccountID(data.CompanyID), msg)
	if err != nil {
		return xerrors.Wrapf(err, "%s 推送失败 msg[%+v]", funcName, msg)
	}

	logger.Infof("%s 推送成功 messageID[%+v] data[%+v]", funcName, messageID.EntryID(), data)
	return nil
}
