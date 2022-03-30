package my_gengine

type EventType string

// 添加业务领域事件类型划分
const (
	// OrderOverdueTime 订单过期时间
	OrderOverdueTime EventType = "orderOverdueTime"
	// SopRuleInstantly 即刻触发事件
	SopRuleInstantly EventType = "instantly"
	// SopRuleAddCustomer 添加客户事件
	SopRuleAddCustomer EventType = "addCustomer"
	// SopRuleFollowStatus 跟进人状态变更事件
	SopRuleFollowStatus EventType = "followStatus"
	// SopRuleFollowTime 跟进时间事件
	SopRuleFollowTime EventType = "followTime"
	// SopRuleDateType 时间类型
	SopRuleDateType EventType = "dateType"
	// SopRuleFixedTime 固定时间触发事件
	SopRuleFixedTime EventType = "fixedTime"
	// SopRuleBirthday 客户生日事件
	SopRuleBirthday EventType = "birthday"
	// SopRuleCustomerTag 客户标签变更事件
	SopRuleCustomerTag EventType = "tag"
	// SopRuleMkScore 客户营销评分事件
	SopRuleMkScore EventType = "score"
)

// Event 事件总线
type Event struct {
	EventType EventType `json:"eventType"`
	Cid       string    `json:"cid"`     // 公司ID
	Cdb       string    `json:"cdb"`     // 分库ID
	Payload   []byte    `json:"payload"` // 埋点实体、定时调度实体
	EventTime string    `json:"eventTime"`
}

// TouchCustomerPayload 客户埋点实体
type TouchCustomerPayload struct {
	CustomerID   string   `json:"customerId"`
	Owner        string   `json:"owner"`    // 跟进人
	UIDs         []string `json:"uids"`     // 共享人
	NewTags      []string `json:"new_tags"` // 新增的企微标签
	Source       int64    `json:"mkScore"`  // 营销评分
	FollowStatus string   `json:"status"`   // 跟进状态
}

// TimerCustomerPayload 客户定时实体
type TimerCustomerPayload struct {
	CustomerID     string   `json:"customerId"`
	Owner          string   `json:"owner"`               // 跟进人
	UIDs           []string `json:"uids"`                // 共享人
	CustomField    Fields   `json:"customField"`         // 客户自定义日期类型字段
	FollowNextTime int64    `json:"log_time"`            // 最近跟进时间
	Birthday       string   `json:"birthdayMonthAndDay"` // 客户生日
	AddTime        int64    `json:"add_time"`            // 添加客户时间
}

// TimerOrderPayload 订单定时实体
type TimerOrderPayload struct {
	Cid         string   `json:"cid"`
	Cdb         string   `json:"cdb"`
	OrderID     string   `json:"orderId"`
	CustomerID  string   `json:"customerId"`
	Owner       string   `json:"owner"`       // 跟进人
	UIDs        []string `json:"uids"`        // 共享人
	CustomField Fields   `json:"customField"` // 客户自定义日期类型字段
}

// Fields 自定义字段
type Fields struct {
	ID          string `json:"id" bson:"_id"`                    // 字段id
	StringValue string `json:"string_value" bson:"string_value"` // 字段类型为：radio（单选），date（日期），datetime（日期时间），text（单行文本），textarea（多行文本）时使用此字段
	Type        string `json:"type"`                             // 自定义字段类别 (目前仅用于销售机会)
}
