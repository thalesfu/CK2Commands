package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布若BurogBarony struct {
	feud.BaseBarony
}

var BBurog布若 feud.Barony = &布若BurogBarony{}

func init() {
	f := BBurog布若.(*布若BurogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burog",
		TitleName: "布若",
		TitleCode: "b_burog",
	}
}
