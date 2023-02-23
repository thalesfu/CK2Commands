package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉雄LaxongBarony struct {
	feud.BaseBarony
}

var BLaxong拉雄 feud.Barony = &拉雄LaxongBarony{}

func init() {
	f := BLaxong拉雄.(*拉雄LaxongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laxong",
		TitleName: "拉雄",
		TitleCode: "b_laxong",
	}
}
