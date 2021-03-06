package radixpool

import (
	"fmt"
	"testing"
	"time"
)

func Test_Pool(t *testing.T) {
	pool, err := NewPool("tcp", "127.0.0.1:6379", 10, 3*time.Second, "", 0)

	rc, err := pool.Get()

	rc.Conn.Cmd("SET", "a", "b")

	str, _ := rc.Conn.Cmd("GET", "a").Str()
	fmt.Println("GET1", str)

	pool.CarefullyPut(rc, &err)

	time.Sleep(4 * time.Second)

	rc, err = pool.Get()

	str, _ = rc.Conn.Cmd("GET", "a").Str()
	fmt.Println("GET2", str)

	pool.CarefullyPut(rc, &err)

	str, _ = pool.Cmd("GET", "a").Str()
	fmt.Println("GET3", str)

}
