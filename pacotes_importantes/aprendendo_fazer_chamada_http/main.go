package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Uma das coisas como desenvolvedor é entender como fazer chamadas http da forma mais performática possível
// Um dos pontos principais que vai fazer a diferença crucial se seu sistema vai ser performático ou não, é os limites de chamadas externas
// que o sistema vai realizar

// trabalhando http com timeout
func main() {
	c := http.Client{Timeout: time.Duration(time.Microsecond)}
	res, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

}
