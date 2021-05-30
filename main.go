package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	showForm()
	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			iniciarMonitorameno()
			fmt.Println("monitoring started.")
		case 2:
			fmt.Println("Going out...")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		default:
			fmt.Println("Option invalid")
			os.Exit(-1)
		}
	}
}

var nomeLido string

func showForm() {
	fmt.Print("Your Nome: ")
	fmt.Scan(&nomeLido)
	//command to clean terminal
	fmt.Print("\033[H\033[2J")
}

func showMenu() {
	fmt.Println("Hello,", nomeLido, "enter one of the options below.")
	fmt.Println("1 - Monitor website")
	fmt.Println("2 - Exit")
}

func readCommand() int {
	var command int

	fmt.Printf("- Option: ")
	fmt.Scan(&command)
	fmt.Println("The chosen command was:", command)
	fmt.Println("")
	return command
}

func siteMonitor() string {
	var site string

	fmt.Print("Site to monitor: ", site)
	fmt.Scan(&site)
	//command to clean terminal
	fmt.Print("\033[H\033[2J")
	return site
}

func testSite(site string) {
	response, err := http.Get(site)
	dt := time.Now()

	if err != nil {
		fmt.Println("Site is not found", err)
	}

	if response.StatusCode == http.StatusOK {
		fmt.Println(dt.Format("01-02-2006 15:04:05 Monday"))
		fmt.Print("Website: ", site, " - Online \n\n")

	} else {
		fmt.Print("The website: ", site, " has a problem. Status Code:", response.StatusCode, "\n\n")
		os.Exit(0)
	}
}

func iniciarMonitorameno() {
	site := siteMonitor()
	for {
		testSite(site)
		time.Sleep(5 * time.Second)
	}
}
