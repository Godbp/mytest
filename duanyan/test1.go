package duanyan

// SopEvent sop事件
type SopEvent struct {
	DoMain          string      `json:"domain"`          // 域 crm、mall....
	SubDoMain       string      `json:"domainChild"`     // 子域 crm - customer、mall - order....
	CompanyID       string      `json:"companyId"`       // 尘峰公司唯一标识
	CorpID          string      `json:"corpId"`          // 企微公司唯一标识
	EventHappenTime string      `json:"eventHappenTime"` // 事件产生时间
	Payload         interface{} `json:"payload"`         // 业务数据
	TraceID         string      `json:"traceId"`         // 追踪ID
}

// SopEventPayload  清洗后的 sop事件
type SopEventPayload struct {
	//CompanyID    string   `json:"companyId"`    // 租户ID
	//Action       string   `json:"action"`       // 事件类型
	//CustomerType int      `json:"customerType"` // 联系人类型 1 普通客户 2 外部联系人
	//CustomerID   string   `json:"customerId"`   // 客户ID
	//Owner        string   `json:"owner"`        // 客户跟进人
	//Operator     string   `json:"Operator"`     // 当前操作者
	//ShareUIDs    []string `json:"shareUIDs"`    // 共享人ID
	//AddQwTag     []string `json:"addQwTag"`     // 添加企微标签
	//MoveQwTag    []string `json:"moveQwTag"`    // 移除企微标签
	//AddTag       []string `json:"addTag"`       // 添加行为标签
	//MoveTag      []string `json:"moveTag"`      // 移除行为标签
	Status string `json:"status"` // 客户跟进状态
}
