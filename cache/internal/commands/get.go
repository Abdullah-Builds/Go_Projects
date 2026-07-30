package commands

import (
	"net"

	"github.com/Abdullah-Builds/cache/internal/cache"
)

func Get(parts []string, cacheServer *cache.Cache, conn net.Conn) (string,bool) {
	if len(parts) != 2 {
			return "ERROR usage: GET <key>\n",false
		}

		key := parts[1]

		value, ok := cacheServer.Get(key)
		if ok {
			return (value + "\n"),true
		} else {
			return  "NOT_FOUND\n",true
		}
}
