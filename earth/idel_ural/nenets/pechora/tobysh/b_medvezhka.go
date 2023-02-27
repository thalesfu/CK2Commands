package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅德韦日卡MedvezhkaBarony struct {
	feud.BaseBarony
}

var BMedvezhka梅德韦日卡 feud.Barony = &梅德韦日卡MedvezhkaBarony{}

func init() {
    f := BMedvezhka梅德韦日卡.(*梅德韦日卡MedvezhkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medvezhka",
		TitleName: "梅德韦日卡",
		TitleCode: "b_medvezhka",
	}
}
