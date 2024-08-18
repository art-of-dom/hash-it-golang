package hashdefs

import (
	"github.com/snksoft/crc"
)

// DataType - Custom type to determine incoming data type
type DataType int

const (
	Ascii DataType = iota
	Hex
	File
	Stdin
	Bytes
)

var (
	CRC24 = &crc.Parameters{Width: 24, Polynomial: 0x1864CFB, Init: 0xB704CE, ReflectIn: false, ReflectOut: false, FinalXor: 0x000000}
)

func MapCrcs() map[string]*crc.Parameters {
	return map[string]*crc.Parameters{
		"X25":        crc.X25,
		"CCITT":      crc.CCITT,
		"CRC16":      crc.CRC16,
		"XMODEM":     crc.XMODEM,
		"XMODEM2":    crc.XMODEM2,
		"CRC24":      CRC24,
		"CRC32":      crc.CRC32,
		"IEEE":       crc.IEEE,
		"CASTAGNOLI": crc.Castagnoli,
		"CRC32C":     crc.CRC32C,
		"KOOPMAN":    crc.Koopman,
		"CRC64ISO":   crc.CRC64ISO,
		"CRC64ECMA":  crc.CRC64ECMA,
	}
}
