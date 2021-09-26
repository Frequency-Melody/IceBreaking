package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// validConfig 判断配置是否合法，注：现在只能判断单级配置, 没啥用，纯属写着玩
func (c *config) validConfig(rule func(v interface{}) (err error)) (valid bool) {
	valid = true // 合法性标志，全部字段不为空则合法\
	if c == nil {
		fmt.Println("config 读取失败")
		return false
	}
	it := c.structIterator()
	for {
		key, value, ok := it()
		if !ok {
			break
		}
		//fmt.Println(key, value)
		if err := rule(value); err != nil {
			fmt.Printf("配置错误：%s: %s", key, err)
			valid = false
		}
	}
	return
}

// structIterator, 结构体迭代器，每次返回一个结构体的属性名和值，用来检验配置合法性
// 接收者：c ：结构体信息
// 返回：key：结构体的属性名;value：属性对应的值；ok：当前值是否有效，为 false 代表已经遍历完
// 参考文献1：[Go语言笔记：struct结构遍历](https://www.cnblogs.com/nyist-xsk/p/9995321.html)
// 参考文献2：[Go语言设计模式实践：迭代器（Iterator）](https://studygolang.com/articles/1695)
func (c *config) structIterator() func() (key string, value interface{}, ok bool) {
	index := 0
	return func() (key string, value interface{}, ok bool) {
		// 按理说应该是传 c，但是现在只实现了单级配置的检验，所以传了 c.Mysql
		keys := reflect.TypeOf(c.Mysql)
		values := reflect.ValueOf(c.Mysql)
		keyNum := keys.NumField()
		if index >= keyNum {
			return "", nil, false
		}
		key = keys.Field(index).Name
		value = values.Field(index).Interface()
		index++
		return key, value, true
	}

}

//配置校验准则，每种类型都有自己的准则
//形参 : v：待判断的值（空接口）
//返回 ：err ：不合法返回错误
func validRule(v interface{}) (err error) {
	if v == nil {
		return errors.New("value is nil")
	}
	switch t := v.(type) {
	case string:
		if strings.TrimSpace(t) == "" {
			return errors.New("参数不能为空")
		}
		return nil
	case int:
		return nil
	case bool:
		return nil
	default:
		return nil
	}
}
