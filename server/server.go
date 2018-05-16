package main 

import (
	"fmt"
	"log"
//	"github.com/gocql/gocql"
	zmq "github.com/pebbe/zmq3"
)

type State struct {
	Tag 	Tag
	Val 	string
	Phase 	int
}

func server_task() {
	// Set the Default State Variables
	state = State{Tag: Tag{Id: "", Ts: 0}, Val: "", Phase: FINALIZE}

	session = getSession("172.0.0.2")
	defer session.Close()
	// Set the ZMQ sockets
	frontend,_ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	frontend.Bind("tcp://5555")

	//  Backend socket talks to workers over inproc
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	fmt.Println("frontend router", "tcp://5555")
	go server_worker()

	//  Connect backend to frontend via a proxy
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatal("Proxy interrupted:", err)
}

func server_worker() {
	worker, _ := zmq.NewSocket(zmq.DEALER)
	defer worker.Close()
	worker.Connect("inproc://backend")
	msg_reply := make([][]byte, 2)

	for i := 0; i < len(msg_reply); i++ {
		msg_reply[i] = make([]byte, 0) // the frist frame  specifies the identity of the sender, the second specifies the content
	}

	for {
		msg,err := worker.RecvMessageBytes(0)
		if err != nil {
			fmt.Println(err)
		}

		message := getMsgFromGob(msg[1])
		msg_reply[0] = msg[0]
		fmt.Println(msg[0])

		tmpMsg := createRep(message)
		tmpGob := getGobFromMsg(tmpMsg)
		msg_reply[1] = tmpGob.Bytes()

		worker.SendMessage(msg_reply)
	}
}