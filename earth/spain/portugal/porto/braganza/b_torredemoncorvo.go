package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托里迪蒙科尔武TorredemoncorvoBarony struct {
	feud.BaseBarony
}

var BTorredemoncorvo托里迪蒙科尔武 feud.Barony = &托里迪蒙科尔武TorredemoncorvoBarony{}

func init() {
	f := BTorredemoncorvo托里迪蒙科尔武.(*托里迪蒙科尔武TorredemoncorvoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torredemoncorvo",
		TitleName: "托里迪蒙科尔武",
		TitleCode: "b_torredemoncorvo",
	}
}
