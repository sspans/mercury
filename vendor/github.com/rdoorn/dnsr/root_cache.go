package dnsr

import (
	"strings"

	"github.com/miekg/dns"
)

//go:generate sh generate.sh

var (
	rootCache *cache
)

func init() {
	rootCache = newCache(strings.Count(root, "\n"))
	for t := range dns.ParseZone(strings.NewReader(root), "", "") {
		if t.Error != nil {
			continue
		}
		rr, ok := convertRR(t.RR)
		if ok {
			rootCache.add(rr.Name, rr)
		}
	}
}
