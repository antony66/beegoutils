package beegoutils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/antony66/go-sshwrapper"
	"github.com/astaxie/beego"
)

// ServerInterface represents interface to a Server
type ServerInterface interface {
	GetName() string
	GetSSHURL() string
}

func getSSHConn(server ServerInterface) (*sshwrapper.SSHConn, error) {
	return sshwrapper.Dial(server.GetSSHURL(),
		beego.AppConfig.String("ssh_agent"),
		false)
}

// RunSSHCommand executes command with stdin on server and returns output and error
func RunSSHCommand(server ServerInterface, command string, stdin io.Reader) ([]byte, error) {
	var outp []byte
	var err error
	fake := beego.AppConfig.DefaultBool("fake_ssh", false)
	buf := new(bytes.Buffer)
	log.Printf("SSH Command (fake=%v) on %s: %s\n", fake, server.GetName(), command)
	if stdin != nil {
		tee := io.TeeReader(stdin, buf)
		if b, e := ioutil.ReadAll(tee); e == nil {
			log.Printf("STDIN: %s\n", b)
		}
	}
	if !fake {
		var conn *sshwrapper.SSHConn
		conn, err = getSSHConn(server)
		defer conn.Close()
		if err == nil {
			outp, err = conn.CombinedOutput(command, buf)
		}
		if err != nil {
			err = fmt.Errorf("Error executing ssh command on %s: %s: %s\n", server.GetName(), err, string(outp))
			log.Println(err)
		}
	}
	return outp, err
}
