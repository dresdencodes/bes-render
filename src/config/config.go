package config

import (
	"log"
	"sync"
	"time"
)

type Config struct {
	PocketAuth 		*PocketAuth
	PocketIP 		string
}

var mu = &sync.Mutex{}
var cfg = &Config{
	PocketIP:"http://149.28.13.238:8091",
}


func Get() *Config {
	mu.Lock(); defer mu.Unlock()
	cpy := *cfg
	return &cpy
}

func Inits() {
	
	// pocket auth 
	pocketAuth, err := PocketAuthorize(Get().PocketIP)
	if err!=nil {
		log.Fatal(err)
	}

	// lock / unlock
	mu.Lock(); defer mu.Unlock()
	cfg.PocketAuth = pocketAuth

	// re run
	time.AfterFunc(time.Duration(180) * time.Second, Inits)

}