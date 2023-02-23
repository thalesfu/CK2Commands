package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福拉ForradhBarony struct {
	feud.BaseBarony
}

var BForradh福拉 feud.Barony = &福拉ForradhBarony{}

func init() {
	f := BForradh福拉.(*福拉ForradhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forradh",
		TitleName: "福拉",
		TitleCode: "b_forradh",
	}
}
