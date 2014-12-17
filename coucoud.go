package main

import (
	"fmt"
)

// utmp (5)
type Utmp struct {
	ut_type		uint16
	ut_pid		uint16
	ut_line		[32]byte
	ut_id		[4]byte
	ut_user		[32]byte
	ut_host		[256]byte
	ut_session	uint32
	ut_tv_sec	uint32
	ut_tv_usec	uint32
	ut_addr_v6	[4]int32
	unused		[20]byte
}

func main() {
	fmt.Printf("coucou\n")
}
