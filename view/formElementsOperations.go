package view

import (
	"time"
	"fmt"
	"errors"
)

func (devView DeviceView) CheckCheckboxByResource(resource string, index int, timeout int) error {
	start := time.Now()
	for {
		current := time.Now()
		delta := current.Sub(start)
		if delta.Seconds() >= float64(timeout) {
			break
		}
		vws, err := devView.GetViewes()
		if err != nil {
			return err
		}
		vw, found := vws.GetByResource(resource, index)
		if found {
			if !vw.Checked {
				return devView.im.TouchScreen.Tap(vw.Center.X, vw.Center.Y)
			}
		}
	}
	return errors.New(fmt.Sprintf("Timeout occured after %d seconds while searching for resource [%s]", timeout, resource))
}

