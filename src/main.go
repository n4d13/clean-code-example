package main

import (
	"github.com/n4d13/clean-code-example/src/db"
	"log"
	"fmt"
	"os"
)

func main(){

	conn, err := db.MakeRedisConnection()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout,"Total keys count: %d", conn.Count("key1", "key2"))


}
