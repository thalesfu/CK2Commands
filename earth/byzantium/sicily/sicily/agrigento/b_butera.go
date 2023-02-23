package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布泰拉ButeraBarony struct {
	feud.BaseBarony
}

var BButera布泰拉 feud.Barony = &布泰拉ButeraBarony{}

func init() {
	f := BButera布泰拉.(*布泰拉ButeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butera",
		TitleName: "布泰拉",
		TitleCode: "b_butera",
	}
}
