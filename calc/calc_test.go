package calc

import (
    "testing"
    "net"
)

func getService() *CalculatorService {
    service, _ := NewCalculatorService(
        InputPath("./subnet-list-1.txt"),
        Debug(true),
    )

    return service
}

func TestNetworkAddressFromCidr(t *testing.T) {
    service := getService()
    m := map[string]net.IP{
        "165.137.176.0/24": net.IP([]byte{165, 137, 176, 0}),
        "165.137.176.50/24": net.IP([]byte{165, 137, 176, 0}),
        "165.137.176.5/16": net.IP([]byte{165, 137, 0, 0}),
        "165.137.176.255/24": net.IP([]byte{165, 137, 176, 0}),
    }

    for cidr, expected := range m {
        actual, err := service.NetworkAddressFromCidr(cidr)

        if err != nil {
            t.Fatalf("error parsing: %s", cidr)
        }

        if expected.String() != actual.String() {
            t.Fatalf("expected[%s] != actual[%s]", expected, actual)
        }
    }
}

func TestBroadcastAddressFromCidr(t *testing.T) {
    service := getService()
    m := map[string]net.IP{
        "165.137.176.0/24": net.IP([]byte{165, 137, 176, 255}),
        "165.137.176.50/24": net.IP([]byte{165, 137, 176, 255}),
        "165.137.176.5/16": net.IP([]byte{165, 137, 255, 255}),
        "165.137.176.255/24": net.IP([]byte{165, 137, 176, 255}),
    }

    for cidr, expected := range m {
        actual, err := service.BroadcastAddressFromCidr(cidr)

        if err != nil {
            t.Fatalf("error parsing: %s", cidr)
        }

        if expected.String() != actual.String() {
            t.Fatalf("expected[%s] != actual[%s]", expected, actual)
        }
    }
}

func TestIncIP(t *testing.T) {
    service := getService()
    type temp struct {
        input net.IP
        next  net.IP
    }

    m := []temp{
        {input: net.IP([]byte{165, 137, 176, 0}), next: net.IP([]byte{165, 137, 176, 1})},
        {input: net.IP([]byte{165, 137, 176, 255}), next: net.IP([]byte{165, 137, 177, 0})},
    }

    for _, x := range m {
        actual := service.IncIP(x.input)

        if x.next.String() != actual.String() {
            t.Fatalf("expected[%s] != actual[%s]", x.next, actual)
        }
    }
}

func TestDecIP(t *testing.T) {
    service := getService()
    type temp struct {
        input net.IP
        prev  net.IP
    }

    m := []temp{
        {input: net.IP([]byte{165, 137, 176, 5}), prev: net.IP([]byte{165, 137, 176, 4})},
        {input: net.IP([]byte{165, 137, 176, 0}), prev: net.IP([]byte{165, 137, 175, 255})},
    }

    for _, x := range m {
        actual := service.DecIP(x.input)

        if x.prev.String() != actual.String() {
            t.Fatalf("expected[%s] != actual[%s]", x.prev, actual)
        }
    }
}

func TestSubnetEnumerate(t *testing.T) {
    service := getService()

    type temp struct {
        cidr      string
        length    int
        network   net.IP
        broadcast net.IP
    }

    m := []temp{
        {
            cidr:"165.137.176.0/24",
            length: 256,
            network: net.IP([]byte{165, 137, 176, 0}),
            broadcast: net.IP([]byte{165, 137, 176, 255}),
        },
        {
            cidr:"165.137.176.0/16",
            length: 65536,
            network: net.IP([]byte{165, 137, 0, 0}),
            broadcast: net.IP([]byte{165, 137, 255, 255}),
        },
    }

    for _, x := range m {
        ips, err := service.SubnetEnumerateFromCidr(x.cidr)

        if err != nil {
            t.Errorf("EnumerateSubnetFromCidr(%s) err: %v", x.cidr, err)
        }

        if x.length != len(ips) {
            t.Errorf("x.length(%s) != len(ips): %v", x.length, len(ips))
        }

        if x.network.String() != ips[0].String() {
            t.Errorf("network not equal [%s] != %s", x.network, ips[0])
        }

        lastIndex := len(ips) - 1

        if x.broadcast.String() != ips[lastIndex].String() {
            t.Errorf("network not equal [%s] != %s", x.network, ips[lastIndex])
        }
    }
}

func TestSubnetEndpoints(t *testing.T) {
    service := getService()

    type temp struct {
        cidr      string
        network   string
        broadcast string
    }

    m := []temp{
        {
            cidr:"165.137.176.0/24",
            network:"165.137.176.0",
            broadcast:"165.137.176.255",
        },
        {
            cidr:"165.137.176.0/16",
            network: "165.137.0.0",
            broadcast: "165.137.255.255",
        },
    }

    for _, x := range m {
        ips, err := service.SubnetEndpointsFromCidr(x.cidr)

        if err != nil {
            t.Errorf("EnumerateSubnetFromCidr(%s) err: %v", x.cidr, err)
        }

        if 2 != len(ips) {
            t.Errorf("x.length(%s) != len(ips): %v", 2, len(ips))
        }

        if x.network != ips[0].String() {
            t.Errorf("network not equal [%s] != %s", x.network, ips[0])
        }

        lastIndex := len(ips) - 1

        if x.broadcast != ips[lastIndex].String() {
            t.Errorf("network not equal [%s] != %s", x.network, ips[lastIndex])
        }
    }
}

func TestDedup(t *testing.T) {
    service := getService()

    input := []string{
        "10.189.4.10",
        "10.189.4.10",
        "10.189.4.254",
        "10.189.4.200",
        "",
    }

    expectedLen := 3
    expected0 := "10.189.4.10"
    expected1 := "10.189.4.254"
    expected2 := "10.189.4.200"

    output := service.dedup(input)

    actualLen := len(output)
    actual0 := output[0]
    actual1 := output[1]
    actual2 := output[2]

    if actualLen != expectedLen {
        t.Errorf("len(output)[%d] doesn't match %d", actualLen, expectedLen)
    }

    if actual0 != expected0 {
        t.Errorf("output[0][%s] doesn't match %s", actual0, expected0)
    }

    if actual1 != expected1 {
        t.Errorf("output[1][%s] doesn't match %s", actual1, expected1)
    }

    if actual2 != expected2 {
        t.Errorf("output[2][%s] doesn't match %s", actual2, expected2)
    }
}
