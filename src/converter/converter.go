package converter

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	/*
		Default inkscape binary path
	*/
	BINARY = "/usr/bin/inkscape"
)

type Converter struct {
	bin string
}

func New() *Converter {
	var c Converter
	c.bin = BINARY
	return &c
}

func (c *Converter) Convert(in []byte) (out []byte, err error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(c.bin, "-p", "--export-type=png", "-o", "-")
	cmd.Stdin = bytes.NewBuffer(in)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if e := cmd.Run(); e != nil {
		err = fmt.Errorf("%s\nSTDERR:\n%s", e.Error(), stderr.String())
		return
	}

	if stdout.Len() == 0 {
		err = fmt.Errorf("Got no data from inkscape")
		return
	}

	out = stdout.Bytes()
	return
}
