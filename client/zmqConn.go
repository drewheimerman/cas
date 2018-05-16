package main 

import (
	"fmt"
	zmq "github.com/pebbe/zmq3"
)

func createDealerSocket() *zmq.Socket {
	dealer,_ := zmq.NewSocket(zmq.DEALER)
	var addr string
	for _,server := range servers {
		addr = "tcp://" + server + ":5555"
		dealer.Connect(addr)
		fmt.Println(addr)
	}
	return dealer
}

// broadcast
func sendToServer(msg Message, dealer *zmq.Socket) {
	msgToSend := getGobFromMsg(msg)
	for i := 0 ; i < len(servers); i++ {
		dealer.SendBytes(msgToSend.Bytes(), 0)
	}
}

// recv tag
func recvTag(dealer *zmq.Socket) Tag {
	msgBytes,_ := dealer.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != QUERY {
		return recvRes(dealer)
	}
	return msg.Tag
}

// recv empty ack
func recvAck(req *zmq.Socket, done chan bool) {
	msgBytes,_ := req.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != PREWRITE {
		recvAck(req, done)
	} else{
		done<-true
		req.Close()
	}
}

// recv data
func recvRes(dealer *zmq.Socket) string {
	msgBytes,_ := dealer.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != FINALIZE {
		return recvRes(dealer)
	}
	return msg.Val
}