package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古兹达尔QuzdarBarony struct {
	feud.BaseBarony
}

var BQuzdar古兹达尔 feud.Barony = &古兹达尔QuzdarBarony{}

func init() {
	f := BQuzdar古兹达尔.(*古兹达尔QuzdarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quzdar",
		TitleName: "古兹达尔",
		TitleCode: "b_quzdar",
	}
}
