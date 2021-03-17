package main

import (
	"fmt"
	"testing"
	"time"
)

//func TestCowQueue(t *testing.T) {
//	queue := NewQueue()
//	queue.Enqueue(1)
//	//queue.Enqueue(2)
//	//queue.Enqueue(3)
//	//
//	//val := queue.Dequeue()
//	//if 1 != val {
//	//	t.Errorf("1 is not %d", val)
//	//}
//	//
//	//val = queue.Dequeue()
//	//if 2 != val {
//	//	t.Errorf("2 is not %d", val)
//	//}
//	//
//	//val = queue.Dequeue()
//	//if 3 != val {
//	//	t.Errorf("3 is not %d", val)
//	//}
//}

//func assert(expect, actual int , t *testing.T) {
//	if expect != actual {
//		t.Errorf("expect is %d but actual %d",expect, actual)
//	}
//}

//func TestQueue_Enqueue(t *testing.T) {
//	fmt.Println("TestQueue_Enqueue start")
//	queue := NewQueue()
//	//for i:=0; i < 10000000; i++ {
//	//for i := 0; i < 10000000; i++ {
//	//	queue.Enqueue(i)
//	//	//fmt.Println(i)
//	//}
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			queue.Enqueue(i)
//			fmt.Println(i)
//		}(i)
//	}
//	size := queue.Size()
//	fmt.Printf("size: %d \n", size)
//	for ; size > 0; {
//		//val := queue.Dequeue()
//		_ = queue.Dequeue()
//		size = queue.Size()
//		//fmt.Printf("value: %d size: %d\n", val, size)
//	}
//}

func TestQueue_EnqueueV2(t *testing.T) {
	fmt.Println("TestQueue_Enqueue start")
	queue := NewQueue()
	//for i:=0; i < 10000000; i++ {
	enquque(0, 2, queue)
	enquque(1, 2, queue)
	size := queue.Size()
	//time.Sleep(time.Second * 1)
	fmt.Printf("size: %d \n", size)
	for size > 0 {
		val := queue.Dequeue()
		size = queue.Size()
		fmt.Printf("value: %d size: %d\n", val, size)
	}
	time.Sleep(time.Second * 1)
}

func enquque(start, inc int, queue *Queue) {
	for i := start; i < 1000; i = i + inc {
		go func(i int) {
			queue.Enqueue(i)
			fmt.Printf("for go %d\n", queue.Size())
			fmt.Println(i)
		}(i)
	}
	fmt.Println("enqueu end")
}
