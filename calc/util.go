package calc

import (
    "net"
    "encoding/binary"
    "math"
    "strings"
    "fmt"
)

func cidr() net.IPNet {
    maxLen := 32
    for l := maxLen; l >= 0; l-- {
        mask := net.CIDRMask(l, maxLen)
        na := ip1.Mask(mask)
        n := net.IPNet{IP: na, Mask: mask}

        if n.Contains(ip2) {
            fmt.Printf("smallest possible CIDR range: %v/%v\n", na, l)
            break
        }
    }
}

func dedup(raw []string) []string {
    dmap := make(map[string]bool)
    dedup := make([]string, 0)

    for _, item := range raw {
        titem := strings.TrimSpace(item)

        if titem == "" {
            continue
        }

        if _, found := dmap[titem]; !found {
            dmap[titem] = true
        }
    }

    for k, _ := range dmap {
        dedup = append(dedup, k)
    }

    return dedup
}

func highlow(iprange []net.IP) []net.IP {
    lowestint := uint32(math.MaxUint32)
    highestint := uint32(0)
    ipints := make([]uint32, 0)

    for _, item := range iprange {
        ipints = append(ipints, ip2int(item))
    }

    for _, item := range ipints {
        if item < lowestint {
            lowestint = item
        }

        if item > highestint {
            highestint = item
        }
    }

    highest := int2ip(highestint)
    lowest := int2ip(lowestint)

    return []net.IP{lowest, highest}
}

// https://gist.github.com/kotakanbe/d3059af990252ba89a82
func iprange(cidr string) ([]net.IP, error) {
    ip, ipnet, err := net.ParseCIDR(cidr)
    if err != nil {
        return nil, err
    }

    var ips []net.IP

    for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ip = inc(ip) {
        ips = append(ips, ip)
    }

    // remove network address and broadcast address
    return ips[1 : len(ips) - 1], nil
}

// https://gist.github.com/kotakanbe/d3059af990252ba89a82
func inc(ip net.IP) net.IP {
    newIp := net.ParseIP(ip.String())

    for j := len(newIp) - 1; j >= 0; j-- {
        newIp[j]++
        if newIp[j] > 0 {
            break
        }
    }

    return newIp
}

// https://gist.github.com/ammario/649d4c0da650162efd404af23e25b86b
func ip2int(ip net.IP) uint32 {
    if len(ip) == 16 {
        return binary.BigEndian.Uint32(ip[12:16])
    }
    return binary.BigEndian.Uint32(ip)
}

func int2ip(nn uint32) net.IP {
    ip := make(net.IP, 4)
    binary.BigEndian.PutUint32(ip, nn)
    return ip
}
