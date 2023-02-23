package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈里穆达MerimdaBarony struct {
	feud.BaseBarony
}

var BMerimda迈里穆达 feud.Barony = &迈里穆达MerimdaBarony{}

func init() {
	f := BMerimda迈里穆达.(*迈里穆达MerimdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merimda",
		TitleName: "迈里穆达",
		TitleCode: "b_merimda",
	}
}
