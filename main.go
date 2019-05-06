package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

const (
	// SECSFILE .... filename contains text to be displayed on stdout every second
	SECSFILE = "secsFile"
	// MNTSFILE .... filename contains text to be displayed on stdout every minue
	MNTSFILE = "mntsFile"
	// HOURSFILE .... filename contains text to be displayed on stdout every hour
	HOURSFILE = "hoursFile"
)

func main() {
	// run for Three Hours[uncomment the below line to run the app]
	// Clock(3 * time.Hour)
}

// Clock ...
func Clock(waitFor time.Duration) {
	done := make(chan bool)
	t1 := Schedule(Reader, HOURSFILE, time.Duration(1)*time.Hour, done)
	t2 := Schedule(Reader, MNTSFILE, time.Duration(1)*time.Minute, done)
	t3 := Schedule(Reader, SECSFILE, time.Duration(1)*time.Second, done)

	time.Sleep(waitFor)

	close(done)
	t3.Stop()
	t2.Stop()
	t1.Stop()
}

// Reader ...
func Reader(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	text, _ := reader.ReadString('\n')
	log.Println(text)
}

// Schedule ....
func Schedule(f func(string), file string, interval time.Duration, done <-chan bool) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				f(file)
			case <-done:
				return
			}
		}
	}()
	return ticker
}
