package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋罗姞利呬BargarhBarony struct {
	feud.BaseBarony
}

var BBargarh跋罗姞利呬 feud.Barony = &跋罗姞利呬BargarhBarony{}

func init() {
    f := BBargarh跋罗姞利呬.(*跋罗姞利呬BargarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bargarh",
		TitleName: "跋罗姞利呬",
		TitleCode: "b_bargarh",
	}
}
