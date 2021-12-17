package packet

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_toBinary(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantBin string
	}{
		{
			"test F",
			args{"F"},
			"1111",
		},
		{
			"test F7",
			args{"F7"},
			"11110111",
		},
		{
			"test F73",
			args{"F73"},
			"111101110011",
		},
		{
			"test D2FE28",
			args{"D2FE28"},
			"110100101111111000101000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBin := ToBinary(tt.args.hex); gotBin != tt.wantBin {
				t.Errorf("toBinary() = %v, want %v", gotBin, tt.wantBin)
			}
		})
	}
}

func Test_newLiteralValuePacket(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		args       args
		wantPacket LiteralValuePacket
		wantOut    string
	}{
		{
			"test 1",
			args{"D2FE28"},
			LiteralValuePacket{PacketHeader{6, Value}, big.NewInt(2021)},
			"000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPacket, gotOut := newLiteralValuePacket(ToBinary(tt.args.in))
			if !reflect.DeepEqual(gotPacket, tt.wantPacket) {
				t.Errorf("newLiteralValuePacket() gotPacket = %v, want %v", gotPacket, tt.wantPacket)
			}
			if gotOut != tt.wantOut {
				t.Errorf("newLiteralValuePacket() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_newOperatorPacket(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		args       args
		wantPacket OperatorPacket
		wantResult *big.Int
		wantOut    string
	}{
		{
			"test Multiple operations (Sum, Mult, Equal)",
			args{"9C0141080250320F1802104A08"},
			OperatorPacket{PacketHeader{4, Equal}, []Packet{
				OperatorPacket{PacketHeader{2, Sum}, []Packet{
					LiteralValuePacket{PacketHeader{2, Value}, big.NewInt(1)},
					LiteralValuePacket{PacketHeader{4, Value}, big.NewInt(3)},
				}},
				OperatorPacket{PacketHeader{6, Mult}, []Packet{
					LiteralValuePacket{PacketHeader{0, Value}, big.NewInt(2)},
					LiteralValuePacket{PacketHeader{2, Value}, big.NewInt(2)},
				}},
			}},
			big.NewInt(1),
			"00",
		},
		{
			"test Equal",
			args{"9C005AC2F8F0"},
			OperatorPacket{PacketHeader{4, Equal}, []Packet{
				LiteralValuePacket{PacketHeader{5, Value}, big.NewInt(5)},
				LiteralValuePacket{PacketHeader{7, Value}, big.NewInt(15)},
			}},
			big.NewInt(0),
			"0000",
		},
		{
			"test Greater",
			args{"F600BC2D8F"},
			OperatorPacket{PacketHeader{7, Greater}, []Packet{
				LiteralValuePacket{PacketHeader{7, Value}, big.NewInt(5)},
				LiteralValuePacket{PacketHeader{5, Value}, big.NewInt(15)},
			}},
			big.NewInt(0),
			"",
		},
		{
			"test Less",
			args{"D8005AC2A8F0"},
			OperatorPacket{PacketHeader{6, Less}, []Packet{
				LiteralValuePacket{PacketHeader{5, Value}, big.NewInt(5)},
				LiteralValuePacket{PacketHeader{2, Value}, big.NewInt(15)},
			}},
			big.NewInt(1),
			"0000",
		},
		{
			"test length + Less",
			args{"38006F45291200"},
			OperatorPacket{PacketHeader{1, Less}, []Packet{
				LiteralValuePacket{PacketHeader{6, Value}, big.NewInt(10)},
				LiteralValuePacket{PacketHeader{2, Value}, big.NewInt(20)},
			}},
			big.NewInt(1),
			"0000000",
		},
		{
			"test count + Max",
			args{"EE00D40C823060"},
			OperatorPacket{PacketHeader{7, Max}, []Packet{
				LiteralValuePacket{PacketHeader{2, Value}, big.NewInt(1)},
				LiteralValuePacket{PacketHeader{4, Value}, big.NewInt(2)},
				LiteralValuePacket{PacketHeader{1, Value}, big.NewInt(3)},
			}},
			big.NewInt(3),
			"00000",
		},
		{
			"test Sum",
			args{"C200B40A82"},
			OperatorPacket{PacketHeader{6, Sum}, []Packet{
				LiteralValuePacket{PacketHeader{6, Value}, big.NewInt(1)},
				LiteralValuePacket{PacketHeader{2, Value}, big.NewInt(2)},
			}},
			big.NewInt(3),
			"",
		},
		{
			"test Mult",
			args{"04005AC33890"},
			OperatorPacket{PacketHeader{0, Mult}, []Packet{
				LiteralValuePacket{PacketHeader{5, Value}, big.NewInt(6)},
				LiteralValuePacket{PacketHeader{3, Value}, big.NewInt(9)},
			}},
			big.NewInt(54),
			"0000",
		},
		{
			"test Min",
			args{"880086C3E88112"},
			OperatorPacket{PacketHeader{4, Min}, []Packet{
				LiteralValuePacket{PacketHeader{5, Value}, big.NewInt(7)},
				LiteralValuePacket{PacketHeader{6, Value}, big.NewInt(8)},
				LiteralValuePacket{PacketHeader{0, Value}, big.NewInt(9)},
			}},
			big.NewInt(7),
			"0",
		},
		{
			"test Max",
			args{"CE00C43D881120"},
			OperatorPacket{PacketHeader{6, Max}, []Packet{
				LiteralValuePacket{PacketHeader{0, Value}, big.NewInt(7)},
				LiteralValuePacket{PacketHeader{5, Value}, big.NewInt(8)},
				LiteralValuePacket{PacketHeader{0, Value}, big.NewInt(9)},
			}},
			big.NewInt(9),
			"00000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPacket, gotOut := newOperatorPacket(ToBinary(tt.args.in))
			if !reflect.DeepEqual(gotPacket, tt.wantPacket) {
				t.Errorf("newOperatorPacket() gotPacket = %v, want %v", gotPacket, tt.wantPacket)
			}
			if gotPacket.GetResult().Cmp(tt.wantResult) != 0 {
				t.Errorf("newOperatorPacket() packet result = %v, want %v", gotPacket.GetResult(), tt.wantResult)
			}
			if gotOut != tt.wantOut {
				t.Errorf("newOperatorPacket() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
