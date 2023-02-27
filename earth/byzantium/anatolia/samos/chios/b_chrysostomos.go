package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫里索斯托莫斯ChrysostomosBarony struct {
	feud.BaseBarony
}

var BChrysostomos赫里索斯托莫斯 feud.Barony = &赫里索斯托莫斯ChrysostomosBarony{}

func init() {
    f := BChrysostomos赫里索斯托莫斯.(*赫里索斯托莫斯ChrysostomosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chrysostomos",
		TitleName: "赫里索斯托莫斯",
		TitleCode: "b_chrysostomos",
	}
}
