package main

import (
    "regexp"
)

type DeviceSettings struct {
    // Name.
    Name string `json:"name"`

    // Debug?
    Debug bool `json:"debug"`
}

func (control *Control) Device(settings *DeviceSettings, ok *bool) error {

    rp, err := regexp.Compile(settings.Name)
    if err != nil {
        return err
    }

    for _, device := range control.model.Devices() {
        if rp.MatchString(device.Name()) {
            device.SetDebugging(settings.Debug)
        }
    }

    *ok = true
    return nil
}
