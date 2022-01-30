/*
场景：
	多个对象可以处理一个请求，但具体由哪个对象处理该请求在运行时自动确定。
	可动态指定一组对象处理请求，或添加新的处理者。
	需要在不明确指定请求处理者的情况下，向多个处理者中的一个提交请求。
 */
//案例来源 https://www.cnblogs.com/amunote/p/15335419.html
package Responsibility_Chain

import "fmt"

//抽象处理者（Handler）角色：定义一个处理请求的接口，包含抽象处理方法和一个后继连接。
type department interface {
	execute(*patient)
	setNext(department)
}

//具体处理者（Concrete Handler）角色：实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
type medical struct {
	next department
}
func (e *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		if e.next != nil {
			e.next.execute(p)
		}
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	if e.next != nil {
		e.next.execute(p)
	}
}
func (e *medical) setNext(next department) {
	e.next = next
}

//具体处理者（Concrete Handler）角色：实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
type cashier struct {
	next department
}
func (e *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
	if e.next != nil {
		e.next.execute(p)
	}
}
func (e *cashier) setNext(next department) {
	e.next = next
}

//具体处理者（Concrete Handler）角色：实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
type doctor struct {
	next department
}
func (e *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		if e.next != nil {
			e.next.execute(p)
		}
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	if e.next != nil {
		e.next.execute(p)
	}
}
func (e *doctor) setNext(next department) {
	e.next = next
}

//具体处理者（Concrete Handler）角色：实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
type reception struct {
	next department
}
func (e *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		if e.next != nil {
			e.next.execute(p)
		}
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	if e.next != nil {
		e.next.execute(p)
	}
}
func (e *reception) setNext(next department) {
	e.next = next
}

type patient struct {
	name string
	registrationDone bool
	doctorCheckUpDone bool
	medicineDone bool
	paymentDone bool
}