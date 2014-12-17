package main

import (
	"fmt"
	"os"
	"time"
	"encoding/binary"
)

const (
	UTMP_PATH = "/var/run/utmp"
	REFRESH_EVERY_NSEC = 5
)

// utmp (5)
type Utmp struct {
	UtType		uint16
	UtPid		uint16
	UtLine		[32]byte
	UtId		[4]byte
	UtUser		[32]byte
	UtHost		[256]byte
	UtSession	uint64
	UtTvSec		uint32
	UtTvUsec	uint32
	UtAddrV6	[4]int32
	Unused		[20]byte
}

type Coucoud struct {
}

// parses the content of UNIX' utmp file
func (c *Coucoud) parseUtmp() error {
	r, err := os.OpenFile(UTMP_PATH, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer r.Close()

	var u Utmp
	for ; err == nil ; err = binary.Read(r, binary.LittleEndian, &u) {
		fmt.Printf("-> %s\n", string(u.UtUser[0:32]))
	}

	return nil
}

// main loop, watch for new users
func (c *Coucoud) loop() {
	for {
		c.parseUtmp()
		fmt.Printf("sleeping for a for -- no while in go :'(\n")
		time.Sleep(REFRESH_EVERY_NSEC * 1e9)
	}
}

func NewCoucoud() *Coucoud {
	return &Coucoud{}
}

func main() {
	c := NewCoucoud()
	c.loop()
}
