package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅德韦日卡Medvezhka_kuznetskBarony struct {
	feud.BaseBarony
}

var BMedvezhka_kuznetsk梅德韦日卡 feud.Barony = &梅德韦日卡Medvezhka_kuznetskBarony{}

func init() {
    f := BMedvezhka_kuznetsk梅德韦日卡.(*梅德韦日卡Medvezhka_kuznetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medvezhka_kuznetsk",
		TitleName: "梅德韦日卡",
		TitleCode: "b_medvezhka_kuznetsk",
	}
}
