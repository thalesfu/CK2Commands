package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱德伯里LedburyBarony struct {
	feud.BaseBarony
}

var BLedbury莱德伯里 feud.Barony = &莱德伯里LedburyBarony{}

func init() {
	f := BLedbury莱德伯里.(*莱德伯里LedburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ledbury",
		TitleName: "莱德伯里",
		TitleCode: "b_ledbury",
	}
}
