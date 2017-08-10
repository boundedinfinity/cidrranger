package calc

import "strings"

func (this *CalculatorService) ignoreLine(line string) bool {
    if line == "" {
        return true
    }

    if strings.HasPrefix(line, "#") {
        return true
    }

    if strings.HasPrefix(line, "//") {
        return true
    }

    return false
}

func (this *CalculatorService) dedup(raw []string) []string {
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
