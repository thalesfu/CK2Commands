package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里斯德拉LisdraBarony struct {
	feud.BaseBarony
}

var BLisdra里斯德拉 feud.Barony = &里斯德拉LisdraBarony{}

func init() {
    f := BLisdra里斯德拉.(*里斯德拉LisdraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lisdra",
		TitleName: "里斯德拉",
		TitleCode: "b_lisdra",
	}
}
