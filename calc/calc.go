package calc

import (
    "net"
    "strings"
    "io"
    "bufio"
    "log"
    "errors"
)

type CalculatorServiceOption func(*CalculatorService) error

type CalculatorService struct {
    Thing1 string
}

func NewCalculatorService(options ...CalculatorServiceOption) (*CalculatorService, error) {
    service := &CalculatorService{}

    for _, option := range options {
        if err := option(service); err != nil {
            return nil, err
        }
    }

    return service, nil
}

func Thing1(input string) CalculatorServiceOption {
    return func(this *CalculatorService) error {
        if input == "" {
            return nil, errors.New("Invalid thing1")
        }

        this.Thing1 = input
        return nil
    }
}

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






