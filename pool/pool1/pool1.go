package pool1

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Conn struct {
	maxConn       int                     // 最大连接数
	maxIdle       int                     // 最大可用连接数
	freeConn      int                     // 线程池空闲连接数
	connPool      []int                   // 连接池
	openCount     int                     // 已经打开的连接数
	waitConn      map[int]chan Permission // 排队等待的连接队列
	waitCount     int                     // 等待个数
	lock          sync.Mutex              // 锁
	nextConnIndex NextConnIndex
	freeConns     map[int]Permission // 连接池的连接
	expiredCh     chan string
}

type Config struct {
	MaxConn int
	MaxIdle int
}

type Permission struct {
	NextConnIndex
	Content     string
	CreatedAt   time.Time
	MaxLifeTime time.Duration
}

type NextConnIndex struct {
	Index int
}

var nowFunc = time.Now

// Ping Ping是一个获取连接并释放连接的过程
func Ping() {

}

// TODO 1、使用channel归还连接 2、连接添加超时时间，如果超过超时时间没有拿到连接则退出
func Prepare(ctx context.Context, config *Config) (conn *Conn) {
	go func() {
		//for {
		//conn.expiredCh = make(chan string, len(conn.freeConns))
		//for _, value := range conn.freeConns {
		//	if value.CreatedAt.Add(value.MaxLifeTime).Before(nowFunc()) {
		//		conn.expiredCh <- "CLOSE"
		//	}
		//}
		//select {
		//case <-conn.waitConn:
		//	fmt.Println("receive")
		//	conn.lock.Lock()
		//	if conn.waitCount > 1 {
		//		conn.waitCount--
		//	}
		//	conn.lock.Unlock()
		//}
		//}

	}()
	return &Conn{
		maxConn:   config.MaxConn,
		maxIdle:   config.MaxIdle,
		openCount: 0,
		connPool:  []int{},
		waitConn:  make(map[int]chan Permission),
		waitCount: 0,
		freeConns: make(map[int]Permission),
	}
}

// 创建连接
func (conn *Conn) New(ctx context.Context) (permission Permission, err error) {
	/**
	1、如果当前连接池已满，即len(freeConns)=0
	2、判定openConn是否大于maxConn，如果大于，则丢弃获取加入队列进行等待
	3、如果小于，则考虑创建新连接
	注意设置超时时间以及用完连接的回收
	*/
	conn.lock.Lock()

	select {
	default:
	case <-ctx.Done():
		conn.lock.Unlock()

		return Permission{}, errors.New("new conn failed, context cancelled!")
	}

	if len(conn.freeConns) > 0 {
		var (
			popPermission Permission
			popReqKey     int
		)
		for popReqKey, popPermission = range conn.freeConns {
			break
		}
		delete(conn.freeConns, popReqKey)
		fmt.Println("log", "use free conn!!!!!", "openCount: ", conn.openCount, " freeConns: ", conn.freeConns)
		conn.lock.Unlock()
		return popPermission, nil
	}

	if conn.openCount >= conn.maxConn { // 当前连接数大于上限，则加入等待队列
		//conn.lock.Lock()
		//conn.waitConn <- 1
		nextConnIndex := getNextConnIndex(conn)

		req := make(chan Permission, 1)
		conn.waitConn[nextConnIndex] = req
		//conn.waitConn = append(conn.waitConn, 1) // TODO 等待队列的处理
		conn.waitCount++
		conn.lock.Unlock()

		select {
		case <-time.After(time.Second * time.Duration(3)):
			fmt.Println("超时，通知主线程退出")
			return
		case ret, ok := <-req: // 有放回的连接, 直接拿来用
			if !ok {
				return Permission{}, errors.New("new conn failed, no available conn release")
			}
			fmt.Println("log", "received released conn!!!!!", "openCount: ", conn.openCount, " freeConns: ", conn.freeConns)
			return ret, nil
		}
		return Permission{}, errors.New("new conn failed")
	}

	// 新建连接
	conn.openCount++
	conn.lock.Unlock()
	nextConnIndex := getNextConnIndex(conn)
	permission = Permission{NextConnIndex: NextConnIndex{nextConnIndex},
		Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
	//conn.freeConns[nextConnIndex] = permission
	fmt.Println("log", "create conn!!!!!", "openCount: ", conn.openCount, " freeConns: ", conn.freeConns)
	return permission, nil
}

func getNextConnIndex(conn *Conn) int {
	currentIndex := conn.nextConnIndex.Index
	//conn.lock.Lock()
	conn.nextConnIndex.Index = currentIndex + 1
	//fmt.Println("connIndex: ", conn.nextConnIndex)
	//conn.lock.Lock()
	return conn.nextConnIndex.Index
}

// 释放连接
func (conn *Conn) Release(ctx context.Context) (result bool, err error) {
	conn.lock.Lock()
	//conn.openCount--
	//conn.freeConn++
	//conn.lock.Unlock()
	//fmt.Println("openCount: ", conn.openCount, " freeConns: ", conn.freeConns)

	if len(conn.waitConn) > 0 {
		var req chan Permission
		var reqKey int
		for reqKey, req = range conn.waitConn {
			break
		}

		permission := Permission{NextConnIndex: NextConnIndex{reqKey},
			Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
		req <- permission
		//conn.freeConns[reqKey] = permission
		//conn.lock.Lock()
		conn.waitCount--
		delete(conn.waitConn, reqKey)
		conn.lock.Unlock()
	} else {
		if conn.openCount > 0 {
			conn.openCount--

			if len(conn.freeConns) < conn.maxIdle { // 确保连接池大小不会超过maxIdle
				nextConnIndex := getNextConnIndex(conn)
				permission := Permission{NextConnIndex: NextConnIndex{nextConnIndex},
					Content: "PASSED", CreatedAt: nowFunc(), MaxLifeTime: time.Second * 5}
				conn.freeConns[nextConnIndex] = permission
			}
		}
		conn.lock.Unlock()
	}
	return
}
