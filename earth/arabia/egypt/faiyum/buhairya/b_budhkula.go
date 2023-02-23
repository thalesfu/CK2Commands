package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赫库拉BudhkulaBarony struct {
	feud.BaseBarony
}

var BBudhkula布赫库拉 feud.Barony = &布赫库拉BudhkulaBarony{}

func init() {
	f := BBudhkula布赫库拉.(*布赫库拉BudhkulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "budhkula",
		TitleName: "布赫库拉",
		TitleCode: "b_budhkula",
	}
}
