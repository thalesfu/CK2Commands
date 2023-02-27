package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡甘布利加CikamburikaBarony struct {
	feud.BaseBarony
}

var BCikamburika锡甘布利加 feud.Barony = &锡甘布利加CikamburikaBarony{}

func init() {
    f := BCikamburika锡甘布利加.(*锡甘布利加CikamburikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cikamburika",
		TitleName: "锡甘布利加",
		TitleCode: "b_cikamburika",
	}
}
