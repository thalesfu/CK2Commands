package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑斯廷斯HastingsBarony struct {
	feud.BaseBarony
}

var BHastings黑斯廷斯 feud.Barony = &黑斯廷斯HastingsBarony{}

func init() {
    f := BHastings黑斯廷斯.(*黑斯廷斯HastingsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hastings",
		TitleName: "黑斯廷斯",
		TitleCode: "b_hastings",
	}
}
