package commands

import (
	"net"

	"github.com/Abdullah-Builds/cache/internal/cache"
)

func Delete(parts []string, cacheServer *cache.Cache, conn net.Conn) (string,bool) {
	if len(parts) != 2 {
			conn.Write([]byte("ERROR usage: DELETE <key>\n"))
			return "ERROR usage: DELETE <key>\n",false
		}

		key := parts[1]

		cacheServer.Delete(key)
		return  "OK\n",true
}
