package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔加斯OrgazBarony struct {
	feud.BaseBarony
}

var BOrgaz奥尔加斯 feud.Barony = &奥尔加斯OrgazBarony{}

func init() {
	f := BOrgaz奥尔加斯.(*奥尔加斯OrgazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orgaz",
		TitleName: "奥尔加斯",
		TitleCode: "b_orgaz",
	}
}
