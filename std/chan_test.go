package std

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {

	//ch := NewChan(1)
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	ch.Send("test")
	//}()
	//v := ch.Recv()
	//fmt.Println(v.IsNil(), "-", v.IsEmpty(), "-", v.String())

	//ch := NewChan(10)
	//go func() {
	//	for k := 0; k <= 7; k++ {
	//		ch.Send(k)
	//	}
	//	ch.Close()
	//}()
	//for k := 0; k < 13; k++ {
	//	v := ch.Recv()
	//	fmt.Println(v.IsNil(), "-", v.IsEmpty(), "-", v.String())
	//	time.Sleep(time.Second)
	//}

	ch := NewChan(10)
	go func() {
		for k := 0; k <= 9; k++ {
			fmt.Println("send:", k, "-", ch.Send(k))
			time.Sleep(1 * time.Second)
		}
	}()
	for k := 0; k < 6; k++ {
		v := ch.Recv()
		fmt.Println(v.IsNil(), "-", v.IsEmpty(), "-", v.String())
		time.Sleep(time.Second)
	}
	ch.Close()
	time.Sleep(5 * time.Second)

}
