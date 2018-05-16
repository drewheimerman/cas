package main

import(
	"fmt"
	"log"
	"github.com/gocql/gocql"
)

// keyspace: demo;
// table cas (key text, id text, ts int, val text, phase int)

// get cassandra session
func getSession(addr string) *gocql.Session {
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = "demo"
	cluster.Consistency = gocql.One
	session,err  := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
	}
	return session
}

// query cassandra to get val with key
func queryGet(key string) string {
	var res string
	arg := fmt.Sprintf("SELECT val FROM cas WHERE key='%s'", key)
	if err := session.Query(arg).Scan(&res); err != nil {
		log.Fatal(err)
	}
	return res
}

func queryPre() {
	arg := fmt.Sprintf("UPDATE cas SET id='%s', val='%s', ts=%d WHERE key='%s' AND phase='fin' AND ts<%d", tv.Tag.Id, tv.Tag.Ts, tv.Val, tv.Key)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
}

func queryFinW(ket string, tag Tag) {
	arg := fmt.Sprintf("UPDATE cas SET phase=%d WHERE key='%s' AND phase=%d AND ts=%d AND id='%s'", tv.Tag.Id, tv.Tag.Ts, tv.Val, tv.Key)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
}

func queryFinR(ket string, tag Tag) {
	var tv Tagval
	
	arg := fmt.Sprintf("UPDATE cas SET phase=%d WHERE key='%s' AND phase=%d AND ts=%d AND id='%s'", tv.Tag.Id, tv.Tag.Ts, tv.Val, tv.Key)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
	
	return tv
}