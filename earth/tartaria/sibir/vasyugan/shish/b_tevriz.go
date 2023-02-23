package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷夫里兹TevrizBarony struct {
	feud.BaseBarony
}

var BTevriz捷夫里兹 feud.Barony = &捷夫里兹TevrizBarony{}

func init() {
	f := BTevriz捷夫里兹.(*捷夫里兹TevrizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tevriz",
		TitleName: "捷夫里兹",
		TitleCode: "b_tevriz",
	}
}
