package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔霍夫VolkhovBarony struct {
	feud.BaseBarony
}

var BVolkhov沃尔霍夫 feud.Barony = &沃尔霍夫VolkhovBarony{}

func init() {
    f := BVolkhov沃尔霍夫.(*沃尔霍夫VolkhovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volkhov",
		TitleName: "沃尔霍夫",
		TitleCode: "b_volkhov",
	}
}
