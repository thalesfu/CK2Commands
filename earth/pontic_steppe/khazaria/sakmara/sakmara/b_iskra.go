package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯克拉IskraBarony struct {
	feud.BaseBarony
}

var BIskra伊斯克拉 feud.Barony = &伊斯克拉IskraBarony{}

func init() {
    f := BIskra伊斯克拉.(*伊斯克拉IskraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iskra",
		TitleName: "伊斯克拉",
		TitleCode: "b_iskra",
	}
}
