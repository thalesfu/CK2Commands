package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗戈伊RogoiBarony struct {
	feud.BaseBarony
}

var BRogoi罗戈伊 feud.Barony = &罗戈伊RogoiBarony{}

func init() {
    f := BRogoi罗戈伊.(*罗戈伊RogoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rogoi",
		TitleName: "罗戈伊",
		TitleCode: "b_rogoi",
	}
}
