package concurrency

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func NormalPrint() {
	log.Println("NormalPrint")
	start := time.Now()

	time.Sleep(time.Second)
	fmt.Println("Hello, World")
	time.Sleep(time.Second)
	fmt.Println("This is")
	time.Sleep(time.Second)
	fmt.Println("Only PRINTING!")

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func GoRoutinesPrint() {
	var wg sync.WaitGroup
	wg.Add(3)
	log.Println("GoRoutinesPrint")

	start := time.Now()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("Hello, World")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("This is")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("Only PRINTING!")
	}()

	wg.Wait()

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func NormalLoop() {
	log.Println("NormalLoop")
	start := time.Now()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func GoRoutinesLoop() {
	var wg sync.WaitGroup
	wg.Add(5)
	log.Println("GoRoutinesLoop")
	start := time.Now()

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(i)
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func ChannelPrint() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	log.Println("ChannelPrint")

	start := time.Now()

	go func(ch1 chan bool) {
		time.Sleep(time.Second)
		fmt.Println("Hello, World")
		ch1 <- true
		close(ch1)
	}(ch1)
	go func(ch1, ch2 chan bool) {
		time.Sleep(time.Second)
		<-ch1
		fmt.Println("This is")
		ch2 <- true
		close(ch2)
	}(ch1, ch2)
	go func(ch2, ch3 chan bool) {
		time.Sleep(time.Second)
		<-ch2
		fmt.Println("Only PRINTING!")
		close(ch3)
	}(ch2, ch3)

	<-ch3

	elapsed := time.Since(start)
	log.Println(elapsed)

}

func ChannelWaitGroupPrint() {
	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	log.Println("ChannelWaitGroupPrint")

	start := time.Now()

	go func(ch1 chan bool) {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("Hello, World")
		ch1 <- true
	}(ch1)
	go func(ch1, ch2 chan bool) {
		defer wg.Done()
		time.Sleep(time.Second)
		<-ch1
		fmt.Println("This is")
		ch2 <- true

	}(ch1, ch2)
	go func(ch2 chan bool) {
		defer wg.Done()
		time.Sleep(time.Second)
		<-ch2
		fmt.Println("Only PRINTING!")
	}(ch2)

	wg.Wait()

	elapsed := time.Since(start)
	log.Println(elapsed)

}

// still deadlock
func ChannelLoop() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(5)
	log.Println("ChannelLoop")
	start := time.Now()

	for i := 0; i < 5; i++ {
		fmt.Printf("start %d\n", i)
		go func(i int) {
			defer wg.Done()
			defer fmt.Printf("end %d\n", i)
			time.Sleep(time.Second)
			num := <-ch
			if num == i {
				fmt.Printf("inif %d\n", num)
				if i < 4 {
					ch <- i + 1
				}
			}
		}(i)
	}

	ch <- 0

	wg.Wait()

	elapsed := time.Since(start)
	log.Println(elapsed)
}

func Worker() {
	var wg sync.WaitGroup
	ch := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- fmt.Sprintf("Worker %d done", i)
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
