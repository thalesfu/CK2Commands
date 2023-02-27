package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛霍尔LochoreBarony struct {
	feud.BaseBarony
}

var BLochore洛霍尔 feud.Barony = &洛霍尔LochoreBarony{}

func init() {
    f := BLochore洛霍尔.(*洛霍尔LochoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lochore",
		TitleName: "洛霍尔",
		TitleCode: "b_lochore",
	}
}
