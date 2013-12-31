# Go package to manage pid files

Usage:

    import "github.com/igorsobreira/pidfile"

    func main() {

        pf := pidfile.New("/var/run/myprogram.pid")
        defer pf.Remove()

    }
