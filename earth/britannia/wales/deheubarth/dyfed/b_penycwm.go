package dyfed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尼库姆PenycwmBarony struct {
	feud.BaseBarony
}

var BPenycwm佩尼库姆 feud.Barony = &佩尼库姆PenycwmBarony{}

func init() {
	f := BPenycwm佩尼库姆.(*佩尼库姆PenycwmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penycwm",
		TitleName: "佩尼库姆",
		TitleCode: "b_penycwm",
	}
}
