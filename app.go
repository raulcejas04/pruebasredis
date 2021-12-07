package main

import (
    "fmt"
    "log"
    "net/http"
    "gopkg.in/redis.v3"
)
type User struct {
        mes1 float64
        mes2 float64
        mes3 float64
}

func main() {
    client := redis.NewClient(&redis.Options{
        Addr: "db:6379",
        Password: "",
        DB: 0,
    })
    
    pong, err := client.Ping().Result()
    if err != nil {
        panic(err)
    }
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", pong)
    })

    log.Fatal(http.ListenAndServe(":5000", nil))
    

    newUser := User{
    	    mes1: 1000000,
    	    mes2: 2000000,
    }
    
    for j := 1; j <= 10000; j++ {
    
	_, err := client.Do("HSET", redis.Args{}.Add(j).AddFlat(newUser))
	if err != nil {
        	return err
    	}
    }


}
