package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 忒耳摩冬ThermodonBarony struct {
	feud.BaseBarony
}

var BThermodon忒耳摩冬 feud.Barony = &忒耳摩冬ThermodonBarony{}

func init() {
	f := BThermodon忒耳摩冬.(*忒耳摩冬ThermodonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thermodon",
		TitleName: "忒耳摩冬",
		TitleCode: "b_thermodon",
	}
}
