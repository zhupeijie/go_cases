package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"

	job_demo "demo/job-demo"
	"demo/timer"
)

type ExtData struct {
	TeaID int64 `json:"tea_id"`
}

const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m4 = 0x0f0f0f0f // 00001111000011110000111100001111
	m8 = 0x00ff00ff // 00000000111111110000000011111111
)

func main() {

	ticker := timer.NewTicker(1, func() {
		fmt.Println("here is test job ", time.Now().Unix())
	})

	defer ticker.Stop()
	ticker.Tick()
	// now := time.Now().UnixNano() / 1e6
	// fmt.Println("---", now)
	// fmt.Println(1e6)
	// 1663205130400
	// 1663121601000
	// return
	return

	// c := job_demo.TJob()
	job_demo.TJob()
	// defer func() {
	// 	c.Stop()
	// }()

	g := gin.New()

	gracehttp.Serve(&http.Server{Addr: ":8081", Handler: g})
	// phone := "13849013659"

}
func reverseBits(n uint32) uint32 {
	n = n>>1&m1 | n&m1<<1
	n = n>>2&m2 | n&m2<<2
	n = n>>4&m4 | n&m4<<4
	n = n>>8&m8 | n&m8<<8
	return n>>16 | n<<16
}

func uuidT() {
	// t1 := time.Now()
	// times := 1
	// for i := 0; i < times; i++ {
	// 	uuid.NewV4().String()
	// go mod
	// }
	// t2 := time.Now()
	fmt.Println(" == ", strings.Replace(uuid2.New().String(), "-", "", -1))

	// for i := 0; i < times; i++ {
	// 	uuid2.New().String()
	// }
	// t3 := time.Now()
	// fmt.Println("++++++", t3.Sub(t2))
	// 50c8389b-fcea-4c22-96de-7006dadea4bc
}
