package main

import "strings"

func fullMatch(selector string, tags []string) (ok bool) {
	verbs := strings.Split(selector, "|")
	for _, verb := range verbs {
		verb = strings.TrimSpace(verb)
		switch {
		case verb == "$any":
			return true
		case verb[:1] == "!":
			for _, tag := range tags {
				if verb[1:] == tag {
					return false
				}
			}

			ok = true
		default:
			for _, tag := range tags {
				if verb == tag {
					ok = true
				}
			}
		}
	}

	return ok
}
