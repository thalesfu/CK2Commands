package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维扎斯VizhasBarony struct {
	feud.BaseBarony
}

var BVizhas维扎斯 feud.Barony = &维扎斯VizhasBarony{}

func init() {
    f := BVizhas维扎斯.(*维扎斯VizhasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vizhas",
		TitleName: "维扎斯",
		TitleCode: "b_vizhas",
	}
}
