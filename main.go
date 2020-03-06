package main

import (
    "errors"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "os"
)

func externalIP() (string, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, iface := range ifaces {
        if iface.Flags & net.FlagUp == 0 {
            continue
        }

        if iface.Flags & net.FlagLoopback != 0 {
            continue
        }

        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }

        for _, addr := range addrs {
            var ip net.IP

            switch v := addr.(type) {
                case *net.IPNet:
                    ip = v.IP
                case *net.IPAddr:
                    ip = v.IP
            }

            if ip == nil || ip.IsLoopback() {
                continue
            }

            ip = ip.To4()

            if ip == nil {
                continue
            }

            return ip.String(), nil
        }
    }

    return "", errors.New("No IP found!")
}

func main() {
    ip, err := externalIP()
    if err != nil {
        fmt.Println(err)
    }

    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println(err)
    }

    output := fmt.Sprintf("%s (%s)\n", hostname, ip)

    myHandler := func(w http.ResponseWriter, req *http.Request) {
        log.Printf("Request from %s", req.RemoteAddr)
        io.WriteString(w, output)
    }

    port := ":8080"

    if len(os.Args[1:]) > 0 {
        port = os.Args[1]
    }

    if port[0] != ':' {
        port = fmt.Sprintf(":%s", port)
    }

    http.HandleFunc("/", myHandler)
    log.Printf("Starting server on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
