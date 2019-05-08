package main

import "fmt"

// var nsList luban

type luban struct {
	nss []nsArray
}

type nsArray struct {
	namespace string
	state     string
	rcs       []rcArray
}
type rcArray struct {
	rcname string
	state  string
	pods   []podArray
}
type podArray struct {
	podName     string
	state       string
	instanceNum string
	IP          string
}

//父类
type BaseNum struct {
	num1 int
	num2 int
}

//加法子类
type Add struct {
	BaseNum
}

//减法子类
type Sub struct {
	BaseNum
}

//子类方法
func (a *Add) Opt() int {
	return a.num1 + a.num2
}

//子类方法
func (a *Add) Test() int {
	return a.num1 + a.num2
}
func (s *Sub) Opt() int {
	return s.num1 - s.num2
}

//定义接口, 即封装

type Opter interface {
	Opt() int
	Test() int
}

//定义空类 以产生 工厂模式 的方法
type lubanDB struct {
}

func (l *lubanDB) initDB(a, b int) int {
	var i Opter
	AddNum := Add{BaseNum{a, b}}
	i = &AddNum
	return i.Test()
}

func main() {
	var db lubanDB
	value := db.initDB(20, 3)
	fmt.Println(value)
}
