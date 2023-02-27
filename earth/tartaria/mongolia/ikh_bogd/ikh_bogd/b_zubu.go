package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阻卜ZubuBarony struct {
	feud.BaseBarony
}

var BZubu阻卜 feud.Barony = &阻卜ZubuBarony{}

func init() {
    f := BZubu阻卜.(*阻卜ZubuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zubu",
		TitleName: "阻卜",
		TitleCode: "b_zubu",
	}
}
