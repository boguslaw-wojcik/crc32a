package crc32a_test

import (
	"fmt"
	"hash/crc32"
	"testing"

	"github.com/boguslaw-wojcik/crc32a"
)

// TestChecksum checks returned checksum for validity by comparing results from PHP implementation.
func TestChecksum(t *testing.T) {
	tests := map[uint32]string{
		404326908:  "123456789",
		2927368811: "Hello World!",
		3750220813: "d41d8cd98f00b204e9800998ecf8427e",
		3035062612: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		3879939681: "I have spread my dreams under your feet; Tread softly because you tread on my dreams.",
	}
	for sum, data := range tests {
		t.Run("", func(t *testing.T) {
			if got := crc32a.Checksum([]byte(data)); got != sum {
				t.Errorf("Checksum() = %v, want %v", got, sum)
			}
		})
	}
}

// TestChecksumHex checks returned checksum hex for validity by comparing results from PHP implementation.
func TestChecksumHex(t *testing.T) {
	tests := map[string]string{
		"181989fc": "123456789",
		"ae7c1a6b": "Hello World!",
		"df87d40d": "d41d8cd98f00b204e9800998ecf8427e",
		"b4e76154": "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		"e7432e61": "I have spread my dreams under your feet; Tread softly because you tread on my dreams.",
	}
	for hex, data := range tests {
		t.Run(hex, func(t *testing.T) {
			if got := crc32a.ChecksumHex([]byte(data)); got != hex {
				t.Errorf("ChecksumHex() = %v, want %v", got, hex)
			}
		})
	}

	checksum = crc32a.Checksum(data)
	fmt.Println(fmt.Sprintf("%x", checksum))
	checksum = crc32.ChecksumIEEE(data)
	fmt.Println(fmt.Sprintf("%x", checksum))
	checksum = crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli))
	fmt.Println(fmt.Sprintf("%x", checksum))
	checksum = crc32.Checksum(data, crc32.MakeTable(crc32.Koopman))
	fmt.Println(fmt.Sprintf("%x", checksum))
}

// variables declared and initialized for benchmark.
var (
	checksum uint32
	data     = []byte("123456789")
)

// BenchmarkCRC32A benchmarks implementation of CRC32A (ITU I.363.5) checksum.
func BenchmarkCRC32A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum = crc32a.Checksum(data)
	}
}

// BenchmarkCRC32B benchmarks standard library implementation of CRC32B (ITU V.42 / IEEE 802.3) checksum.
func BenchmarkCRC32B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum = crc32.ChecksumIEEE(data)
	}
}

// BenchmarkCRC32C benchmarks standard library implementation of CRC32C (Castagnoli) checksum.
func BenchmarkCRC32C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum = crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli))
	}
}

// BenchmarkCRC32K benchmarks standard library implementation of CRC32K (Koopman) checksum.
func BenchmarkCRC32K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum = crc32.Checksum(data, crc32.MakeTable(crc32.Koopman))
	}
}
