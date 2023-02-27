package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科特卡KotkaBarony struct {
	feud.BaseBarony
}

var BKotka科特卡 feud.Barony = &科特卡KotkaBarony{}

func init() {
    f := BKotka科特卡.(*科特卡KotkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotka",
		TitleName: "科特卡",
		TitleCode: "b_kotka",
	}
}
