package main

import (
    "fmt"
)

const (
    minute float64 = 1.0
    hour   float64 = 60 * minute
    day    float64 = 24 * 60 * minute

    normalIsolationTime      = 2 * day
    emergencyIsolationTime   = 2 * hour
    maxRetentionTime         = 7 * day
    defaultWaitTimeout       = 10 * minute

    waitLLDPUploadTimeout    = defaultWaitTimeout
    waitRebootTimeout        = defaultWaitTimeout
    waitOsCleanResultTimeout = defaultWaitTimeout
)

func main() {
    fmt.Printf("%.2f\n", minute)
    fmt.Printf("%.2f\n", hour)
    fmt.Printf("%.2f\n", day)

    fmt.Printf("%.2f\n", maxRetentionTime)

    fmt.Printf("%t\n", isValid())

}

func isValid() bool {
    return normalIsolationTime >= emergencyIsolationTime
}
