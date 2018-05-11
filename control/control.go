package control

import (
	"fmt"

	term "github.com/nsf/termbox-go"
	"gobot.io/x/gobot/platforms/dji/tello"
)

const speed = 40

func InitControl(driver *tello.Driver) {
	err := term.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer term.Close()

	fmt.Println("Drone: control intitialized")
	for {
		event := term.PollEvent()
		switch event.Type {
		case term.EventKey:
			switch event.Ch {
			case 'w':
				term.Sync()
				driver.Forward(speed)
				fmt.Println("w Pressed")
			case 'a':
				term.Sync()
				driver.Left(speed)
				fmt.Println("a pressed")
			case 's':
				term.Sync()
				driver.Backward(speed)
				fmt.Println("s pressed")
			case 'd':
				term.Sync()
				driver.Right(speed)
				fmt.Println("d pressed")
			case 'q':
				term.Sync()
				driver.CounterClockwise(speed)
				fmt.Println("q pressed")
			case 'e':
				term.Sync()
				driver.Clockwise(speed)
				fmt.Println("e pressed")
			case 'r':
				term.Sync()
				driver.Up(speed)
				fmt.Println("r pressed")
			case 'f':
				term.Sync()
				driver.Down(speed)
				fmt.Println("f pressed")
			case 'l':
				term.Sync()
				driver.Land()
			case 't':
				term.Sync()
				driver.TakeOff()
				fmt.Println("t pressed")
			default:
				fmt.Println("hello fuck dog")
			}
		default:
			term.Sync()
		}

	}
}
