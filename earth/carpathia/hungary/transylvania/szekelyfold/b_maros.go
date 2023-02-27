package szekelyfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛罗什MarosBarony struct {
	feud.BaseBarony
}

var BMaros毛罗什 feud.Barony = &毛罗什MarosBarony{}

func init() {
    f := BMaros毛罗什.(*毛罗什MarosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maros",
		TitleName: "毛罗什",
		TitleCode: "b_maros",
	}
}
