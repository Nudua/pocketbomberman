package gbutils

import "fmt"

//http://gbdev.gg8.se/wiki/articles/The_Cartridge_Header

//014E-014F - Global Checksum
//The global checksum for a gbrom is a 16bit summation of all the bytes in the rom minus the two checksum bytes located at 0x14E-0x14F, it is ignored by a real game boy.

//FixGlobalChecksum fixes the global checksum for 'Game Boy' and 'Game Boy Color' games
func FixGlobalChecksum(romData []byte) {
	checkSumLocationStart := 0x14E
	checkSumLocationEnd := 0x14F

	checksum := 0

	for i := 0; i < len(romData); i++ {
		if i == checkSumLocationStart || i == checkSumLocationEnd {
			continue
		}
		checksum += int(romData[i])
	}

	checksum = checksum & 65535

	upper := byte(checksum >> 8)
	lower := byte(checksum & 0xFF)

	romData[checkSumLocationStart] = upper
	romData[checkSumLocationEnd] = lower

	fmt.Printf("Global checksum updated: 0x%x%x\n", upper, lower)
}
