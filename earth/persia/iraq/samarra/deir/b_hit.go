package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希特HitBarony struct {
	feud.BaseBarony
}

var BHit希特 feud.Barony = &希特HitBarony{}

func init() {
	f := BHit希特.(*希特HitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hit",
		TitleName: "希特",
		TitleCode: "b_hit",
	}
}
