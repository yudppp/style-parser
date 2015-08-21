package style

import (
	"strings"
)

type Parser struct{}

func (p *Parser) ParseFromString(s string) map[string]string {

	if s == "" {
		return nil
	}

	styles := strings.Split(s, ";")

	set := make(map[string]string, len(styles))

	for _, v := range styles {
		ss := strings.Split(v, ":")
		if len(ss) != 2 {
			continue
		}
		propety := strings.TrimSpace(ss[0])
		value := strings.TrimSpace(ss[1])
		set[propety] = value
	}
	return set
}
