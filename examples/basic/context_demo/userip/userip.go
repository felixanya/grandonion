package userip

import (
    "context"
    "fmt"
    "net"
    "net/http"
)

func FromRequest(req *http.Request) (net.IP, error) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        return nil, fmt.Errorf("userip: %s is not IP:Port format", req.RemoteAddr)
    }

    userIp := net.ParseIP(ip)
    if userIp == nil {
        return nil, fmt.Errorf("userip: %s is not IP:Port format", req.RemoteAddr)
    }

    return userIp, nil
}

type key int

const userIpKey key = 0

func NewContext(ctx context.Context, userIp net.IP) context.Context {
    return context.WithValue(ctx, userIpKey, userIp)
}

func FromContext(ctx context.Context) (net.IP, bool) {
    userIp, ok := ctx.Value(userIpKey).(net.IP)
    return userIp, ok
}
