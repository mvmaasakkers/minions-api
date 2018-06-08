package mongodb

import (
	"gopkg.in/mgo.v2"
	"net"
	"crypto/tls"
	"strings"
	"time"
)

type DB struct {

	Session  *mgo.Session
	Settings struct {
		Host     string
		Database string
	}
}

func (d *DB) Open() error {
	// Only open session if none is set yet (for testing)
	if d.Session == nil {
		dialInfo, err := mgoParseURI(d.Settings.Host)
		if err != nil {
			return err
		}

		d.Session, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DB) Close() error {
	if nil != d.Session {
		d.Session.Close()
		return nil
	}
	return nil
}

func mgoParseURI(url string) (*mgo.DialInfo, error) {
	isSSL := strings.Index(url, "ssl=true") > -1
	// Remove ssl option because it is unsupported by mgo ParseURL
	url = strings.Replace(url, "ssl=true", "", 1)

	// Remove other options that are unsupported by mgo ParseURL
	url = strings.Replace(url, "retryWrites=true", "", 1)

	dialInfo, err := mgo.ParseURL(url)

	if err != nil {
		return nil, err
	}

	if isSSL {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true

		dialInfo.Timeout = time.Second * 10
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	return dialInfo, err
}
