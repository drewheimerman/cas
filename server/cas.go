package main 

import(
)

func createRep(input Message) Message {
	var output Message
	switch input.OpType{
	case QUERY:
		t := queryTag(input.Key)
		output = Message{OpType: QUERY, Tag: t, Key: "", Val: ""}
	case PREWRITE:
		querySet(input)
		output = Message{OpType: PREWRITE, Tag: Tag{Id: "", Ts: 0}, Key: "", Val: ""}
	case WFINALIZE:
		queryWFinal(input)
		output = Message{OpType: WFINALIZE, Tag: Tag{Id: "", Ts: 0}, Key: "", Val: ""}
		go gossip(input)
	case RFINALIZE:
		res := queryRFinal(input)
		output = Message{OpType: WFINALIZE, Tag: Tag{Id: "", Ts: 0}, Key: input.Key, Val: res}
		go gossip(output)
	case GOSSIP:
		queryGossip(input)
	}	
	return output
}

func gossip(msg Message){
	msgToSend := getGobFromMsg(msg)
	for idx,server := range servers{
		sock,_ := zmq.NewSocket(zmq.PUSH)
		sock.Connect(server)
		sock.SendBytes(msgToSend.Bytes(), 0)
		sock.Close()
	}
}
