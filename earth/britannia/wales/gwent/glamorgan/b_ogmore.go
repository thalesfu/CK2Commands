package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥格莫尔OgmoreBarony struct {
	feud.BaseBarony
}

var BOgmore奥格莫尔 feud.Barony = &奥格莫尔OgmoreBarony{}

func init() {
    f := BOgmore奥格莫尔.(*奥格莫尔OgmoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ogmore",
		TitleName: "奥格莫尔",
		TitleCode: "b_ogmore",
	}
}
