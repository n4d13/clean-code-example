package main

import (
	"github.com/n4d13/clean-code-example/src/db"
	"log"
	"fmt"
	"os"
	"github.com/n4d13/clean-code-example/src/config"
)

func main() {

	conf := config.GetInstance()
	conn, err := db.MakeRedisConnection(conf.Address(), conf.DialTimeout,
		conf.ReadTimeout, conf.WriteTimeout, conf.PoolSize, conf.PoolTimeout)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "Total keys count: %d", conn.Count("key1", "key2"))

}
