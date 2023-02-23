package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿曼努斯AmanusBarony struct {
	feud.BaseBarony
}

var BAmanus阿曼努斯 feud.Barony = &阿曼努斯AmanusBarony{}

func init() {
	f := BAmanus阿曼努斯.(*阿曼努斯AmanusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amanus",
		TitleName: "阿曼努斯",
		TitleCode: "b_amanus",
	}
}
