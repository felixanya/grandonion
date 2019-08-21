package configure

type Producer interface {
    run() error
}

type KProducer struct {
    Hosts       []string    `toml:"hosts"`
    Topic       string      `toml:"topic"`
}

func (kp *KProducer) run() error {

    return nil
}
