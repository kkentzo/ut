package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func convert(arg string) (time.Time, error) {
	t, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to convert %s", arg)
	}

	var sec, nsec int64

	switch len(arg) {
	case 10:
		// seconds
		sec = t
		nsec = 0
	case 13:
		// milliseconds
		sec = t / 1000
		nsec = (t % 1000) * 1_000_000
	case 16:
		// microseconds
		sec = t / 1_000_000
		nsec = (t % 1_000_000) * 1000
	case 19:
		// nanoseconds
		sec = t / 1_000_000_000
		nsec = t % 1_000_000_000
	default:
		return time.Time{}, fmt.Errorf("unrecognized resolution in timestamp %s. Supported timestamps are in seconds, milliseconds, microseconds or nanoseconds.", arg)
	}

	return time.Unix(sec, nsec), nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if t, err := convert(arg); err != nil {
			fmt.Printf("[error] %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println(t.Format(time.RFC3339Nano))
		}
	}

}
