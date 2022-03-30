package main

import (
	"fmt"
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
	"github.com/sirupsen/logrus"
	"mytest/myString"
)

// gengine 集群部署 、高并发
// 规则生产封装
//
// 日志链路
// 架构设计，性能
// 业务接入，代码修改（配置化）

// User 定义想要注入的结构体
type User struct {
	Name  string
	Addr  string
	Age   float64
	IsBoy bool
	AAA   string
	Res   bool
}

func (u *User) GetNum(i int64) int64 {
	return i
}

func (u *User) Print(s string) {
	fmt.Println(s)
}

func (u *User) Say() {
	fmt.Println("hello world")
}

//定义规则
const rule1 = `
rule "name test" "i can"  salience 0
begin
		if 7 == User.Age && "Calo" == User.Name {
			User.Male = true
			User.Print("明年我8岁了")
		}
end
`

// Builder 规则过滤器
type Builder struct {
	Rule   string      `json:"rule"`   // 执行规则
	Name   string      `json:"name"`   // 规则名称（唯一）
	Remark string      `json:"remark"` // 备注
	Level  int64       `json:"level"`  // 优先级
	Obj    interface{} `json:"obj"`    // 待执行对象
}

func main() {
	//test.MyEncrypt("18982108252", "vms-t2.tezign.com")
	//test.GetImg("https://itg-tezign-files.tezign.com/t2/75fc4ab0d78fea079ae0cda8ae68493b.jpg?Expires=2270430124&OSSAccessKeyId=LTAIiH7NZflLSZy3&Signature=YPST8kSnpR99RJs7wqD%2FvzUVwgk%3D&response-content-disposition=attachment%3B%20filename%3D%22test%2520pic%252012.4.jpg%22%3B%20filename%2A%3Dutf-8%27%27test%2520pic%252012.4.jpg&response-content-type=application%2Foctet-stream")
	//r := myString.Slice("我爱中国", -1, -4)
	//r := my_number.ChuFa(18982108252, 1000)
	//r := my_url.GetHost("https://itg-tezign-files.tezign.com/5f09838316b108b549e824fe77e23a62.jpg?Expires=2270972276&OSSAccessKeyId=LTAIiH7NZflLSZy3&Signature=zNqS6cRl1QG0knlJAwB5PIAxKBw%3D&response-content-disposition=attachment%3B%20filename%3D%22RT31oo-%25E8%2588%2592%25E8%2580%2590%25E7%2588%25BD%25E8%25BA%25AB%25E9%25A6%2599%25E4%25BD%2593%25E5%2596%25B7%25E9%259B%25BE%2520%25E6%25B8%2585%25E9%259A%2590%25E5%258A%25B2%25E7%2588%25BD%252012X150ML-690208128.jpg%22%3B%20filename%2A%3Dutf-8%27%27RT31oo-%25E8%2588%2592%25E8%2580%2590%25E7%2588%25BD%25E8%25BA%25AB%25E9%25A6%2599%25E4%25BD%2593%25E5%2596%25B7%25E9%259B%25BE%2520%25E6%25B8%2585%25E9%259A%2590%25E5%258A%25B2%25E7%2588%25BD%252012X150ML-690208128.jpg&response-content-type=application%2Foctet-stream")
	//r, filePath, _ := my_url.CreateTempFile(ctx.TODO(), "https://itg-tezign-files.tezign.com/5f09838316b108b549e824fe77e23a62.jpg?Expires=2270972276&OSSAccessKeyId=LTAIiH7NZflLSZy3&Signature=zNqS6cRl1QG0knlJAwB5PIAxKBw%3D&response-content-disposition=attachment%3B%20filename%3D%22RT31oo-%25E8%2588%2592%25E8%2580%2590%25E7%2588%25BD%25E8%25BA%25AB%25E9%25A6%2599%25E4%25BD%2593%25E5%2596%25B7%25E9%259B%25BE%2520%25E6%25B8%2585%25E9%259A%2590%25E5%258A%25B2%25E7%2588%25BD%252012X150ML-690208128.jpg%22%3B%20filename%2A%3Dutf-8%27%27RT31oo-%25E8%2588%2592%25E8%2580%2590%25E7%2588%25BD%25E8%25BA%25AB%25E9%25A6%2599%25E4%25BD%2593%25E5%2596%25B7%25E9%259B%25BE%2520%25E6%25B8%2585%25E9%259A%2590%25E5%258A%25B2%25E7%2588%25BD%252012X150ML-690208128.jpg&response-content-type=application%2Foctet-stream")
	myString.SliceList([]string{"1", "2", "3", "4", "5", "6"}, 7)
	//fmt.Printf("%s filePath", r)

}

func rule() {
	orString := map[string][]string{
		"Name": {"==", "pb"},
	}
	orInterface := map[string][]interface{}{
		"Age":   {">", 1},
		"IsBoy": {"==", true},
		"Addr":  {"==", "18"},
	}
	sliceSting := [][3]string{
		{"AAA", "==", "aaa"},
	}
	user := &User{
		Name:  "pb",
		Addr:  "四川",
		Age:   2.2,
		IsBoy: true,
		AAA:   "aaa",
		Res:   false,
	}
	//struct {
	//	ID string `json:"_id"`
	//	SopID string `json:"sop_id"`
	//	Rule string `json:"rule"`
	//}{}
	//type == "timer" // 自定义日期
	//struct {
	//	ID string `json:"_id"`
	//	PriID string `json:"priId"`
	//	Info struct{} `json:"info"`
	//	Type string `json:"type"` // type
	//	Value string `json:"value"`
	//	Key string `json:"key"`
	//}{}
	b := NewBuilder("sop", "任务名称", 0, user)
	b.OrMapStringBuilder(orString)
	b.OrMapInterfaceBuilder(orInterface)
	b.OrSliceStringBuilder(sliceSting)
	b.Build()
	b.Exe()

	fmt.Printf("user.Res=[%v]", user.Res)
}

// NewBuilder 构造器
func NewBuilder(name, remark string, level int64, obj interface{}) *Builder {
	return &Builder{
		Rule:   "",
		Name:   name,
		Remark: remark,
		Level:  level,
		Obj:    obj,
	}
}

// OrMapStringBuilder 或条件
func (b *Builder) OrMapStringBuilder(arg map[string][]string) {
	if len(b.Rule) > 0 {
		b.Rule = b.Rule + " && "
	}
	rule := ""
	for k, v := range arg {
		if len(v) != 2 {
			continue
		}
		rule += fmt.Sprintf("%s %s \"%s\" || ", "Builder."+k, v[0], v[1])
	}
	b.Rule += "(" + rule[:len(rule)-3] + ")"

}

// OrMapInterfaceBuilder 或条件
func (b *Builder) OrMapInterfaceBuilder(arg map[string][]interface{}) {
	if len(b.Rule) > 0 {
		b.Rule = b.Rule + " && "
	}
	rule := ""
	for k, v := range arg {
		if len(v) != 2 {
			continue
		}
		v1, ok := v[0].(string)
		if !ok {
			continue
		}
		v2, ok := v[1].(string)
		if ok {
			rule += fmt.Sprintf("%s %s \"%s\" || ", "Builder."+k, v1, v2)
			continue
		}
		rule += fmt.Sprintf("%s %s %v || ", "Builder."+k, v1, v[1])
	}
	b.Rule += "(" + rule[:len(rule)-3] + ")"
}

// OrSliceStringBuilder 或条件
func (b *Builder) OrSliceStringBuilder(arg [][3]string) {
	if len(b.Rule) > 0 {
		b.Rule = b.Rule + " && "
	}
	rule := ""
	for _, v := range arg {
		rule += fmt.Sprintf("%s %s \"%s\" || ", "Builder."+v[0], v[1], v[2])
	}
	b.Rule += "(" + rule[:len(rule)-3] + ")"

}

// Build 构建
func (b *Builder) Build() {
	b.Rule = fmt.Sprintf(`rule "%s" "%s"  salience %d begin if %s { Builder.Res = true } end`, b.Name, b.Remark, b.Level, b.Rule)
}

// Exe 执行
func (b *Builder) Exe() {
	dataContext := context.NewDataContext()
	// 注入初始化的结构体
	dataContext.Add("Builder", b.Obj)

	ruleBuilder := builder.NewRuleBuilder(dataContext)
	fmt.Printf("rule is [%s]", b.Rule)
	// 注入规则
	err := ruleBuilder.BuildRuleFromString(b.Rule)

	if err != nil {
		logrus.Errorf("err:%s ", err)
	}

	eng := engine.NewGengine()

	// 执行规则
	err = eng.Execute(ruleBuilder, true)
	if err != nil {
		logrus.Errorf("execute rule error: %v", err)
	}
}