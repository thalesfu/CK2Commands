package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰迪尤LlanddewBarony struct {
	feud.BaseBarony
}

var BLlanddew兰迪尤 feud.Barony = &兰迪尤LlanddewBarony{}

func init() {
	f := BLlanddew兰迪尤.(*兰迪尤LlanddewBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanddew",
		TitleName: "兰迪尤",
		TitleCode: "b_llanddew",
	}
}
