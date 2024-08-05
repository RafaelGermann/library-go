package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
//Generate secret_key

	func init() {
		chave := make([]byte, 64)

		if _, err := rand.Read(chave); err != nil {
			log.Fatal(err)
		}
		stringBase64 := base64.StdEncoding.EncodeToString(chave)
		fmt.Println(stringBase64)
	}
*/
func main() {
	config.Carregar()
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
