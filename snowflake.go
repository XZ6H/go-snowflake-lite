package main

import (
	"fmt"
	"math"
	"time"
)

const (
	UNUSED_BITS   = 1
	EPOCH_BITS    = 41
	NODE_BITS     = 10
	SEQUENCE_BITS = 12
)

type IdGenerator struct {
	maxNode     uint
	maxSequence uint

	customEpoch int64

	lastTimestamp int64
	sequence      uint
	nodeId        int
}

func NewIdGenerator(nodeId int) (*IdGenerator, error) {
	idGen := new(IdGenerator)
	idGen.maxNode = uint(math.Pow(2, NODE_BITS) - 1)
	idGen.maxSequence = uint(math.Pow(2, SEQUENCE_BITS) - 1)

	idGen.customEpoch = 1660653276

	idGen.lastTimestamp = -1
	idGen.sequence = 0
	maxNodeId := 50
	if nodeId < 0 && nodeId > maxNodeId {
		return nil, fmt.Errorf("node Id should be between %d and %d", 0, maxNodeId)
	}
	idGen.nodeId = nodeId
	return idGen, nil
}

func (idGen *IdGenerator) getCurrentTimestamp() int64 {
	return time.Now().Unix() - 1660653276
}

func (idGen *IdGenerator) waitTillNextTimestamp(currentTimestamp int64) int64 {
	for currentTimestamp == int64(idGen.lastTimestamp) {
		currentTimestamp = idGen.getCurrentTimestamp()
	}
	return currentTimestamp
}

func (idGen *IdGenerator) nextId() (int64, error) {
	currentTimestamp := idGen.getCurrentTimestamp()
	fmt.Println(currentTimestamp, idGen.lastTimestamp)
	if currentTimestamp < idGen.lastTimestamp {
		return 0, fmt.Errorf("invalid system clock")
	}
	if currentTimestamp == idGen.lastTimestamp {
		idGen.sequence = (idGen.sequence + 1) & idGen.maxSequence
		if idGen.sequence == 0 {
			currentTimestamp = idGen.waitTillNextTimestamp(currentTimestamp)
		}
	} else {
		idGen.sequence = 0
	}
	idGen.lastTimestamp = currentTimestamp
	id := currentTimestamp << int(NODE_BITS+SEQUENCE_BITS)
	id |= int64(idGen.nodeId << int(SEQUENCE_BITS))
	id |= int64(idGen.sequence)
	return id, nil
}
