package utils

import (
	"fmt"
	"net"
	"strings"

	"github.com/Abdullah-Builds/cache/internal/cache"
	"github.com/Abdullah-Builds/cache/internal/commands"
)

func PerformOperations(line string, cacheServer *cache.Cache, conn net.Conn) {
	line = strings.TrimSpace(line)

	if line == "" {
		return
	}

	parts := strings.Fields(line)

	command := strings.ToUpper(parts[0])

	switch command {

	case "SET":
		msg, isvalid := commands.Set(parts, cacheServer, conn)
		fmt.Println(msg)
		conn.Write([]byte(msg+ "\n"))

		if !isvalid {
			return
		}

	case "GET":
		msg, isvalid := commands.Get(parts, cacheServer, conn)
		conn.Write([]byte(msg))

		if !isvalid {
			return
		}

	case "DELETE":
		msg, isvalid := commands.Get(parts, cacheServer, conn)
		conn.Write([]byte(msg))

		if !isvalid {
			return
		}

	case "PING":
		conn.Write([]byte("PONG\n"))

	default:
		conn.Write([]byte("ERROR unknown command\n"))
	}
}
