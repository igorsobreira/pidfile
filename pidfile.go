// Package pidfile writes a pidfile while the program is running
package pidfile

import (
	"fmt"
	"log"
	"os"
)

type Pidfile struct {
	Name string
}

// New creates a new Pidfile and writes the current PID
func New(name string) *Pidfile {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("pidfile: failed to open %s (%s)", name, err)
		return nil
	}
	defer file.Close()

	pid := fmt.Sprintf("%d", os.Getpid())
	file.Write([]byte(pid))

	return &Pidfile{name}
}

func (pf *Pidfile) Remove() {
	err := os.Remove(pf.Name)
	if err != nil {
		log.Printf("pidfile: failed to remove %s (%s)", pf.Name, err)
	}
}
