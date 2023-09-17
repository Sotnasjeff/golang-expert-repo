package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		//Usando solução atómica
		//atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você tem acesso a essa página %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)
}
