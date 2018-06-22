package web

import (
	"fmt"
	"os/exec"
)

// runtimes
const LINUX_RT = "linux"
const WINDOWS_RT = "windows"
const MAC_RT = "darwin"

// commands
const LINUX_BROWSER = "xdg-open"
const WINDOWS_BROWSER = "rundll32"
const WINDOWS_BROWSER_OPTS = "url.dll,FileProtocolHandler"
const MAC_BROWSER = "open"

func Init() {
	// StartBrowser(url) TODO: get browser to start when command is called
}

func ServeVideoWS(vidStream <-chan []byte, w http.ResponseWriter, r *http.Request) string {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	for frame := range vidStream {
		err := conn.WriteMessage(frame)
		handle(err)
	}
}

func StartBrowser(url string) {
	switch runtime.GOOS {
	case LINUX_RT:
		err := exec.Command(LINUX_BROWSER, url).Start()
	case WINDOWS_RT:
		err := exec.Command(WINDOWS_BROWSER, WINDOWS_BROWSER_OPTS, url).Start()
	case MAC_RT:
		err := exec.Command(MAC_BROWSER, url).Start()
	default:
		err = fmt.Errorf("could not detect operating system when trying to start browser")
	}
}

func handle(err error) {
	fmt.Println(err)
}
