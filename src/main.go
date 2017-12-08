package main

import "os"

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage           : ./damm <input>\nValid Example   : ./damm 123; echo \"$?\"\nInvalid Example : ./damm 12; echo \"$?\"\n"))
		return
	}

	if !Validate(os.Args[1]) {
		os.Exit(1)
		return
	}

	os.Exit(0)
}
