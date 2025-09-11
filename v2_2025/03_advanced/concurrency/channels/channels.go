package main

import (
	"fmt"
)

/*
	We can set debugChan to true in source code of chan.go file to see more massages while working with channels

	const (
		maxAlign  = 8
		hchanSize = unsafe.Sizeof(hchan{}) + uintptr(-int(unsafe.Sizeof(hchan{}))&(maxAlign-1))
		debugChan = false <-
	)
*/

/*
	How the channel struct look

	type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	timer    *timer // timer feeding this chan
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
*/

func main() {
	//playground with channels

	/*
		this line of code allocates memory in the heap and save pointer (address) to hchan struct

		qcount = 0 at the moment
		dataqsiz = 4

		buf has unsafe.Pointer to buffer

		recvX = 0
		sendX = 0
	*/

	//1. Cone head. computer module. tube body. engine. 4 piece of hi-tech rocket
	rocketLegoCh := make(chan string, 4)

	/*
		First of all we lock mutex to safely work a round with channel, and be afraid other post office send at the same time

		Here I bump sendX value to 1, cuz sendX it is an index that indicates to buf's circular queue.
		tell us

		When to put data in the channel we COPY it, to further modification, original stays untouched unless we send a pointer

		sendX = 1
	*/
	coneHeadPartOriginal := "Cone Head"
	go NasaFactoryApproval(rocketLegoCh, coneHeadPartOriginal)

	coneHeadPartApproved := <-rocketLegoCh

	fmt.Println(coneHeadPartOriginal)
	fmt.Println(coneHeadPartApproved)

	/*
		This example shows us how Circular Queue actually works, we fulled the buffer and sendX has no another option then
		jump back to the beginning. This means channel is full, and it leads to another logic

		sendX = 0

		It will increase the index and write copy value to queue if buffer's cell is empty otherwise we move to another logic
	*/
	go NasaFactoryApproval(rocketLegoCh, coneHeadPartOriginal, "Computer Module", "Tube Body", "Engine")

	/*
		TOPIC NAME: buffer overflow

		at this point we only send data to channel, and it is full. SpaceX comes late we're already building the rocket
		this piece may go for the next rocket.

		This goroutine is on hold (pause) with a help of gopark() function by Go's scheduler
		and added to channel's sendQ field
	*/
	go SpaceXFactoryApproval(rocketLegoCh, coneHeadPartOriginal)

	//var rocketAssembler strings.Builder

	fmt.Printf("channel %v", rocketLegoCh)
}

func NasaFactoryApproval(rocketLegoCh chan string, parts ...string) {
	approvedText := " approved by NASA"

	if len(parts) < 2 {
		rocketLegoCh <- parts[0] + approvedText
		return
	}

	for _, part := range parts {
		rocketLegoCh <- part + approvedText
	}
}

func SpaceXFactoryApproval(rocketLegoCh chan string, parts ...string) {
	approvedText := " approved by SpaceX"

	if len(parts) < 2 {
		rocketLegoCh <- parts[0] + approvedText
		return
	}

	for _, part := range parts {
		rocketLegoCh <- part + approvedText
	}
}
