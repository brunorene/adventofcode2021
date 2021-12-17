package packet

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

const (
	Sum     = 0
	Mult    = 1
	Min     = 2
	Max     = 3
	Value   = 4
	Greater = 5
	Less    = 6
	Equal   = 7
)

var zeroRightPad = regexp.MustCompile(`^[0]*$`)

type (
	Header interface {
		GetType() PacketType
		GetVersion() byte
	}

	Packet interface {
		Header
		GetResult() *big.Int
		GetChildren() []Packet
	}

	PacketType byte

	Message struct {
		packets []Packet
	}

	PacketHeader struct {
		Version byte
		Type    PacketType
	}

	LiteralValuePacket struct {
		PacketHeader
		Value *big.Int
	}

	OperatorPacket struct {
		PacketHeader
		Children []Packet
	}
)

func (p LiteralValuePacket) GetResult() *big.Int {
	return p.Value
}

func (p OperatorPacket) GetResult() *big.Int {
	switch p.Type {
	case Sum:
		sum := big.NewInt(0)
		for _, p := range p.Children {
			sum.Add(sum, p.GetResult())
		}

		return sum
	case Mult:
		sum := big.NewInt(1)
		for _, p := range p.Children {
			sum.Mul(sum, p.GetResult())
		}

		return sum
	case Min:
		return minMax(p.Children, -1)
	case Max:
		return minMax(p.Children, 1)
	case Less:
		return compare(p.Children, -1)
	case Greater:
		return compare(p.Children, 1)
	case Equal:
		return compare(p.Children, 0)
	default:
		panic("unknown operation")
	}
}

func (p LiteralValuePacket) GetChildren() []Packet {
	return []Packet{}
}

func (p OperatorPacket) GetChildren() []Packet {
	return p.Children
}

func (p PacketHeader) GetType() PacketType {
	return p.Type
}

func (p PacketHeader) GetVersion() byte {
	return p.Version
}

func compare(packets []Packet, outcome int) *big.Int {
	if len(packets) != 2 {
		panic("min: len packets != 2")
	}
	if packets[0].GetResult().Cmp(packets[1].GetResult()) == outcome {
		return big.NewInt(1)
	}

	return big.NewInt(0)
}

func minMax(packets []Packet, outcome int) (chosen *big.Int) {
	chosen = packets[0].GetResult()
	if len(packets) == 1 {
		return
	}
	if packets[0].GetResult().Cmp(packets[1].GetResult()) != outcome {
		chosen = packets[1].GetResult()
	}
	for i := 0; i < len(packets); i++ {
		if chosen.Cmp(packets[i].GetResult()) != outcome {
			chosen = packets[i].GetResult()
		}
	}

	return
}

func newPacket(in string) (Packet, string) {
	packetType, err := strconv.ParseUint(in[3:6], 2, 8)
	check(err)

	switch PacketType(packetType) {
	case Value:
		return newLiteralValuePacket(in)
	default:
		return newOperatorPacket(in)
	}
}

func newLiteralValuePacket(in string) (packet LiteralValuePacket, out string) {
	version, err := strconv.ParseUint(in[0:3], 2, 8)
	check(err)
	packetType, err := strconv.ParseUint(in[3:6], 2, 8)
	check(err)

	header := PacketHeader{byte(version), PacketType(packetType)}

	var number string
	for i := 6; ; i += 5 {
		groupBit := in[i : i+1]
		number += in[i+1 : i+5]
		if groupBit == "0" { //last group
			out = in[i+5:]
			break
		}
	}
	val, success := new(big.Int).SetString(number, 2)
	if !success {
		panic("fail big int")
	}
	return LiteralValuePacket{header, val}, out
}

func newOperatorPacket(in string) (packet OperatorPacket, out string) {
	version, err := strconv.ParseUint(in[0:3], 2, 8)
	check(err)
	packetType, err := strconv.ParseUint(in[3:6], 2, 8)
	check(err)

	header := PacketHeader{byte(version), PacketType(packetType)}

	var packets []Packet

	switch in[6] {
	case '0': // 15 bit == subpacket length
		length, err := strconv.ParseUint(in[7:22], 2, 16)
		check(err)
		packets, out = NewMessage(in[22 : 22+length])

		return OperatorPacket{header, packets}, in[22+length:]
	case '1': // 11 bit == subpacket count
		count, err := strconv.ParseUint(in[7:18], 2, 16)
		check(err)
		subSequence := in[18:]
		for i := 0; i < int(count); i++ {
			var innerPacket Packet
			innerPacket, subSequence = newPacket(subSequence)
			packets = append(packets, innerPacket)
		}

		return OperatorPacket{header, packets}, subSequence
	default:
		panic("invalid length bit")
	}
}

func MessageResults(packets []Packet, remainder string) (results []*big.Int) {
	for _, p := range packets {
		results = append(results, p.GetResult())
	}

	return
}

func NewMessage(sequence string) (packets []Packet, remainder string) {
	for len(sequence) > 0 {
		if zeroRightPad.MatchString(sequence) {
			return packets, sequence
		}

		var packet Packet
		packet, sequence = newPacket(sequence)

		packets = append(packets, packet)
	}

	return
}

func ToBinary(hex string) (bin string) {
	for _, c := range hex {
		n, err := strconv.ParseUint(string(c), 16, 32)
		check(err)
		bin += fmt.Sprintf("%04b", n)
	}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
