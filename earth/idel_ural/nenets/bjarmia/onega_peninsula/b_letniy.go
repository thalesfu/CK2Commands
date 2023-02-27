package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列特尼LetniyBarony struct {
	feud.BaseBarony
}

var BLetniy列特尼 feud.Barony = &列特尼LetniyBarony{}

func init() {
    f := BLetniy列特尼.(*列特尼LetniyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "letniy",
		TitleName: "列特尼",
		TitleCode: "b_letniy",
	}
}
