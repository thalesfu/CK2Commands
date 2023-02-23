package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 休鲁阿SiuruanBarony struct {
	feud.BaseBarony
}

var BSiuruan休鲁阿 feud.Barony = &休鲁阿SiuruanBarony{}

func init() {
	f := BSiuruan休鲁阿.(*休鲁阿SiuruanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siuruan",
		TitleName: "休鲁阿",
		TitleCode: "b_siuruan",
	}
}
