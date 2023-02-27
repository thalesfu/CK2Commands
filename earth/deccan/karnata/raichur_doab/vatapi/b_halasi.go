package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃罗斯HalasiBarony struct {
	feud.BaseBarony
}

var BHalasi诃罗斯 feud.Barony = &诃罗斯HalasiBarony{}

func init() {
    f := BHalasi诃罗斯.(*诃罗斯HalasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halasi",
		TitleName: "诃罗斯",
		TitleCode: "b_halasi",
	}
}
