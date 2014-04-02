package main

import (
	// "github.com/tiancaiamao/ouster"
	"github.com/tiancaiamao/ouster/config"
	"github.com/tiancaiamao/ouster/login"
	"github.com/tiancaiamao/ouster/packet"
	"github.com/tiancaiamao/ouster/player"
	"github.com/tiancaiamao/ouster/scene"
	"log"
	"net"
)

func main() {
	log.Println("Starting the server.")

	scene.Initialize()

	listener, err := net.Listen("tcp", config.ServerPort)
	checkError(err)

	log.Println("Game Server OK.")

	for {
		conn, err := listener.Accept()
		if err == nil {
			go handleClient(conn)
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleClient(conn net.Conn) {
	log.Println("accept a connection...")
	// check username/password, load player info, and so on...
	playerData, err := login.Login(conn)
	if err != nil {
		return
	}

	// get the map that player current in and connect to it
	m := scene.Query(playerData.Map)

	ch := make(chan interface{})
	// pos := ouster.FPoint{
	// 	X: float32(playerData.Pos.X),
	// 	Y: float32(playerData.Pos.Y),
	// }

	agent := new(player.Player)
	agent.Player2scene = make(chan interface{})
	agent.Scene2player = make(chan interface{})
	playerId, succ := m.Login(agent)

	// turn into a player agent
	if succ {
		// open a goroutine to read from conn
		go func() {
			for {
				data, err := packet.Read(conn)
				if err != nil {
					// write a reset packet...
					continue
				}
				ch <- data
			}
		}()

		agent.Init(playerId, playerData, conn, ch)
		agent.Go()
	}
}
