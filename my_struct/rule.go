package my_struct

// 文本的function枚举
const (
	StringFunctionContain    = "contain"    // 包含
	StringFunctionNotContain = "notContain" // 不包含
	StringFunctionEmpty      = "empty"      // 为空
	StringFunctionNotEmpty   = "notEmpty"   // 非空
)

// 日期的function枚举
const (
	DateFunctionBetween  = "between"  // 范围
	DateFunctionGt       = "gt"       // 大于
	DateFunctionLt       = "lt"       // 大于
	DateFunctionEqual    = "equal"    // 等于
	DateFunctionRecently = "recently" // 最近多少天
	DateFunctionToday    = "today"    // 等于今天
)

// 月日的function枚举
const (
	MonthDateFunctionBetween  = "between"  // 范围
	MonthDateFunctionGt       = "gt"       // 大于
	MonthDateFunctionLt       = "lt"       // 大于
	MonthDateFunctionEqual    = "equal"    // 等于
	MonthDateFunctionRecently = "recently" // 最近多少天
	MonthDateFunctionToday    = "today"    // 等于今天
)

// 数字类
const (
	NumberFunctionGt       = "gt"       // 大于
	NumberFunctionLt       = "lt"       // 小于
	NumberFunctionEqual    = "equal"    // 等于
	NumberFunctionNotEqual = "notEqual" // 不等于
	NumberFunctionGte      = "gte"      // 大于等于
	NumberFunctionLte      = "lte"      // 小于等于
	NumberFunctionBetween  = "between"  // 范围
)

// 数组类
const (
	SliceFunctionContainAll    = "containAll"    // 包含所有
	SliceFunctionContainAny    = "containAny"    // 包含任一
	SliceFunctionNotContainAll = "notContainAll" // 不包含所有
	SliceFunctionNotContainAny = "notContainAny" // 不包含任一
)

var SliceOperationMap = map[string]string{
	SliceFunctionContainAll:    "",
	SliceFunctionContainAny:    "",
	SliceFunctionNotContainAll: "",
	SliceFunctionNotContainAny: "",
}

var NumberOperationMap = map[string]string{}
