package proxy

import (
    "log"
    "net"
)

type Proxy struct {
    config      *Config

    listener    net.Listener

    quit        chan bool
}

func NewProxy(cfg *Config) *Proxy {
    return &Proxy{
        config: cfg,
        quit: make(chan bool),
    }
}

func (p *Proxy) Start() error {
    var err error
    p.listener, err = net.Listen("tcp", p.config.Addr)
    if err != nil {
        return err
    }

OUTFOR:
    for {
        select {
        case <-p.quit:
            _ = p.listener.Close()
            break OUTFOR
        default:
            conn, err := p.listener.Accept()
            if err != nil {
                log.Println("accept error", err.Error())
                continue
            }

            // TODO
            go p.newConn(conn).serve()
        }

    }

    return nil
}

func (p *Proxy) Close() {
    close(p.quit)
}
