package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Nudua/pocketbomberman/buttonswap"
	"github.com/Nudua/pocketbomberman/gbutils"
)

func main() {

	//No rom file were supplied, exit and display usage.
	if len(os.Args) < 2 {
		printUsage()
		keepConsoleWindowOpen()
		return
	}

	fileName := os.Args[1]

	//Probably not needed?
	if len(fileName) == 0 {

		printUsage()
		keepConsoleWindowOpen()
		return
	}

	//get all bytes from the rom
	romData, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Unable to open the file '%v' please check the filepath...\n", fileName)
		keepConsoleWindowOpen()
		return
	}

	//Check if file is valid
	gameVersion := buttonswap.IsKnownRom(romData)

	//Possibly unknown rom dump or invalid file
	if len(gameVersion) == 0 {
		fmt.Println("Unknown rom file, patch anyway? y/yes n/no")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		text = strings.ToLower(text)

		text = strings.TrimSpace(text)

		if text != "yes" && text != "y" {
			fmt.Println("Patching aborted...")
			keepConsoleWindowOpen()
			return
		}

	} else {
		fmt.Printf("Valid rom '%v' was found.\n\n", gameVersion)
	}

	if buttonswap.Patch(romData) {

		//Recalculate the global rom checksum, not really needed, but nice for bgb
		gbutils.FixGlobalChecksum(romData)

		patchedFileName := getPatchedFileName(fileName)

		fullPatchedFileName := filepath.Join(filepath.Dir(fileName), patchedFileName)

		err = ioutil.WriteFile(fullPatchedFileName, romData, 0644)

		if err != nil {
			fmt.Println("There was an issue while writing the patched rom to the hdd...")
			keepConsoleWindowOpen()
			return
		}

		fmt.Printf("\nPatched rom was successfully created at \"%v\"\n", fullPatchedFileName)
		keepConsoleWindowOpen()

	} else {
		fmt.Println("Nothing to patch, incorrect file?")
		keepConsoleWindowOpen()
	}
}

func getPatchedFileName(fileName string) string {
	short := filepath.Base(fileName)

	ext := filepath.Ext(fileName)

	newFileName := short[:strings.Index(short, ext)]

	newFileName = newFileName + "(Patched)" + ext

	return newFileName
}

func printUsage() {
	fmt.Print("Pocket Bomberman A+B Button Swap 1.0 by Nudua\n\n")

	if runtime.GOOS == "windows" {
		fmt.Println("Usage: buttonswap.exe \"C:\\path\\to\\roms\\Pocket Bomberman.gbc\"")

		fmt.Println("\nTip: You can just drag and drop the rom over buttonswap.exe")

	} else {
		fmt.Println("Usage: ./buttonswap \"/Home/User/romfilename.gbc\"")
		fmt.Println("\nNote: Make sure the program has execute permissions (744) 'chmod buttonswap 744'")
	}
}

func keepConsoleWindowOpen() {
	fmt.Println("Done...")
	//Only do this on windows because the window closes automatically otherwise
	if runtime.GOOS == "windows" {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}
}
