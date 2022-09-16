package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"

	"hotup/srv"
)

// 热更新实现s
func main() {
	// HotByGrace()
	HotByShutdown()
}

func HotByShutdown() {
	e := gin.Default()


	// 注册router
	router(e)
	tl, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	sr := http.Server{
		Addr:    ":8002",
		Handler: e,
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("here is err %+v\n", err)
	}
	fmt.Printf("here is path, %s\n", wd)
	srv.TestGetWd()
	go listenSign(context.Background(), &sr)
	err = sr.Serve(tl) // 启动服务
	if err != nil {
		panic(err)
	}
}

func listenSign(ctx context.Context, srv *http.Server) {
	fmt.Println("here is listen sign to close")
	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		select {
		case <-signCh:
			fmt.Printf(" here has get sys sign")
			err := srv.Shutdown(context.Background())
			if err != nil {
				panic(err)
			}
			return
		default:

		}
	}
}

func HotByGrace() {
	e := gin.Default()

	// 注册router
	router(e)
	server := &http.Server{
		Addr:    ":8001",
		Handler: e,
	}
	err := gracehttp.Serve(server)
	if err != nil {
		panic(err)
	}
}

func router(e *gin.Engine) {
	e.GET("/hello", SayHello)
}

func SayHello(c *gin.Context) {
	t := c.Query("t")
	fmt.Println("here is  ", t)
	t1, _ := strconv.Atoi(t)
	if t1 == 0 {
		t1 = 2
	}
	time.Sleep(time.Second * time.Duration(t1))
	c.JSON(200, "success3333")
}
