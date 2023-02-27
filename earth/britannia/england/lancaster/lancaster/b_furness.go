package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗内斯FurnessBarony struct {
	feud.BaseBarony
}

var BFurness弗内斯 feud.Barony = &弗内斯FurnessBarony{}

func init() {
    f := BFurness弗内斯.(*弗内斯FurnessBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furness",
		TitleName: "弗内斯",
		TitleCode: "b_furness",
	}
}
