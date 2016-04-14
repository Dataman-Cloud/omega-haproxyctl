package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Runtime struct {
	Binary   string
	SockFile string
}

// Sets the weight of a backend
func (r *Runtime) SetWeight(backend string, server string, weight int) (string, error) {
	result, err := r.cmd("set weight " + backend + "/" + server + " " + strconv.Itoa(weight) + "\n")

	if err != nil {
		return "", err
	} else {
		return result, nil
	}

}

// Executes a arbitrary HAproxy command on the unix socket
func (r *Runtime) cmd(cmd string) (string, error) {

	// connect to haproxy
	conn, err_conn := net.Dial("unix", r.SockFile)
	defer conn.Close()

	if err_conn != nil {
		return "", errors.New("Unable to connect to Haproxy socket")
	} else {

		fmt.Fprint(conn, cmd)

		response := ""

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			response += (scanner.Text() + "\n")
		}
		if err := scanner.Err(); err != nil {
			return "", err
		} else {
			return response, nil
		}

	}
}
