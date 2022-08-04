package duanyan

import "fmt"

func Test2() {
	sop := map[string]interface{}{
		//"companyId":    "W00000000176",
		//"action":       "sopCategory",
		//"customerType": 1,
		//"customerId":   "f39be893-ec9a-11ec-8301-62258f8a6a3c",
		//"owner":        "13",
		//"Operator":     "13",
		//"shareUIDs":    []string{},
		//"addQwTag":     nil,
		//"moveQwTag":    nil,
		//"addTag":       nil,
		//"moveTag":      nil,
		"status": "uuid0",
	}
	inter(sop)

}

func inter(sop interface{}) {
	s, ok := sop.(SopEventPayload)
	if !ok {
		fmt.Printf("%+v", ok)
	}
	fmt.Printf("%+v", s)
}
