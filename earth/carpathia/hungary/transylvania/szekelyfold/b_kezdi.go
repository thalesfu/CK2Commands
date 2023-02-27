package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯兹迪KezdiBarony struct {
	feud.BaseBarony
}

var BKezdi凯兹迪 feud.Barony = &凯兹迪KezdiBarony{}

func init() {
    f := BKezdi凯兹迪.(*凯兹迪KezdiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kezdi",
		TitleName: "凯兹迪",
		TitleCode: "b_kezdi",
	}
}
