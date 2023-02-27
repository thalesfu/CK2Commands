package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米拉萨MylasaBarony struct {
	feud.BaseBarony
}

var BMylasa米拉萨 feud.Barony = &米拉萨MylasaBarony{}

func init() {
    f := BMylasa米拉萨.(*米拉萨MylasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mylasa",
		TitleName: "米拉萨",
		TitleCode: "b_mylasa",
	}
}
