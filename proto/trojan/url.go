package trojan

import (
	"errors"
	"net/url"
)

// ParseURL is ...
func ParseURL(s string) (server, path, password, transport, mux, domain string, err error) {
	u, err := url.Parse(s)
	if err != nil {
		return
	}

	server = u.Host
	if u.User == nil {
		err = errors.New("no user info")
		return
	}

	if s := u.User.Username(); s != "" {
		password = s
	} else {
		err = errors.New("no password")
		return
	}

	path = u.Path

	transport = u.Query().Get("transport")
	switch transport {
	case "":
		transport = "tls"
	case "tls", "websocket":
	default:
		err = errors.New("wrong transport")
		return
	}

	mux = u.Query().Get("mux")
	switch mux {
	case "":
		mux = "off"
	case "off", "v1", "v2":
	default:
		err = errors.New("wrong mux config")
		return
	}

	domain = u.Fragment
	if domain == "" {
		err = errors.New("no domain name")
	}
	return
}
