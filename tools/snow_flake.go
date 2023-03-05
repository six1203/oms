package tools

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// 设置起始时间戳, time.Now().UnixNano()
	epoch = int64(1678015337674)

	nodeBits  = uint(10)              // 节点 ID 所占位数
	stepBits  = uint(12)              // 序列号所占位数
	nodeMax   = -1 ^ (-1 << nodeBits) // 最大节点 ID
	stepMask  = -1 ^ (-1 << stepBits) // 序列号掩码
	timeShift = nodeBits + stepBits   // 时间戳左移位数
	nodeShift = stepBits              // 节点 ID 左移位数
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	node      int64
	step      int64
}

func NewSnowFlake(node int64) (*Snowflake, error) {
	if node < 0 || node > nodeMax {
		return nil, fmt.Errorf("节点 ID 超出范围: %d", node)
	}
	return &Snowflake{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1000000
	if s.timestamp == now {
		s.step = (s.step + 1) & stepMask
		if s.step == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.step = 0
	}

	s.timestamp = now
	id := (now-epoch)<<timeShift | (s.node << nodeShift) | s.step
	return id
}

func GenerateId() int64 {
	// 生成种子
	rand.Seed(time.Now().UnixNano())

	// 随机选择一个节点
	node := rand.Intn(10) + 1

	// 生成雪花算法实例
	snowFlake, err := NewSnowFlake(int64(node))
	if err != nil {
		panic(err)
	}
	return snowFlake.NextID()
}
