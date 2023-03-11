package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维戈尼奇VygonichiBarony struct {
	feud.BaseBarony
}

var BVygonichi维戈尼奇 feud.Barony = &维戈尼奇VygonichiBarony{}

func init() {
    f := BVygonichi维戈尼奇.(*维戈尼奇VygonichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vygonichi",
		TitleName: "维戈尼奇",
		TitleCode: "b_vygonichi",
	}
}
