package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基比拉KibyraBarony struct {
	feud.BaseBarony
}

var BKibyra基比拉 feud.Barony = &基比拉KibyraBarony{}

func init() {
    f := BKibyra基比拉.(*基比拉KibyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kibyra",
		TitleName: "基比拉",
		TitleCode: "b_kibyra",
	}
}
