package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔旺根EllwangenBarony struct {
	feud.BaseBarony
}

var BEllwangen埃尔旺根 feud.Barony = &埃尔旺根EllwangenBarony{}

func init() {
    f := BEllwangen埃尔旺根.(*埃尔旺根EllwangenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ellwangen",
		TitleName: "埃尔旺根",
		TitleCode: "b_ellwangen",
	}
}
