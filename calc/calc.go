package calc

import (
    "net"
    "strings"
    "io"
    "bufio"
    "log"
)

var (
    Debug = true
)

func FullRange(reader io.Reader) ([]net.IP, error) {
    raw := make([]string, 0)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        raw = append(raw, line)
    }

    if Debug {
        log.Printf("len(raw): %d\n", len(raw))
    }

    dedupped := dedup(raw)

    if Debug {
        log.Printf("len(dedupped): %d\n", len(dedupped))
    }

    ipr := make([]net.IP, 0)

    for _, item := range dedupped {
        temp1, err := iprange(item)

        if err != nil {
            return []net.IP{}, err
        }

        temp2 := highlow(temp1)

        for _, ip := range temp2 {
            ipr = append(ipr, ip)
        }
    }

    hl := highlow(ipr)

    return hl, nil
}






