package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图恩ThunBarony struct {
	feud.BaseBarony
}

var BThun图恩 feud.Barony = &图恩ThunBarony{}

func init() {
    f := BThun图恩.(*图恩ThunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thun",
		TitleName: "图恩",
		TitleCode: "b_thun",
	}
}
