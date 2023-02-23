package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄂特冈腾格里OtgontengerBarony struct {
	feud.BaseBarony
}

var BOtgontenger鄂特冈腾格里 feud.Barony = &鄂特冈腾格里OtgontengerBarony{}

func init() {
	f := BOtgontenger鄂特冈腾格里.(*鄂特冈腾格里OtgontengerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otgontenger",
		TitleName: "鄂特冈腾格里",
		TitleCode: "b_otgontenger",
	}
}
