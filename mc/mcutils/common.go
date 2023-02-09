package mcutils

import (
	"Hyperion/mc/mcversions"
	"Hyperion/mc/packet"
)

type nextState int

const (
	Status nextState = 1
	Login  nextState = 2
)

func GetHandshakePacket(ip string, port int, protocol int, nextState nextState) (pk packet.Packet) {
	pk = packet.Marshal(
		0x00,
		packet.VarInt(protocol),
		packet.String(ip),
		packet.UnsignedShort(port),
		packet.VarInt(nextState),
	)
	return
}

func GetLoginPacket(name string, versionProtocol int) (pk packet.Packet) {
	if versionProtocol == mcversions.V1_19_1 || versionProtocol == mcversions.V1_19 {
		pk = packet.Marshal(
			0x00,
			packet.String(name),
			packet.Boolean(false),
			packet.Boolean(false),
		)
	} else if versionProtocol > mcversions.V1_19_1 {
		pk = packet.Marshal(
			0x00,
			packet.String(name),
			packet.Boolean(false),
		)
	} else {
		pk = packet.Marshal(
			0x00,
			packet.String(name),
		)
	}
	return
}
