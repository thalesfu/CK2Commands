package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔科洛ArkoloBarony struct {
	feud.BaseBarony
}

var BArkolo阿尔科洛 feud.Barony = &阿尔科洛ArkoloBarony{}

func init() {
    f := BArkolo阿尔科洛.(*阿尔科洛ArkoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkolo",
		TitleName: "阿尔科洛",
		TitleCode: "b_arkolo",
	}
}
