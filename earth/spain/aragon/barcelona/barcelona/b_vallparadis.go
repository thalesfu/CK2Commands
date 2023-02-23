package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴利帕拉迪斯VallparadisBarony struct {
	feud.BaseBarony
}

var BVallparadis巴利帕拉迪斯 feud.Barony = &巴利帕拉迪斯VallparadisBarony{}

func init() {
	f := BVallparadis巴利帕拉迪斯.(*巴利帕拉迪斯VallparadisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vallparadis",
		TitleName: "巴利帕拉迪斯",
		TitleCode: "b_vallparadis",
	}
}
