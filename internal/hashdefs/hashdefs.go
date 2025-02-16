package hashdefs

import (
	"github.com/snksoft/crc"
	"hash"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"strings"
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
func MapHashes() map[string]func() hash.Hash {
	return map[string]func() hash.Hash{
		"ADLER32": func() hash.Hash { return adler32.New() },
		"CRC32": func() hash.Hash { return crc32.NewIEEE() },
		"CRC64": func() hash.Hash { return crc64.New(crc64.MakeTable(crc64.ISO)) },
		"FNV32": func() hash.Hash { return fnv.New32() },
		"FNV32A": func() hash.Hash { return fnv.New32a() },
		"FNV64": func() hash.Hash { return fnv.New64() },
		"FNV64A": func() hash.Hash { return fnv.New64a() },
		"FNV128": func() hash.Hash { return fnv.New128() },
		"FNV128A": func() hash.Hash { return fnv.New128a() },
		"MD5": md5.New, 
		"SHA1": sha1.New, 
		"SHA224": sha256.New224, 
		"SHA256": sha256.New,
		"SHA384": sha512.New384,
		"SHA512_224": sha512.New512_224, 
		"SHA512_256": sha512.New512_256,
		"SHA512": sha512.New,
	}
}

func GetHash(name string) hash.Hash {
	crcmap := MapCrcs()
	p := crcmap[strings.ToUpper(name)]

	if p != nil {
		return crc.NewHash(p)
	}

	hashmap := MapHashes()
	f := hashmap[strings.ToUpper(name)]

	if f != nil {
		return f()
	}

	return nil
}

