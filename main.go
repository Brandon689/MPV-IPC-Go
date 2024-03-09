package main

import (
	"fmt"
	"os"
	"time"

	npipe "gopkg.in/natefinch/npipe.v2"
)

func main() {
	fmt.Println("ok")
	conn, err := npipe.Dial(`\\.\pipe\mpv-pipe`)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	//fmt.Fprintf(conn, `{"command": ["get_property", "playback-time"]}`+"\n")
	//fmt.Fprintf(conn, `{"command": ["set_property", "volume", 25]}`+"\n")
	//
	//status, err := bufio.NewReader(conn).ReadString('\n')
	//if err != nil {
	//	fmt.Println("Failed to read:", err)
	//	os.Exit(1)
	//}
	////"C:\Users\Brandon\Videos\brighter.mp4"
	//fmt.Println("Status:", status)
	time.Sleep(3 * time.Second)
	// Replace with the path to your video file
	videoFilePath := `C:\\Users\\Brandon\\Videos\\brighter.mp4`
	videoFilePath2 := `C:\\Users\\Brandon\\Videos\\some.mp4`

	// Send the loadfile command
	fmt.Fprintf(conn, `{"command": ["loadfile", "%s"]}`+"\n", videoFilePath)

	//startTime := 60 // Start at 60 seconds

	// Send the loadfile command with the start time
	fmt.Fprintf(conn, `{"command": ["loadfile", "%s", "replace", "start=03:11:10"]}`+"\n", videoFilePath)

	time.Sleep(3 * time.Second)

	fmt.Fprintf(conn, `{"command": ["loadfile", "%s", "replace", "start=03:11:10"]}`+"\n", videoFilePath2)

}
