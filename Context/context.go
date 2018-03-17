package Context

type MsgContext struct {
	Tpl    string
	Params map[string]interface{}
	Drive  MsgDrive
}

type MsgDrive interface {
	Send(*MsgContext) (string)
}

func (m *MsgContext) Send() (string) {
	return m.Drive.Send(m)
}
