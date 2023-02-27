package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰基LanjiBarony struct {
	feud.BaseBarony
}

var BLanji兰基 feud.Barony = &兰基LanjiBarony{}

func init() {
    f := BLanji兰基.(*兰基LanjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lanji",
		TitleName: "兰基",
		TitleCode: "b_lanji",
	}
}
