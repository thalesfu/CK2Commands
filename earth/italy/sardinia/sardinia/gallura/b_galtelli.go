package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加尔泰利GaltelliBarony struct {
	feud.BaseBarony
}

var BGaltelli加尔泰利 feud.Barony = &加尔泰利GaltelliBarony{}

func init() {
	f := BGaltelli加尔泰利.(*加尔泰利GaltelliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galtelli",
		TitleName: "加尔泰利",
		TitleCode: "b_galtelli",
	}
}
