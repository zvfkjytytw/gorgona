package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	gorgonaApp "github.com/zvfkjytytw/gorgona/internal/app"
)

func main() {
	var (
		// Maximum parallel flows
		maxGorutines int
		// Env variables of the maximum parallel flows
		envMaxGorutines string = "MAX_FLOWS"
 	)

	flag.IntVar(&maxGorutines, "f", 3, "Maximum parallel flows")
	flag.Parse()

	envF, ok := os.LookupEnv(envMaxGorutines)
	if ok {
		value, err := strconv.Atoi(envF)
		if err != nil {
			maxGorutines = value
		}
	}

	app, err := gorgonaApp.New(maxGorutines)
	if err != nil {
		fmt.Printf("test failed")
		os.Exit(1)
	}

	app.Run()
}

// import (
// 	"log"
// 	"time"

// 	"github.com/tebeka/selenium"
// 	"github.com/tebeka/selenium/chrome"
// )

// func main() {
// 	// Set up ChromeDriver options
// 	opts := []selenium.ServiceOption{}
// 	caps := selenium.Capabilities{
// 		"browserName": "chrome",
// 	}

// 	// Configure Chrome-specific capabilities
// 	chromeCaps := chrome.Capabilities{
// 		Args: []string{
// 			"--no-sandbox",
// 			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
// 		},
// 	}
// 	caps.AddChrome(chromeCaps)

// 	// Start ChromeDriver service
// 	service, err := selenium.NewChromeDriverService("chromedriver", 4444, opts...)
// 	if err != nil {
// 		log.Fatal("error starting ChromeDriver service:", err)
// 	}
// 	defer service.Stop()

// 	// Connect to ChromeDriver instance
// 	driver, err := selenium.NewRemote(caps, "")
// 	if err != nil {
// 		log.Fatal("error creating WebDriver:", err)
// 	}
// 	defer driver.Quit()

// 	log.Println("WebDriver initialized successfully")

// 	err = driver.Get("http://185.104.249.14/")
// 	if err != nil {
// 		log.Fatal("error navigating to website:", err)
// 	}

// 	startTest, err := driver.FindElement(selenium.ByXPATH, "//button")
// 	if err != nil {
// 		log.Fatalf("error finding start button: %v", err)
// 	}
// 	startTest.Click()
// 	time.Sleep(time.Second)

// 	var finish bool
// 	for !finish {
// 		wrong, err := driver.FindElements(selenium.ByXPATH, "//h3[@class='failure']")
// 		if err == nil && len(wrong) > 0 {
// 			finish = true
// 			log.Fatal("test failed. Wrong answer")
// 			continue
// 		}

// 		success, err := driver.FindElements(selenium.ByXPATH, "//h3[@class='success']")
// 		if err == nil && len(success) > 0 {
// 			finish = true
// 			log.Print("test complite")
// 			continue
// 		}

// 		elems, err := driver.FindElements(selenium.ByXPATH, "//input[@type='radio']")
// 		if err != nil {
// 			log.Fatalf("test failed. No elements: %v", err)
// 		}
// 		index, longest := 0, 0
// 		for i, elem := range elems {
// 			value, err := elem.GetAttribute("value")
// 			if err != nil {
// 				value = "-"
// 			}
// 			if len(value) > longest {
// 				index = i
// 				longest = len(value)
// 			}

// 			// log.Printf("Elem: %d\t\tValue: %v\t\tLong: %d\n", i, value, len(value))

// 		}
// 		// log.Printf("Select answer: %d\t\t%v\n", index, elems[index])
// 		elems[index].Click()

// 		submit, err := driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
// 		if err != nil {
// 			log.Fatalf("Error finding element by XPath: %v", err)
// 		}
// 		submit.Click()
// 		time.Sleep(time.Second)
// 	}

// 	// Wait for page to load
// 	// time.Sleep(2 * time.Second)
// }
