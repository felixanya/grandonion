package agent_plugin

var defaultPluginPaths = []string{"/var/lib/prjname/plugin/", "/var/lib/run/prjname/plugin/"}

type Scanner struct {
    pluginPaths     []string
    errors          []error
}

func NewScanner(path ...string) *Scanner {
    return &Scanner{
        pluginPaths: append(defaultPluginPaths, path...),
        errors: nil,
    }
}

func (s *Scanner) LoadPlugin() error {

    return nil
}

func (s *Scanner) ListAllPlugin() ([]Plugin, error) {

    return nil, nil
}

func (s *Scanner) ListAvailablePlugin() ([]Plugin, error) {

    return nil, nil
}

func (s *Scanner) listPlugin() ([]Plugin, error) {

    return nil, nil
}
