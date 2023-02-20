package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func fGoRoutine() {

	f("direct") // go primitive

	go f("goroutine")     // A goroutine is a lightweight thread of execution.
	go func(msg string) { // start anonymous function
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second) // time.Sleep() primitive
	fmt.Println("done")
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroups() {
	var wg sync.WaitGroup // sync.WorkGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done() // defer primitive
			worker(i)
		}()
	}

	wg.Wait()
}

func channels() {
	messages := make(chan string) // unbuffered  chan primitive
	go func() {
		messages <- "ping"
	}()

	fmt.Println(<-messages) //channel will accept sends (chan <-) if a corresponding
	// receive ( <- chan )exists
	//sends and receives block until both are ready

}
func bufferedChannels() {
	messages := make(chan string, 2) //channel is buffered now it can send 2 messages without
	// a corresponding receiver
	go func() {
		messages <- "ping" // <- en -> zijn ook primitives
		messages <- "xxxx"
		worker(1)
	}()
	fmt.Println(<-messages) // buffered, receive only last send
	fmt.Println(<-messages)
}

func workerSync(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string) { //chan<- may only receive
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) { //<-chan- may only send
	msg := <-pings
	pongs <- msg
}

func channelBlocking() {
	done := make(chan bool)
	go workerSync(done)
	<-done // blocking receive
}

func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func chanelSelect() {
	c1 := make(chan string, 2)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
		c1 <- "one second time"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
		c2 <- "two second time"
	}()

	for i := 0; i < 4; i++ {
		select { // select primitive
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}

func timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // time.After primitive
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

func nonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages: // non blocking receive
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: // non blocking send
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func closingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			//2-value form of receive, the more value will be false if jobs has been closed and all
			//values in the channel have already been received
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) //close primitive
	fmt.Println("sent all jobs")
	<-done
}

func rangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

func timer() {
	timer1 := time.NewTimer(2 * time.Second) // time.NewTimer()

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}

func ticker() {
	ticker := time.NewTicker(500 * time.Millisecond) // time.NewTicker()
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workerPools() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}

func atomicCounter() {
	var ops uint64
	var wg sync.WaitGroup // wg is a struct
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}

func main() {
	//processes.ManyLanguageElements()
	//fGoRoutine()
	//waitGroups()
	//channels()
	//bufferedChannels()
	//chanelSelect()
	//timeouts()
	//nonBlockingChannelOperations()

	//closingChannels()

	//rangeOverChannels()

	//timer()
	//ticker()

	atomicCounter()
}
