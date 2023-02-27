package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯图波维StupoviBarony struct {
	feud.BaseBarony
}

var BStupovi斯图波维 feud.Barony = &斯图波维StupoviBarony{}

func init() {
    f := BStupovi斯图波维.(*斯图波维StupoviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stupovi",
		TitleName: "斯图波维",
		TitleCode: "b_stupovi",
	}
}
