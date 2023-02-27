package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克赖因堡KrainburgBarony struct {
	feud.BaseBarony
}

var BKrainburg克赖因堡 feud.Barony = &克赖因堡KrainburgBarony{}

func init() {
    f := BKrainburg克赖因堡.(*克赖因堡KrainburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krainburg",
		TitleName: "克赖因堡",
		TitleCode: "b_krainburg",
	}
}
