package main 

import (
//	zmq "github.com/pebbe/zmq3"
)

func write(key string, val []byte) {
	t := query()
	t.update(ID)
	preWrite(t,key,val)
	wfinalize(t)
}

func read(key string) []byte {
	t := query()
	v := rfinalize(t)
	return v
}

func query() Tag {
	dealer := createDealerSocket()
	defer dealer.Close()

	t := Tag{Id: "", Ts: 0}
	msg := Message{OpType: QUERY, Tag: t, Key: "", Val: ""}
	
	go sendToServer(msg, dealer)

	for i := 0; i < len(servers)/2 + 1; i++ {
		tmp := recvTag(dealer)
		if t.smaller(tmp) {
			t = tmp
		}
	}
	return t
}

func preWrite(t Tag, key string, val string){
	//////////////////////////
	// w := phi.encode(val) //
	//////////////////////////
	msg := Message{OpType: PREWRITE, Tag: t, Key: key Val: ""}
	
	done := make(chan bool)
	for idx,server := range servers {
		msg.Val = w[idx]
		req := createReqSocket(server)
		go sendToServer(msg, req)
		go recvAck(req, done)
	}
	
	for i := 0; i < len(servers)/2 + 1; i++ {
		<-done
	}
}

func wfinalize(t Tag) {
	dealer := createDealerSocket()
	defer dealer.Close()

	msg := Message{OpType: FINALIZE, Tag: t, Key: "", Val: ""}
	
	go sendToServer(msg, dealer)

	for i := 0; i < len(servers)/2 + 1; i++ {
		recvRes(dealer)
	}
}

func rfinalize(t Tag) string {
	dealer := createDealerSocket()
	defer dealer.Close()

	msg := Message{OpType: FINALIZE, Tag: t, Key: "", Val: ""}
	
	go sendToServer(msg, dealer)

	w := make([]string)
	for i := 0; i < len(servers)/2 + 1; i++ {
		w = append(w,recvRes(dealer))
	}
	////////////////////////
	//return phi.decode(w)//
	////////////////////////
}