package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TwinProduction/go-color"
	"github.com/martinlindhe/notify"
)

func main() {
	for {
		showMenu()
		comando := readCommand()

		switch comando {
		case 1:
			startMonitoring()
			fmt.Println("monitoring started.")
		case 2:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Going out...")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		default:
			fmt.Print("\033[H\033[2J")
			print(color.Red + " --- ENTER A VALID OPTION --- ")
			showMenu()
		}
	}
}

func readCommand() int {
	var command int

	fmt.Print(color.Ize(color.Green, "- Option: "))
	fmt.Scan(&command)
	fmt.Println("The chosen command was:", command)
	fmt.Println("")
	return command
}

func siteMonitor() string {
	fmt.Print("\033[H\033[2J")
	var site string

	print(color.Green + "Site to monitor: " + site + color.Reset)

	fmt.Scan(&site)
	//command to clean terminal
	fmt.Print("\033[H\033[2J")
	return site
}

func requestSite(site string) {
	response, err := http.Get(site)
	dt := time.Now()

	if err != nil {
		fmt.Println("Site is not found", err)
	}

	switch response.StatusCode {
	case http.StatusOK:
		fmt.Println(dt.Format("01-02-2006 15:04:05 Monday"))
		print(color.Green + "Website: " + site + " - Online \n\n " + color.Reset)
	default:
		print(color.Red+"The website: "+site+" has a problem. Status Code:", response.StatusCode, "\n\n"+color.Reset)
		notify.Alert("app name", "ALERT!", "Site failed: "+site, "path/to/icon.png")
		os.Exit(0)
	}
}

func startMonitoring() {
	var setTimeDuration int
	site := siteMonitor()

	showMenuTimer()
	interval := readCommand()

	switch interval {
	case 1:
		fmt.Print(color.Ize(color.Green, "Time interval: "))
		fmt.Scan(&setTimeDuration)
		for {
			requestSite(site)
			time.Sleep(time.Duration(setTimeDuration) * time.Hour)
		}
	case 2:
		fmt.Print(color.Ize(color.Green, "Time interval: "))
		fmt.Scan(&setTimeDuration)
		for {
			requestSite(site)
			time.Sleep(time.Duration(setTimeDuration) * time.Minute)
		}
	case 3:
		fmt.Print(color.Ize(color.Green, "Time interval: "))
		fmt.Scan(&setTimeDuration)
		for {
			requestSite(site)
			time.Sleep(time.Duration(setTimeDuration) * time.Second)
		}
	case 4:
		fmt.Print("\033[H\033[2J")
		fmt.Println("Going out...")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	default:
		fmt.Print("\033[H\033[2J")
		print(color.Red + " --- ENTER A VALID OPTION --- ")
		showMenuTimer()
	}
}

func showMenu() {
	fmt.Println("Enter one of the options below.")
	fmt.Println(color.Ize(color.Cyan, "1 - Monitor website"))
	fmt.Println(color.Ize(color.Red, "2 - Exit"))
}

func showMenuTimer() {
	fmt.Println("Enter one of the options below.")
	fmt.Print(color.Ize(color.Cyan, "1 - Hour | 2 - Minute | 3 - Second | 4 - Exit: \n\n"))
}
