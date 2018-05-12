package control

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

const speed = 40

func InitControl(driver *tello.Driver) {

	keyb := keyboard.NewDriver()

	fmt.Println("Drone: control intitialized")

	work := func() {

		keyb.On(keyboard.Key, func(data interface{}) {

			k := data.(keyboard.KeyEvent)

			switch k.Key {
			case keyboard.W:
				driver.Forward(speed)
				fmt.Println("w Pressed")
			case keyboard.A:
				driver.Left(speed)
				fmt.Println("a pressed")
			case keyboard.S:
				driver.Backward(speed)
				fmt.Println("s pressed")
			case keyboard.D:
				driver.Right(speed)
				fmt.Println("d pressed")
			case keyboard.Q:
				driver.CounterClockwise(speed)
				fmt.Println("q pressed")
			case keyboard.E:
				driver.Clockwise(speed)
				fmt.Println("e pressed")
			case keyboard.R:
				driver.Up(speed)
				fmt.Println("r pressed")
			case keyboard.F:
				driver.Down(speed)
				fmt.Println("f pressed")
			case keyboard.L:
				driver.Land()
			case keyboard.T:
				driver.TakeOff()
				fmt.Println("t pressed")
			default:
				fmt.Println("hello fuck dog")
			}

		})

	}

	robot := gobot.NewRobot("keyboardbot",
		[]gobot.Connection{},
		[]gobot.Device{keyb},
		work,
	)

	fmt.Println("keyboard initialized")

	robot.Start()
}
