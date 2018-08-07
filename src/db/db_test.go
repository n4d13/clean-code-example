package db

import "testing"
import (
	"github.com/alicebob/miniredis"
	"time"
)

func TestMakeRedisConnection(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	conn, err := MakeRedisConnection(s.Addr(), 5*time.Second,
		1*time.Second, 1*time.Second, 2, 30*time.Second)

	if err != nil {
		t.Fatal(err)
	}

	conn.client.Set("key1", "value1", 0)
	conn.client.Set("key2", "value2", 0)

	count := conn.Count("key1", "key2")

	if count != 2 {
		t.Fatalf("Failed assert. Expecting %d but got %d", 2, count)
	}

}


