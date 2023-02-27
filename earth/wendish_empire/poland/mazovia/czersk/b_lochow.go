package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃胡夫LochowBarony struct {
	feud.BaseBarony
}

var BLochow沃胡夫 feud.Barony = &沃胡夫LochowBarony{}

func init() {
    f := BLochow沃胡夫.(*沃胡夫LochowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lochow",
		TitleName: "沃胡夫",
		TitleCode: "b_lochow",
	}
}
