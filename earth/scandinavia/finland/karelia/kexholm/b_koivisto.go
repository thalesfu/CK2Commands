package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科伊维斯托KoivistoBarony struct {
	feud.BaseBarony
}

var BKoivisto科伊维斯托 feud.Barony = &科伊维斯托KoivistoBarony{}

func init() {
    f := BKoivisto科伊维斯托.(*科伊维斯托KoivistoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koivisto",
		TitleName: "科伊维斯托",
		TitleCode: "b_koivisto",
	}
}
