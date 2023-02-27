package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉加尔HangalBarony struct {
	feud.BaseBarony
}

var BHangal汉加尔 feud.Barony = &汉加尔HangalBarony{}

func init() {
    f := BHangal汉加尔.(*汉加尔HangalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hangal",
		TitleName: "汉加尔",
		TitleCode: "b_hangal",
	}
}
