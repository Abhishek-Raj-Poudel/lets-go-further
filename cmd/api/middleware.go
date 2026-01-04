package main

import (
	"fmt"
	"net/http"
	"net"
	"sync"
	"golang.org/x/time/rate"
	"time"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// func (app *application) rateLimit(next http.Handler) http.Handler {
// 	limiter := rate.NewLimiter(2, 4)

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if !limiter.Allow() {
// 			app.rateLimitExceededResponse(w, r)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
// 

func (app *application) rateLimit(next http.Handler) http.Handler{
	type client struct{
		limiter *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu sync.Mutex
		clients = make(map[string]*client)
	)
// check every minute when a client has been seen. if the ip was last seen 3 min ago then remove it	
	go func(){
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip,client := range clients{
				if time.Since(client.lastSeen) > 3*time.Minute{
					delete(clients,ip)
				}
			}
			mu.Unlock()
		}
	}()
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// first take ip address of that request
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err !=nil {
			app.serverErrorResponse(w,r,err)
			return
		}
		// lock it so that all threads access to this ip
		mu.Lock()
		
		// check if ip already exist's in map if doesn't then create a New rate limiter
		if _,found := clients[ip]; !found {
			// Create and add a new client struct to the map if it doesn't already exist.
			clients[ip] = &client{limiter:rate.NewLimiter(2,4)} 
		}
	// Update the last seen time for the client	
		clients[ip].lastSeen = time.Now()
		
		if !clients[ip].limiter.Allow(){
			mu.Unlock()
			app.rateLimitExceededResponse(w,r)
			return
		}
		
		mu.Unlock()

		next.ServeHTTP(w,r)
	})
}


