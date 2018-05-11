package main

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")

	work := func() {
		drone.TakeOff()

		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func GetMPlayerInput() (io.WriteCloser, error) {
	mPlayer := exec.Command("mplayer", "-vo", "x11", "-fps", "30", "-")
	mPlayer.Start()
	return mPlayer.StdinPipe()
}

func WriteCameraOutputToMplayer(droneVideoOutput chan []byte, mPlayerIn io.WriteCloser) {
	for frame := range droneVideoOutput {
		if _, err := mPlayerIn.Write(frame); err != nil {
			fmt.Printf("failed to write frame to movie player: %v\n", err)
		}
	}
}
