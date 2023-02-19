package main

import (
	"fmt"
	"sync"
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

	rangeOverChannels()
}
