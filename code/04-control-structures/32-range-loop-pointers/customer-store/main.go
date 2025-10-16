package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func main() {
	s := Store{
		m: make(map[string]*Customer),
	}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	print(s.m)
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer) //
		s.m[customer.ID] = &customer  //只创建一个固定地址的变量
	}
}

func (s *Store) storeCustomers2(customers []Customer) {
	for _, customer := range customers {
		current := customer // 创建一个局部变量
		s.m[current.ID] = &current
	}
}

func (s *Store) storeCustomers3(customers []Customer) {
	for i := range customers {
		customer := &customers[i] //赋值第i个元素的指针
		s.m[customer.ID] = customer
	}
}

func print(m map[string]*Customer) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%#v\n", k, v)
	}
}

/*
➜  customer-store git:(main) ✗ go run main.go
0x1400012a000
0x1400012a018
0x1400012a030
key=1, value=&main.Customer{ID:"1", Balance:10}
key=2, value=&main.Customer{ID:"2", Balance:-10}
key=3, value=&main.Customer{ID:"3", Balance:0}
*/
