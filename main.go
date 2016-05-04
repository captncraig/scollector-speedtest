package main

import (
	"fmt"
	"time"

	"github.com/zpeters/speedtest/debug"
	"github.com/zpeters/speedtest/settings"
	"github.com/zpeters/speedtest/sthttp"
	"github.com/zpeters/speedtest/tests"
)

func main() {
	debug.QUIET = true

	sthttp.CONFIG = sthttp.GetConfig()

	servers := sthttp.GetServers()
	servers = sthttp.GetClosestServers(servers)
	best := sthttp.GetFastestServer(servers)
	dl := tests.DownloadTest(best)
	ul := tests.UploadTest(best)
	ping := best.Latency

	settings.ALGOTYPE = "avg"
	dlavg := tests.DownloadTest(best)
	ulavg := tests.UploadTest(best)
	ts := time.Now().UTC().Unix()
	fmt.Printf("speedtest.dl %d %f measure=max server=%s \n", ts, dl, best.ID)
	fmt.Printf("speedtest.ul %d %f measure=max server=%s \n", ts, ul, best.ID)
	fmt.Printf("speedtest.dl %d %f measure=avg server=%s \n", ts, dlavg, best.ID)
	fmt.Printf("speedtest.ul %d %f measure=avg server=%s \n", ts, ulavg, best.ID)
	fmt.Printf("speedtest.ping %d %f \n", ts, ping)
}
