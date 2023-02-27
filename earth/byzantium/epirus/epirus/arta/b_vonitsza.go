package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尼察VonitszaBarony struct {
	feud.BaseBarony
}

var BVonitsza沃尼察 feud.Barony = &沃尼察VonitszaBarony{}

func init() {
    f := BVonitsza沃尼察.(*沃尼察VonitszaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vonitsza",
		TitleName: "沃尼察",
		TitleCode: "b_vonitsza",
	}
}
