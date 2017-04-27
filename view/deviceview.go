// TODO : Documentation

package view

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antonienko/goandroid/device"
	"github.com/antonienko/goandroid/display"
	"github.com/antonienko/goandroid/input"
	"strings"
)

type DeviceView struct {
	dev  device.Device
	im   input.InputManager
	disp display.Display
}

func NewDeviceView(dev device.Device) DeviceView {
	im := input.NewInputManager(dev)
	return DeviceView{dev: dev, im: im}
}

func (devView DeviceView) GetViewes() (Views, error) {
	hierarchy, err := devView.GetHierarchy()
	if err != nil {
		return Views{}, err
	}
	return hierarchy.ConvertToViews()
}

func (devView DeviceView) GetHierarchy() (Hierarchy, error) {
	var tag = "UI hierchary dumped to:"
	retriesLeft := 10
	out, _ := devView.makeDump()
	for !strings.Contains(out, tag) {
		retriesLeft--
		if !retriesLeft {
			return Hierarchy{}, errors.New(fmt.Sprintf("Unable to locate [%s] in output : %s", tag, out))
		}
		out, _ = devView.makeDump()
	}
	parts := strings.Split(out, "dumped to:")
	if len(parts) != 2 {
		return Hierarchy{}, errors.New(fmt.Sprintf("Unable to locate file location in output : %s", out))
	}
	xml_location := parts[1]
	xml_location = strings.TrimSpace(xml_location)

	xml_data, err := devView.dev.Shell("cat", xml_location)
	if err != nil {
		return Hierarchy{}, err
	}
	xml_data = strings.TrimSpace(xml_data)
	xml_hierarchy := new(Hierarchy)
	err = xml.Unmarshal([]byte(xml_data), xml_hierarchy)
	if err != nil {
		return Hierarchy{}, err
	}

	return *xml_hierarchy, nil
}

func (devView DeviceView) makeDump() (string,error) {
	out, err := devView.dev.Shell("uiautomator dump /data/window_dump.xml")
	if err != nil {
		return Hierarchy{}, err
	}
	return out,nil
}
