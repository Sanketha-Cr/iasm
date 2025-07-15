package obj

import (
    "os"
    "os/exec"
)

// CurrentOS is the OS that iasm is built for.
const CurrentOS = Linux

func link(dest string, file string) error {
    cmd := exec.Command("ld", "-o", dest, file)
    cmd.Stdin = nil
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}