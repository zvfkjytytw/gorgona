package tests

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func QuizTest() error{
	// Set up ChromeDriver options
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// Configure Chrome-specific capabilities
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		},
	}
	caps.AddChrome(chromeCaps)

	// Start ChromeDriver service
	service, err := selenium.NewChromeDriverService("chromedriver", 4444, opts...)
	if err != nil {
		// log.Fatal("error starting ChromeDriver service:", err)
		return fmt.Errorf("error starting ChromeDriver service: %v", err)
	}
	defer service.Stop()

	// Connect to ChromeDriver instance
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		// log.Fatal("error creating WebDriver:", err)
		return fmt.Errorf("error creating WebDriver: %v", err)
	}
	defer driver.Quit()

	err = driver.Get("http://185.104.249.14/")
	if err != nil {
	// 	log.Fatal("error navigating to website:", err)
		return fmt.Errorf("error navigating to website: %v", err)
	}	

	startTest, err := driver.FindElement(selenium.ByXPATH, "//button")
	if err != nil {
		// log.Fatalf("error finding start button: %v", err)
		return fmt.Errorf("error finding start button: %v", err)
	}
	startTest.Click()
	time.Sleep(time.Second)

	var finish bool
	for !finish {
		wrong, err := driver.FindElements(selenium.ByXPATH, "//h3[@class='failure']")
		if err == nil && len(wrong) > 0 {
			finish = true
			// log.Fatal("test failed. Wrong answer")
			// continue
			return fmt.Errorf("test failed. Wrong answer")
		}

		success, err := driver.FindElements(selenium.ByXPATH, "//h3[@class='success']")
		if err == nil && len(success) > 0 {
			finish = true
			// log.Print("test complite")
			// continue
			return nil
		}

		elems, err := driver.FindElements(selenium.ByXPATH, "//input[@type='radio']")
		if err != nil {
			// log.Fatalf("test failed. No elements: %v", err)
			return fmt.Errorf("test failed. no elements: %v", err)
		}
		index, longest := 0, 0
		for i, elem := range elems {
			value, err := elem.GetAttribute("value")
			if err != nil {
				value = "-"
			}
			if len(value) > longest {
				index = i
				longest = len(value)
			}

			// log.Printf("Elem: %d\t\tValue: %v\t\tLong: %d\n", i, value, len(value))

		}
		// log.Printf("Select answer: %d\t\t%v\n", index, elems[index])
		elems[index].Click()

		submit, err := driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
		if err != nil {
			// log.Fatalf("error finding next button: %v", err)
			return fmt.Errorf("error finding next button: %v", err)
		}
		submit.Click()
		time.Sleep(time.Second)
	}

	return nil
}
