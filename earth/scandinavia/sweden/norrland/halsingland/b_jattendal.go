package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶滕达尔JattendalBarony struct {
	feud.BaseBarony
}

var BJattendal耶滕达尔 feud.Barony = &耶滕达尔JattendalBarony{}

func init() {
	f := BJattendal耶滕达尔.(*耶滕达尔JattendalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jattendal",
		TitleName: "耶滕达尔",
		TitleCode: "b_jattendal",
	}
}
