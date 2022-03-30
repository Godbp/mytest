package test1

import "github.com/zhuxiujia/GoMybatis"

func init() {
	engine := GoMybatis.GoMybatisEngine{}.New()
	db, err := engine.Open("mysql", "root:guo6RjqO4OBlinv2VXAt@tcp(1.117.244.157:4000)/gaoji")
	engine.WriteMapperPtr()
}
