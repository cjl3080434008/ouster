package scene

import (
	"bytes"
	"github.com/tiancaiamao/ouster"
	"github.com/tiancaiamao/ouster/data"
	"github.com/tiancaiamao/ouster/packet"
	"github.com/tiancaiamao/ouster/packet/darkeden"
	// "github.com/tiancaiamao/ouster/skill"
	"log"
	"net"
	"time"
)

type PlayerClass uint8

const (
	_     = iota
	BRUTE = iota
)

type PlayerState uint8

const (
	STAND PlayerState = iota
	MOVE
)

// package scene have import player, if player import scene it would be a circle
// so use a interface to avoid direct use of scene.
type scene interface {
	Pos(uint32) (ouster.FPoint, error)
	To(uint32) (ouster.FPoint, error)
	Creature(uint32) ouster.Creature
	Monster(uint32)
	String() string
}

const (
	LEFT      = 53
	RIGHT     = 49
	UP        = 34
	DOWN      = 55
	LEFTUP    = 50
	RIGHTUP   = 48
	LEFTDOWN  = 52
	RIGHTDOWN = 54
)

// mostly the same as data.Player, but this is in memory instead.
type Player struct {
	Id   uint32 // set by scene.Login
	zone *Zone  // set by scene.Login

	name  string
	class PlayerClass
	hp    int
	mp    int
	speed float32
	level int

	strength     int
	agility      int
	intelligence int

	carried []int

	conn   net.Conn
	client <-chan packet.Packet
	send   chan<- packet.Packet
	aoi    <-chan ObjectIDType

	read      <-chan interface{}
	write     chan<- interface{}
	nearby    map[ObjectIDType]struct{}
	heartbeat <-chan time.Time
	ticker    uint32

	// Own by scene...write allowed only by scene agent
	State PlayerState
}

// implement Create
func (player *Player) Agility() int {
	return player.agility
}

func (player *Player) Strength() int {
	return player.strength
}

func (player *Player) Intelligence() int {
	return player.intelligence
}

func (player *Player) Damage() int {
	return player.strength
}

func (player *Player) Dodge() int {
	return player.agility
}

func (player *Player) ToHit() int {
	return player.agility
}

func (player *Player) HP() int {
	return 4*player.strength + player.level
}

// provide for scene to use
func (player *Player) Speed() float32 {
	return player.speed
}

func (player *Player) Defense() int {
	return player.strength
}

func NewPlayer(conn net.Conn, a <-chan ObjectIDType, rd <-chan interface{}, wr chan<- interface{}) *Player {
	return &Player{
		name:  "test",
		class: 1,
		hp:    110,
		mp:    110,
		conn:  conn,
		speed: 0.5,

		aoi:   a,
		read:  rd,
		write: wr,

		nearby:    make(map[ObjectIDType]struct{}),
		heartbeat: time.Tick(50 * time.Millisecond),
	}
}

func (player *Player) NearBy() map[ObjectIDType]struct{} {
	return player.nearby
}

func (player *Player) handleClientMessage(pkt packet.Packet) {
	hp := darkeden.GCStatusCurrentHP{
		ObjectID:  2351,
		CurrentHP: 133,
	}

	switch pkt.Id() {
	case darkeden.PACKET_CG_CONNECT:
		player.send <- &darkeden.GCUpdateInfoPacket{}
	case darkeden.PACKET_CG_READY:
		log.Println("get a CG Ready Packet!!!")
		player.send <- &darkeden.GCSetPositionPacket{
			X:   145,
			Y:   237,
			Dir: 2,
		}
	case darkeden.PACKET_CG_MOVE:
		player.write <- pkt
		// move := pkt.(darkeden.CGMovePacket)
		// moveOk := darkeden.GCMoveOKPacket{
		// 	Dir: move.Dir,
		// 	X:   move.X,
		// 	Y:   move.Y,
		// }
		// player.send <- moveOk

		// addBat := &darkeden.GCAddBat{
		// 	ObjectID:    2352,
		// 	MonsterName: "bat",
		// 	X:           149,
		// 	Y:           242,
		// 	Dir:         1,
		// 	CurrentHP:   111,
		// 	MaxHP:       133,
		// 	GuildID:     1,
		// }

	case darkeden.PACKET_CG_ATTACK:
		if hp.CurrentHP > 0 {
			hp.CurrentHP -= 5
			player.send <- hp
			player.send <- darkeden.GCAttackMeleeOK1(hp.ObjectID)
		}

	case darkeden.PACKET_CG_BLOOD_DRAIN:
	case darkeden.PACKET_CG_VERIFY_TIME:
	}
}

type BaseAttack struct{}

func (_ BaseAttack) ExecuteTarget(from, to ouster.Creature) (int, bool) {
	return 10, true
}

type SkillEffect struct {
	Id   int
	To   uint32
	Succ bool
	Hurt int
}

// func (player *Player) execute(pkt packet.SkillPacket) {
// 	skl := skill.Query(pkt.Id)
// 	switch skl.(type) {
// 	case skill.SelfSkill:
//
// 	case skill.TargetSkill:
// 		skill := skl.(skill.TargetSkill)
// 		target := player.Scene.Creature(pkt.Target)
// 		hurt, ok := skill.ExecuteTarget(player, target)
//
// 		player.write <- SkillEffect{
// 			Id:   pkt.Id,
// 			To:   pkt.Target,
// 			Succ: ok,
// 			Hurt: hurt,
// 		}
// 	case skill.RegionSkill:
// 	}
// }

type CMovePacketAck struct{}

func (this *Player) handleSceneMessage(msg interface{}) {
	switch msg.(type) {
	case darkeden.GCMoveOKPacket:
		raw := msg.(darkeden.GCMoveOKPacket)
		this.send <- raw
	default:
		log.Println("handleSceneMessage receive a unknown msg")
	}
}

func (this *Player) handleAoiMessage(id ObjectIDType) {
	if _, ok := this.nearby[id]; !ok {
		log.Println(id, "enter aoi...")
		this.nearby[id] = struct{}{}
		if id.Monster() {
			log.Println("it's a monster...send message")
			monster := this.zone.Monster(id.Index())
			info := data.MonsterType2MonsterInfo[monster.MonsterType]
			hp := info.STR*5 + uint16(info.Level)
			addMonster := &darkeden.GCAddMonster{
				ObjectID:    uint32(id),
				MonsterType: monster.MonsterType,
				MonsterName: info.Name,
				MainColor:   7,
				SubColor:    174,
				X:           uint8(monster.aoi.X()),
				Y:           uint8(monster.aoi.Y()),
				Dir:         2,
				CurrentHP:   monster.HP,
				MaxHP:       hp,
			}
			this.send <- addMonster
		}
	}

}

func (this *Player) heartBeat() {
	this.ticker++
}

func (this *Player) loop() {
	// var msg interface{}
	for {
		select {
		case msg, ok := <-this.client:
			if !ok {
				// kick the player off...
				return
			} else {
				// log.Println("before handleClientMessage...")
				this.handleClientMessage(msg)
				// log.Println("after handleClientMessage...")
			}
		// case msg = <-this.read:
		// log.Println("before handleSceneMessage...")
		// this.handleSceneMessage(msg)
		// log.Println("after handleSceneMessage...")
		// case id := <-this.aoi:
		// log.Println("before handleAoiMessage...")
		// log.Println("after handleAoiMessage...")
		case <-this.heartbeat:
			this.heartBeat()
		}
	}
}

func (player *Player) Go() {
	read := make(chan packet.Packet, 1)
	write := make(chan packet.Packet, 1)
	player.send = write
	player.client = read

	// open a goroutine to read from conn
	go func() {
		reader := darkeden.NewReader()
		for {
			data, err := reader.Read(player.conn)
			if err != nil {
				log.Println(err)
				player.conn.Close()
				close(read)
				return
			}
			// log.Println("packet before send to chan", data)
			read <- data
			// log.Println("packet after send to chan", data)
		}
	}()

	// open a goroutine to write to conn
	go func() {
		writer := darkeden.NewWriter()
		for {
			pkt := <-write
			log.Println("write channel get a pkt ", pkt.String())
			err := writer.Write(player.conn, pkt)
			if err != nil {
				log.Println(err)
				continue
			}

			buf := &bytes.Buffer{}
			writer.Write(buf, pkt)
			log.Println("send packet to client: ", buf.Bytes())
		}
	}()

	player.loop()
}
