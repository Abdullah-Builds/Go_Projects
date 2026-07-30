package main

<<<<<<< HEAD
import "fmt"

func main() {
	fmt.Println("cache server")
=======
import (
	"log"
	"net"
	"time"

	"github.com/Abdullah-Builds/cache/internal/cache"
	"github.com/Abdullah-Builds/cache/internal/handler"
)

func main() {

	var CacheServer = cache.New()
	CacheServer.StartCleanup(time.Second)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Cache server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go handler.HandleConnection(conn, CacheServer)
	}
>>>>>>> 43439c3 (feat: cache version 1)
}
