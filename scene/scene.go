package scene

import (
	"github.com/tiancaiamao/ouster/aoi"
	"github.com/tiancaiamao/ouster/packet/darkeden"
	// "log"
)

func loop(m *Zone) {
	for {
		for id, player := range m.players {
			select {
			case msg := <-player.agent2scene:
				m.processPlayerInput(uint32(id), msg)
			default:
				break
			}
		}

		select {
		case <-m.quit:
			// 处理退出消息
		case <-m.event:
			// 处理地图事件...比如boss刷了之类的
		case <-m.heartbeat:
			//50ms的心跳,目前还不确定做什么...npc,怪物...player的逻辑放到
			//自身的goroutine中
			m.HeartBeat()
		}
	}
}

func (m *Zone) Go() {
	go loop(m)
}

func (m *Zone) processPlayerInput(playerId uint32, msg interface{}) {
	switch msg.(type) {
	case darkeden.CGMovePacket:
		move := msg.(darkeden.CGMovePacket)
		player := m.Player(playerId)
		player.send <- darkeden.GCMoveOKPacket{
			X:   move.X,
			Y:   move.Y,
			Dir: move.Dir,
		}
		m.aoi.Nearby(uint16(move.X), uint16(move.Y), func(entity *aoi.Entity) {
			dx := int(move.X) - int(entity.X())
			dy := int(move.Y) - int(entity.Y())
			if dx*dx+dy*dy <= 64 {
				player.handleAoiMessage(ObjectIDType(entity.Id()))
			}
		})

		// case player.SkillEffect:
		// 	log.Println("scene receive and process a SkillEffect")
		// 	// raw := msg.(player.SkillEffect)
		// 	handle := m.Player(playerId)
		// 	pc := handle.pc
		//
		// 	nearby := pc.NearBy()
		// 	for _, playerId := range nearby {
		// 		p := m.Player(playerId)
		// 		if p != nil {
		// 			// p.write <- packet.SkillTargetEffectPacket{
		// 			// 	Skill: raw.Id,
		// 			// 	From:  playerId,
		// 			// 	To:    raw.To,
		// 			// 	Hurt:  raw.Hurt,
		// 			// 	Succ:  raw.Succ,
		// 			// }
		// 		}
		// 	}
	}
}
