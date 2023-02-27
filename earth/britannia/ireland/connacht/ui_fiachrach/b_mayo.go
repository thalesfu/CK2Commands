package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅奥MayoBarony struct {
	feud.BaseBarony
}

var BMayo梅奥 feud.Barony = &梅奥MayoBarony{}

func init() {
    f := BMayo梅奥.(*梅奥MayoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayo",
		TitleName: "梅奥",
		TitleCode: "b_mayo",
	}
}
