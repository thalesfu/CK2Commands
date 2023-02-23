package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图姆布拉ToumboulaBarony struct {
	feud.BaseBarony
}

var BToumboula图姆布拉 feud.Barony = &图姆布拉ToumboulaBarony{}

func init() {
	f := BToumboula图姆布拉.(*图姆布拉ToumboulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toumboula",
		TitleName: "图姆布拉",
		TitleCode: "b_toumboula",
	}
}
