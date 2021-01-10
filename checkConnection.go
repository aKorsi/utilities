package utilities

import (
	"errors"
	"net"
	"time"
)

func CheckConnection(host , port string, timeOut time.Duration) error {
	connection := net.JoinHostPort(host,port)
	conn, err := net.DialTimeout("tcp", connection, timeOut)
	if err != nil {
		return errors.New(connection + " => Connection error: " + err.Error())
	}
	if conn != nil {
		_ = conn.Close()
	}

	return nil
}
