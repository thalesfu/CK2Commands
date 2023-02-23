package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰尼科特RanikotBarony struct {
	feud.BaseBarony
}

var BRanikot兰尼科特 feud.Barony = &兰尼科特RanikotBarony{}

func init() {
	f := BRanikot兰尼科特.(*兰尼科特RanikotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranikot",
		TitleName: "兰尼科特",
		TitleCode: "b_ranikot",
	}
}
