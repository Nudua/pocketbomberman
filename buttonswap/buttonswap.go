package buttonswap

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/Nudua/pocketbomberman/bytesearch"
)

//IsKnownRom Check if the rom is one of the three known versions
func IsKnownRom(romData []byte) string {
	hash := md5.New()
	hash.Write(romData)

	stringHash := fmt.Sprintf("%x", hash.Sum(nil))

	stringHash = strings.ToUpper(stringHash)

	knownRoms := map[string]string{
		"3050CDF055B1D036F85C9C629B03FBF2": "Pocket Bomberman (Europe) - Game Boy",
		"50971F6CBB3E4395CB66213D7F096B46": "Pocket Bomberman (Japan) - Game Boy",
		"2F6B6379F8C7CE5D66A198162F345EAA": "Pocket Bomberman (USA, Europe) - Game Boy Color",
	}

	name, exists := knownRoms[stringHash]

	if exists {
		return name
	}

	return ""
}

//ButtonSwapPatch A-B Button swap patch for Pocket Bomberman
func Patch(romData []byte) bool {

	jump := []byte{0xFA, 0x83, 0xFF, 0xCB, 0x47}
	//0xFA, 0x83, 0xFF, //ld a, (FF83) load current button state
	//0xCB, 0x47        //bit 1,a check if the first bit (0x20) is set (Jump) - B

	jumpPrev := []byte{0xFA, 0x82, 0xFF, 0xCB, 0x47}
	//0xFA, 0x82, 0xFF,   //ld a, (FF83) load previous button state
	//0xCB, 0x47          //bit 1,a check if the first bit (0x02) is set (Jump) - B

	bomb := []byte{0xFA, 0x83, 0xFF, 0xCB, 0x4F}
	//0xFA, 0x83, 0xFF,   //ld a, (FF83) load previous button state
	//0xCB, 0x4F          //bit 0,a check if the first bit (0x01) is set (Bomb) - A

	bombPrev := []byte{0xFA, 0x82, 0xFF, 0xCB, 0x4F}
	//0xFA, 0x82, 0xFF,   //ld a, (FF83) load previous button state
	//0xCB, 0x4F          //bit 0,a check if the first bit (0x01) is set (Bomb) - A

	jumpIndexes := bytesearch.IndexOfAll(romData, jump)

	jumpPrevIndexes := bytesearch.IndexOfAll(romData, jumpPrev)

	//Merge them because they replace the same value
	jumpIndexes = append(jumpIndexes, jumpPrevIndexes...)

	//Swap B (bit 0) to A (bit 1)
	bombIndexes := bytesearch.IndexOfAll(romData, bomb)

	bombPrevIndexes := bytesearch.IndexOfAll(romData, bombPrev)

	//Merge them because they replace the same value
	bombIndexes = append(bombIndexes, bombPrevIndexes...)

	//Check if any indexes were actually found
	if len(jumpIndexes) == 0 && len(bombIndexes) == 0 {
		return false
	}

	var bitOne byte = 0x47
	var bitZero byte = 0x4F

	fmt.Print("Patching...")

	for _, index := range jumpIndexes {
		romData[index+4] = bitZero

		fmt.Print("...")
	}

	for _, index := range bombIndexes {
		romData[index+4] = bitOne

		fmt.Print("...")
	}

	fmt.Print("\n")

	return true
}
