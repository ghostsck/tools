package main

import (
	"fmt"
	//"testing"
	"tools/snowflake"
)

//
//func TestSnowFlakeByGo(t *testing.T) {
//	// 测试脚本
//
//	// 生成节点实例
//	worker, err := NewSnowFlake(1)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	ch := make(chan int64)
//	count := 10000
//	// 并发 count 个 goroutine 进行 snowflake ID 生成
//	for i := 0; i < count; i++ {
//		go func() {
//			id := worker.GetId()
//			ch <- id
//		}()
//	}
//
//	defer close(ch)
//
//	m := make(map[int64]int)
//	for i := 0; i < count; i++  {
//		id := <- ch
//		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
//		_, ok := m[id]
//		if ok {
//			t.Error("ID is not unique!\n")
//			return
//		}
//		// 将 id 作为 key 存入 map
//		m[id] = i
//	}
//	// 成功生成 snowflake ID
//	fmt.Println("All", count, "snowflake ID Get successed!")
//}

//var IdWorker *SnowFlake.Worker
//
func main() {
	// 传入当前节点id 此id在机器集群中一定要唯一 且从0开始排最多1024个节点，可以根据节点的不同动态调整该算法每毫秒生成的id上限(如何调整会在后面讲到)
	snow, _ := snowflake.NewSnowFlake(0)

	// 获得唯一id
	id := snow.GetId()
	fmt.Println(id)
}
