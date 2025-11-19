package main

import (
	"fmt"
	"sync"
)

func main() {
	customer := Customer{}
	_ = customer.UpdateAge1(-1) // fatal error: all goroutines are asleep - deadlock!
	//_ = customer.UpdateAge2(-1) // id: , age: 0
	//_ = customer.UpdateAge3(-1) // id: , age: 0
}

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

// 从上下文格式化一个键，返回一个格式化结构体的错误，
// 在这两种情况下，格式化字符串都会导致：数据竞争和死锁
// 因为 UpdateAge1 已经获得了互斥锁，所以 String 方法将无法获得它，导致死锁
func (c *Customer) UpdateAge1(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// 结构体定义了String方法，%v会调用该方法获取值
	// %v	the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}

	c.age = age
	return nil
}

// fatal error: all goroutines are asleep - deadlock!

// 首先检查输入，如果输入有效，则获取锁
func (c *Customer) UpdateAge2(age int) error {
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.age = age
	return nil
}

// 不调用 String(), 直接访问 id 字段，不会导致死锁
func (c *Customer) UpdateAge3(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer id %s", c.id)
	}

	c.age = age
	return nil
}

// 结构体定义了String方法，%v会调用该方法获取值
func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id: %s, age: %d", c.id, c.age)
}
