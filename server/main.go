package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var defaultMax = 5000

type Code struct {
	num   int
	tries uint
	lck   *sync.Mutex
}

func main() {

	c := getCode()

	router := mux.NewRouter()
	router.HandleFunc("/crack/{num}", c.crackTheCode)
	router.HandleFunc("/generate/{max}", c.generateNewCode)
	router.HandleFunc("/cheat", c.cheatAndSee)
	router.HandleFunc("/tries", c.countTries)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}

}

func getCode() Code {
	//new lock
	var mutex = &sync.Mutex{}
	//seed source and generate random number 0 - defaultMax
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	number := rand.Intn(defaultMax)
	//populate the Code struct
	c := Code{
		lck:   mutex,
		num:   number,
		tries: 0,
	}
	return c
}

func (c *Code) generateNewCode(w http.ResponseWriter, r *http.Request) {
	//grab the maximum number from the request
	vars := mux.Vars(r)
	maxStr := vars["max"]
	max, err := strconv.Atoi(maxStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Max number must be an integer")
		return
	}
	//grab the lock and defer its unlock
	c.lck.Lock()
	defer c.lck.Unlock()
	//seed a source with time and generate a num from 0 to max
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	//assign the number, reset the tries counter and return a 200
	c.num = rand.Intn(max)
	c.tries = 0
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS: New code from 0 to %d has been set", max)
	return
}

func (c *Code) crackTheCode(w http.ResponseWriter, r *http.Request) {
	//grab the attempted number from the request
	vars := mux.Vars(r)
	numberStr := vars["num"]
	num, err := strconv.Atoi(numberStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: Number must be an integer")
		return
	}
	//grab the lock
	c.lck.Lock()
	defer c.lck.Unlock()
	//increase tries counter and check number
	c.tries++
	if num != c.num {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "FAIL: The attempted number was incorrect")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "PASS: Guessed the correct number: %d", c.num)
	}
	return
}

func (c *Code) cheatAndSee(w http.ResponseWriter, r *http.Request) {
	c.lck.Lock()
	defer c.lck.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "CHEAT: The number is %d", c.num)
	return
}

func (c *Code) countTries(w http.ResponseWriter, r *http.Request) {
	c.lck.Lock()
	defer c.lck.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "TRIES: %d", c.tries)
	return
}
