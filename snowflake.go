package snowflake

import (
	"errors"
	"fmt"
	"time"
)

/*
twitter的snowflake算法 -- go实现


*/

const (
	START_STMP int64 = 1572143464315 // 起始时间戳

	SEQUENCE_BIT   uint8 = 12 // 序号
	MACHINE_BIT    uint8 = 5  //机器标识占用的位数
	DATACENTER_BIT uint8 = 5  //数据中心占用的位数
	TIMESTMP_BIT   uint8 = 41

	//  每部分最大
	MAX_SEQUENCE         = -1 ^ (-1 << SEQUENCE_BIT)
	MAX_MACHINE          = -1 ^ (-1 << MACHINE_BIT)
	MAX_DATACENTER int64 = -1 ^ (-1 << DATACENTER_BIT)
	MAX_TIMESTMP   int64 = -1 ^ (-1 << TIMESTMP_BIT)

	// 每部分左移的位置
	MACHINE_LEFT    = SEQUENCE_BIT
	DATACENTER_LEFT = MACHINE_LEFT + MACHINE_BIT
	TIMESTMP_LEFT   = DATACENTER_LEFT + DATACENTER_BIT
)

type SnowFlake struct {
	dataCenterId uint64 // 数据中心
	machineId    uint64 //   机器
	sequence     uint64 // 序号
	lastStmp     int64  // 上一次时间戳
}

func NewSnowFlake(datacenterId, machineId int64) (*SnowFlake, error) {
	if datacenterId > MAX_DATACENTER || datacenterId < 0 {
		return nil, errors.New("Illegal Argument Exception :datacenterId  out  of  range")
	}
	if machineId > MAX_MACHINE || machineId < 0 {
		return nil, errors.New("Illegal Argument Exception :machineId  out  of  range")
	}

	return &SnowFlake{
		dataCenterId: uint64(datacenterId),
		machineId:    uint64(machineId),
		sequence:     0,
		lastStmp:     0,
	}, nil
}

func (self *SnowFlake) NextID() (uint64, error) {
	var (
		currTimeStamp int64
		sequence      uint64
	)
	currTimeStamp = self.currentTimeStampMilliSecond()
	if currTimeStamp < self.lastStmp {
		return 0, errors.New("current timestamp less than last timestamp")
	}

	// 同一毫秒内生成
	if currTimeStamp == self.lastStmp {
		self.sequence = (self.sequence + 1) & MAX_SEQUENCE
		if self.sequence == 0 {
			currTimeStamp = self.nextTimeStampMilliSecond()
		}

	} else {
		sequence = 0
	}
	self.lastStmp = currTimeStamp
	self.sequence = sequence
	a := uint64(currTimeStamp-START_STMP)<<TIMESTMP_LEFT |
		self.dataCenterId<<DATACENTER_LEFT |
		self.machineId<<MACHINE_LEFT |
		sequence
	fmt.Println(a)
	return a, nil

}

// 获取当前时间戳 毫秒级别
func (self *SnowFlake) currentTimeStampMilliSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取下一毫秒的时间戳
func (self *SnowFlake) nextTimeStampMilliSecond() int64 {
	var (
		currTimeStamp int64
	)
TOP:
	currTimeStamp = self.currentTimeStampMilliSecond()
	if currTimeStamp <= self.lastStmp {
		goto TOP

	}
	return currTimeStamp
}
