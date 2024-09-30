package switchobject

import "github.com/Netsocs-Team/driver.sdk_go/pkg/objects"

type Switch interface {
	objects.ObjectChild
	TurnOn()
	TurnOff()
	Toggle()
	IsOn() bool
}

type SwitchObject struct {
	isOn bool
}

func (s *SwitchObject) TurnOn() {
	s.isOn = true
}

func (s *SwitchObject) TurnOff() {
	s.isOn = false
}

func (s *SwitchObject) IsOn() bool {
	return s.isOn
}

func (s *SwitchObject) Toggle() {
	s.isOn = !s.isOn
}

func (s *SwitchObject) ObjectChild() objects.Object {

	return objects.Object{
		ObjectID:        "switch",
		Name:            "SwitchNormal",
		Type:            "switch",
		State:           -1,
		StatesAvailable: []objects.StateObject{},
		StateAttributes: objects.StateAttributes{},
		DeviceID:        0,
		Enabled:         false,
		Icon:            "",
	}
}
