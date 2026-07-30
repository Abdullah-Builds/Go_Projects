package commands

import (
	"net"
	"strconv"
	"time"

	"github.com/Abdullah-Builds/cache/internal/cache"
)

func Set(parts []string, cacheServer *cache.Cache, conn net.Conn) (string, bool) {

	var ttl time.Duration

	if len(parts) == 4 {
		seconds, err := strconv.Atoi(parts[3])
		if err != nil {
			conn.Write([]byte("ERROR invalid TTL\n"))
			return "ERROR invalid TTL\n", false
		}

		ttl = time.Duration(seconds) * time.Second
	}

	// if len(parts) != 3 {
	// 	return "ERROR usage: SET <key> <value>\n", false
	// }

	key := parts[1]
	value := parts[2]

	cacheServer.Set(key, value, ttl)
	return "OK", true
}
