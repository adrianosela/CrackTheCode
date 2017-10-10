package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

var goroutines = 10
var maxNum = 5000

//Crack tries all the numspace from 0 to the max
func Crack(c *cli.Context) {
	ch := make(chan bool)

	attemptsPerGoRtn := maxNum / goroutines

	for i := 0; i < goroutines; i++ {
		from := i * attemptsPerGoRtn
		until := (i + 1) * attemptsPerGoRtn
		go tryToCrackTheCode(i, from, until, ch)
	}

	quit := make(chan bool)
	go startTimer(quit)

	select {
	case <-ch:
		fmt.Println("..................TERMINATING..................")
		return
	case <-quit:
		fmt.Println("..................TERMINATING..................")
		return
	}

}

func tryToCrackTheCode(goroutineNum, from, until int, c chan bool) {
	cli := http.DefaultClient

	for i := from; i < until; i++ {

		fmt.Printf("GoRoutine #%d is trying %d\n", goroutineNum, i)

		req, err := http.NewRequest("GET", "http://localhost:8080/crack/"+strconv.Itoa(i), nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := cli.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("GoRoutine #%d cracked the code. Code: %d\n", goroutineNum, i)
			c <- true
		}

		time.Sleep(time.Millisecond * 100)

	}

	fmt.Printf("GoRoutine #%d determined the code is not from %d to %d ", goroutineNum, from, until)
}

func startTimer(c chan bool) {
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second * 10)
		fmt.Printf("INFO: %d seconds have elapsed!\n", i*10)
	}
	fmt.Println("TIMED OUT")
	c <- true
	return
}

//CrackTheCode sdasd
func CrackTheCode(c *cli.Context) error {
	num := c.Int("num")
	tryAll := c.Bool("all")
	//the c.Int() function returns 0 if the flag value is not found
	if num == 0 && !tryAll {
		cli.ShowCommandHelp(c, "crack")
		return errors.New("[ERROR] -num is a mandatory flag if -all is not set")
	}
	if num > maxNum {
		return fmt.Errorf("[ERROR] The maximum allowed is %d", maxNum)
	}
	//if the try all flag is set
	if tryAll {
		Crack(c)
		return nil
	}

	req, err := http.NewRequest("GET", "http://localhost:8080/crack/"+strconv.Itoa(num), nil)
	if err != nil {
		log.Fatal(err)
	}
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}

//GenerateNewCode generates a new code
func GenerateNewCode(c *cli.Context) error {
	req, err := http.NewRequest("GET", "http://localhost:8080/generate/"+strconv.Itoa(maxNum), nil)
	if err != nil {
		log.Fatal(err)
	}
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}

func Cheat(c *cli.Context) error {
	req, err := http.NewRequest("GET", "http://localhost:8080/cheat", nil)
	if err != nil {
		log.Fatal(err)
	}
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}

func Tries(c *cli.Context) error {
	req, err := http.NewRequest("GET", "http://localhost:8080/tries", nil)
	if err != nil {
		log.Fatal(err)
	}
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
