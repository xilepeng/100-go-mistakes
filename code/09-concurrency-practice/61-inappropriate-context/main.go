package main

import (
	"context"
	"net/http"
	"time"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		err := publish(r.Context(), response)
		// Do something with err
		_ = err
	}()

	writeResponse(response)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 异步处理发布操作，如果上下文被取消，发布消息的操作可以被中断
	// 不管写回 HTTP 响应需要多长时间，我们都可以调用 publish
	go func() {
		err := publish(context.Background(), response) // 使用空上下文而不是 HTTP 请求上下文
		// Do something with err
		_ = err
	}()

	writeResponse(response)
}

// 一旦我们返回响应，上下文就会被取消，异步操作也可能会被意外停止。
// 最佳实践：创建自定义上下文，不携带取消信号
// 传递给 publish 的上下文永远不会过期或被取消，但它将携带父上下文的值
func handler3(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 调用 publish 并分离取消信号
	go func() {
		// 在 HTTP 上下文之上使用分离
		err := publish(detach{ctx: r.Context()}, response)
		// Do something with err
		_ = err
	}()

	writeResponse(response)
}

// 创建自定义上下文，不携带取消信号
type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}

func doSomeTask(context.Context, *http.Request) (string, error) {
	return "", nil
}

func publish(context.Context, string) error {
	return nil
}

func writeResponse(string) {}
