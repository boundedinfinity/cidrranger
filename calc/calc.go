package calc

import (
    "net"
    "strings"
    "io"
    "bufio"
    "log"
    "os"
)

type CalculatorService struct {
    InputPath              string
    Logger                 *log.Logger
    Debug                  bool
    RemoveBroadcastAddress bool
    RemoveNetworkAddress   bool
}

func NewCalculatorService(options ...CalculatorServiceOption) (*CalculatorService, error) {
    service := &CalculatorService{
        Logger: log.New(os.Stdout, "CalculatorService", log.Lshortfile),
        Debug: false,
    }

    if err := service.SetOptions(options...); err != nil {
        return nil, err
    }

    return service, nil
}

func (this *CalculatorService) SetOptions(options ...CalculatorServiceOption) error {
    for _, option := range options {
        if err := option(this); err != nil {
            return err
        }
    }

    return nil
}

func (this *CalculatorService) loadFromFile() (io.ReadCloser, error) {
    file, err := os.Open(this.InputPath)

    if err != nil {
        return nil, err
    }

    reader := bufio.NewReader(file)

    return struct {
        io.Reader
        io.Closer
    }{reader, file}, nil
}

func (this *CalculatorService) FullRange() ([]net.IP, error) {
    reader, err := this.loadFromFile()

    if err != nil {
        return nil, err
    }

    defer reader.Close()

    raw := make([]string, 0)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if !this.ignoreLine(line) {
            raw = append(raw, line)
        }
    }

    if this.Debug {
        this.Logger.Printf("len(raw): %d\n", len(raw))
    }

    dedupped := this.dedup(raw)

    if this.Debug {
        this.Logger.Printf("len(dedupped): %d\n", len(dedupped))
    }

    ipr := make([]net.IP, 0)

    //for _, item := range dedupped {
        //temp1, err := this.SubnetEndpointsFromCidr(item)

        //if err != nil {
        //    return []net.IP{}, err
        //}
	//
        //temp2 := this.highlow(temp1)
	//
        //for _, ip := range temp2 {
        //    ipr = append(ipr, ip)
        //}
    //}

    //hl := this.highlow(ipr)

    return ipr, nil
}






