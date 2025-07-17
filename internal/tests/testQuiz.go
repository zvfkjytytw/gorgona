package tests

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	quizURL          string = "http://185.104.249.14/"
	chromeDriverPort int    = 4444
)

func TestQuiz() error {
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
	// service, err := selenium.NewChromeDriverService("chromedriver", 4444, opts...)\
	service, err := selenium.NewChromeDriverService("chromedriver", chromeDriverPort, opts...)
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

	// Connect to test server
	err = driver.Get(quizURL)
	if err != nil {
		return fmt.Errorf("error navigating to website: %v", err)
	}

	// Search and press the start button
	startTest, err := driver.FindElement(selenium.ByXPATH, "//button")
	if err != nil {
		return fmt.Errorf("error finding start button: %v", err)
	}
	startTest.Click()
	time.Sleep(time.Second)

	var finish bool
	for !finish {
		// Check next page for the fault case
		wrong, err := driver.FindElements(selenium.ByXPATH, "//h3[@class='failure']")
		if err == nil && len(wrong) > 0 {
			finish = true
			return fmt.Errorf("test failed. Wrong answer")
		}

		// Check next page for the success case
		success, err := driver.FindElements(selenium.ByXPATH, "//h3[@class='success']")
		if err == nil && len(success) > 0 {
			finish = true
			return nil
		}

		// Go through all the answers and choose the longest one.
		elems, err := driver.FindElements(selenium.ByXPATH, "//input[@type='radio']")
		if err != nil {
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
		}
		elems[index].Click()

		// Search and press the submit button
		submit, err := driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
		if err != nil {
			return fmt.Errorf("error finding next button: %v", err)
		}
		submit.Click()
		time.Sleep(time.Second)
	}

	return nil
}
