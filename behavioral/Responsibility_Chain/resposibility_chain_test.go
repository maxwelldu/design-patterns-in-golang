package Responsibility_Chain

import (
	"testing"
)

//客户类（Client）角色：创建处理链，并向链头的具体处理者对象提交请求，它不关心处理细节和请求的传递过程。
func TestHandler_Handle(t *testing.T) {
	medical := &medical{}

	// 收银完去药房
	cashier := &cashier{}
	cashier.setNext(medical)

	// 医生检查完去收银处
	doctor := &doctor{}
	doctor.setNext(cashier)

	// 挂号完去看医生
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "zhangsan"}
	// 病人先挂号
	reception.execute(patient)
}
