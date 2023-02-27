package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大维索奇诺耶VelikovisochnoyeBarony struct {
	feud.BaseBarony
}

var BVelikovisochnoye大维索奇诺耶 feud.Barony = &大维索奇诺耶VelikovisochnoyeBarony{}

func init() {
    f := BVelikovisochnoye大维索奇诺耶.(*大维索奇诺耶VelikovisochnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velikovisochnoye",
		TitleName: "大维索奇诺耶",
		TitleCode: "b_velikovisochnoye",
	}
}
