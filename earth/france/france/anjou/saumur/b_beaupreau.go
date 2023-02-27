package saumur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博普雷欧BeaupreauBarony struct {
	feud.BaseBarony
}

var BBeaupreau博普雷欧 feud.Barony = &博普雷欧BeaupreauBarony{}

func init() {
    f := BBeaupreau博普雷欧.(*博普雷欧BeaupreauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaupreau",
		TitleName: "博普雷欧",
		TitleCode: "b_beaupreau",
	}
}
