package calc

import (
    "testing"
    "net"
)

func TestDedup(t *testing.T) {
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

    output := dedup(input)

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

func TestHighlow(t *testing.T) {
    input := []net.IP{
        net.ParseIP("10.189.4.10"),
        net.ParseIP("10.189.4.1"),
        net.ParseIP("10.189.4.254"),
        net.ParseIP("10.189.4.200"),
    }

    hl := highlow(input)

    if len(hl) != 2 {
        t.Errorf("len(hl)[%d] doesn't match 2", 2)
    }

    if hl[0].String() != "10.189.4.1" {
        t.Errorf("hl[0][%s] doesn't match 10.189.4.1", hl[0])
    }

    if hl[1].String() != "10.189.4.254" {
        t.Errorf("hl[1][%s] doesn't match 10.189.4.254", hl[1])
    }
}

func TestIprange(t *testing.T) {
    iprange, err := iprange("10.189.4.0/24")

    if err != nil {
        t.Errorf("iprange(10.189.4.0/24) err", err.Error())
    }

    iplen := len(iprange)
    lastIndex := iplen - 1

    if len(iprange) != 254 {
        t.Errorf("len(iprange)[%d] doesn't match 254", iplen)
    }

    if iprange[0].String() != "10.189.4.1" {
        t.Errorf("iprange[0][%s] doesn't match 10.189.4.1", iprange[0])
    }

    if iprange[lastIndex].String() != "10.189.4.254" {
        t.Errorf("iprange[%d][%s] doesn't match 10.189.4.254", lastIndex, iprange[lastIndex])
    }
}

func TestInc(t *testing.T) {
    netIp := net.ParseIP("10.189.4.1")
    nextIp := inc(netIp)

    if nextIp.String() != "10.189.4.2" {
        t.Errorf("nextIp[%s] doesn't match 10.189.4.2", netIp)
    }
}

func TestIp2Int(t *testing.T) {
    netIp := net.ParseIP("10.189.4.1")
    intIp := ip2int(netIp)

    if intIp != uint32(180159489) {
        t.Errorf("intIp[%d] doesn't match 180159489", intIp)
        t.Fail()
    }
}

func TestInt2Ip(t *testing.T) {
    input := uint32(180159489)
    netIp := int2ip(input)

    if netIp.String() != "10.189.4.1" {
        t.Errorf("netIp[%s] doesn't match 10.189.4.2", netIp)
    }
}
