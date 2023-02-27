package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶拉佩特拉IerapetraBarony struct {
	feud.BaseBarony
}

var BIerapetra耶拉佩特拉 feud.Barony = &耶拉佩特拉IerapetraBarony{}

func init() {
    f := BIerapetra耶拉佩特拉.(*耶拉佩特拉IerapetraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ierapetra",
		TitleName: "耶拉佩特拉",
		TitleCode: "b_ierapetra",
	}
}
