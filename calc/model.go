package calc

import "net"

type Holder struct {
    Low  net.IP
    High net.IP
    Net  net.IPNet
}
