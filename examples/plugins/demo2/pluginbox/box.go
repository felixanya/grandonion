package pluginbox

var PluginBoxes map[string]PluginBox

type PluginBox interface {
    Flag() bool

    Execute(in interface{}) (out , error)
}
