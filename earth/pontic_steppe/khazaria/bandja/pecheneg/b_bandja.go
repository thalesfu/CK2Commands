package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班贾BandjaBarony struct {
	feud.BaseBarony
}

var BBandja班贾 feud.Barony = &班贾BandjaBarony{}

func init() {
    f := BBandja班贾.(*班贾BandjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandja",
		TitleName: "班贾",
		TitleCode: "b_bandja",
	}
}
