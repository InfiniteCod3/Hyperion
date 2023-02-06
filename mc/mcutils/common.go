package mcutils

import (
	"Hyperion/mc"
	"Hyperion/mc/packet"
)

type nextState int

const (
	Status nextState = 1
	Login  nextState = 2
)

const loginStart = 0x00

func WriteHandshake(conn *mc.Connection, ip string, port int, protocol int, nextState nextState) (err error) {
	err = conn.WritePacket(
		packet.Marshal(
			0x00,
			packet.VarInt(protocol),
			packet.String(ip),
			packet.UnsignedShort(port),
			packet.VarInt(nextState),
		),
	)
	return
}

func WriteLoginPacket(conn *mc.Connection, name string, hasUUID bool, uuid *packet.UUID) (err error) {
	if hasUUID {
		err = conn.WritePacket(
			packet.Marshal(
				0x00,
				packet.String(name),
				packet.Boolean(hasUUID),
				uuid,
			),
		)
	} else {
		err = conn.WritePacket(
			packet.Marshal(
				0x00,
				packet.String(name),
				packet.Boolean(hasUUID),
			),
		)
	}
	return
}
