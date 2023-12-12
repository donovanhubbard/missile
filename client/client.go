/*
Copyright © 2023 Donovan Hubbard
*/
package client

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

func Connect(hosts ...string) {
	fmt.Println("Running")
	fmt.Println(hosts)

	// Create a ServerSelector
	ss := new(memcache.ServerList)
	ss.SetServers(hosts...)
	mc := memcache.NewFromSelector(ss)
	Ping(mc)
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	thing, _ := mc.Get("foo")
	fmt.Println(thing)
	addr, err := ss.PickServer("foo")

	if err != nil {
		fmt.Errorf("Failed to pick item.")
	}

	fmt.Printf("The item was saved at server %s\n", addr)

	mc.Close()
	fmt.Println("Done")
}

func Ping(mc *memcache.Client) {
	err := mc.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pong")
	}
}
