package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type MyError struct {
	Code    int
	Message string
}

// 实现 Error() 方法
func (e MyError) Error() string {
	return fmt.Sprintf("Code:%d,Message:%s", e.Code, e.Message)
}

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		// return "", err
		// return "", errors.New("文件不存在") // 自定义错误
		// return "", fmt.Errorf("文件不存在（也可以用fmt.Errorf 返回）") // 自定义错误
		return "", &MyError{Code: 400, Message: err.Error()}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, err
}

// 创建一个错误
var err1 = errors.New("我是错误1")
var err2 = fmt.Errorf("我是错误2")

func f() error {
	return err1
}

func f2() error {
	return &MyError{Code: 500, Message: "not found:我是断言错误"}
}

func main() {
	// filename := "unfile.txt"
	filename := "file.txt"
	content, err := readFile(filename)
	if err != nil {
		fmt.Println("读文件发生错误：", err)
		return
	}
	fmt.Println(content)

	// errors.Is() // 判断一个错误是否等于另一个错误
	reserr1 := f()

	if errors.Is(reserr1, err1) {
		fmt.Println("我们是同一个错误或包装了同一个错误")
	} else {
		fmt.Println("我们不是同一个错误")
	}

	// errors.As 方法来检查被包装的错误是否是属于某种类型的错误。
	// errors.As 递归地解开一个错误，如果包装链中的一个错误符合预期的类型，则返回 true
	var myErr *MyError
	err = f2()
	if errors.As(err, &myErr) {
		fmt.Printf("MyError:code=%d,message=%s", myErr.Code, myErr.Message)
	} else {
		fmt.Println(err)
	}

}
