package main

import (
    "github.com/tiancaiamao/ouster/packet"
)

// 传interface相当于传指针，消息更小
// 增加了解包的代价
type AgentMessage interface {
    Sender() *Agent
}

func (agent *Agent) Sender() *Agent {
    return agent
}

type MoveMessage struct {
    *Agent
    packet.CGMovePacket
}

type SkillBroadcastMessage struct {
    *Agent
    packet.Packet
}

type MeteorStrikeMessage struct {
    *Agent
    EffectMeteorStrike
    X   uint8
    Y   uint8
}