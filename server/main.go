package main 

import(
	"github.com/gocql/gocql"
)
	


var (
	state 	Tag
	session *gocql.Session
)

const GET=0
const SET=1

func main(){
	server_task()
}
