package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
import "golang.org/x/sync/errgroup"

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	s := http.Server{
		Addr:    ":8080",
		Handler: &pingHandler{},
	}
	startHttpServer := func() error {
		fmt.Println("http server start")
		// s.Shutdown()调用时，退出
		return s.ListenAndServe()
	}
	listenOsSignal := func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-c:
			// 接收到信号，调用s.Shutdown()，退出
			fmt.Println("listenOsSignal: signal is coming")
			// 尽管这样会闭包一个s，但这种实现方式，
			// 比再起一个goroutine 等待ctx.done()返回，然后再执行s.Shutdown()
			// 逻辑上更清晰和直接
			return s.Shutdown(context.Background())
		case <-ctx.Done():
			// httpServer 报错，退出
			fmt.Println("listenOsSignal: http server happened error")
			return fmt.Errorf("http server happened error:%w", ctx.Err())
		}
	}
	eg.Go(startHttpServer)
	eg.Go(listenOsSignal)
	if err := eg.Wait(); err != nil {
		fmt.Println("http server exit")
		fmt.Printf("errgroup err:%v \r\n", err)
	}
	fmt.Println("exit")
}

type pingHandler struct {
}

func (p *pingHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "week03 \r\n")
}
