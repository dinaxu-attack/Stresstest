package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var success int
var all int
var errors int

var wg sync.WaitGroup

func MainStress(target string, count string) {
	sttime := time.Now().Format("15:04:05")
	dcount, _ := strconv.Atoi(count)
	wg.Add(1)
	for i := 0; i <= 100; i++ {
		go stress(target, dcount)
		go stress(target, dcount)
		go stress(target, dcount)
		go stress(target, dcount)
		go stress(target, dcount)
		go stress(target, dcount)
		go stress(target, dcount)
	}
	wg.Wait()

	println()
	println("Success: " + strconv.Itoa(success) + "\nErrors: " + strconv.Itoa(errors))
	fmt.Println(sttime + " - " + time.Now().Format("15:04:05"))

}

func stress(target string, count int) {

	for i := 0; i <= count; i++ {

		if all >= count {
			break
		}

		req, err := http.NewRequest("GET", target, nil)
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		}

		_, err = client.Do(req)

		all++
		if err != nil {
			errors++
		} else {
			success++
		}

		fmt.Printf("\rSent: " + strconv.Itoa(all))

	}
	wg.Done()
}
