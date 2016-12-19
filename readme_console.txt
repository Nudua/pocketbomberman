------------------------------------
Pocket Bomberman A+B Button Swap 1.0
------------------------------------
This patch will swap A and B so that the 'A' button becomes Jump and 'B' places bombs, which is what most platform games use by default.
Works for the American/European Game Boy Color version and the European and Japanese Game Boy version.
This tool runs on Windows, Mac OS X (10.7 and above) and most Linux distributions.

 Not Patched        Patched
  B  -  A     ->    B  -  A    
 Jump  Bomb        Bomb  Jump

--------
Patching
--------
Windows (Windows XP or later)
	1) Drag and drop the GB/GBC version of Pocket Bomberman over "buttonswap.exe"
	2) The program will then create a new rom named "filename(patched).gbc" in the same directory as the rom is.
	
	Alt) Or use it from the command prompt Win+R: 
		buttonswap.exe "C:\path\to\roms\Pocket Bomberman.gbc"
		
Linux and Mac OS X (10.7 and above)
	1) Navigate to the appropriate subfolder for your system (i.e. Linux or Mac for OSX)
	2) From the terminal:
		./buttonswap "/Home/User/romfilename.gbc"
	3) The program will then create a new rom named "filename(patched).gbc" in the same directory as the rom is.
	
	*Note: use "./buttonswap32" for 32bit systems.
	*Note2: Make sure the program has execute permissions (744) 'chmod buttonswap 744'

Tested on Windows 10 Pro x64 and Ubuntu 16.04.1 x64

------
Source
------
This utility was written in Go (1.7.1) (https://golang.org/)

And the source code is available at: https://github.com/Nudua/pocketbomberman

------
Credits
-------
Made by Tor Ramstad aka Nudua, for any issues or questions feel free to contact via the info below.

Twitter:  https://twitter.com/TheRealNudua
Web:      http://nudua.com
Twitch:   http://twitch.tv/Nudua