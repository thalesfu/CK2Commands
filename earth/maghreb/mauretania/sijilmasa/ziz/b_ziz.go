package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济兹ZizBarony struct {
	feud.BaseBarony
}

var BZiz济兹 feud.Barony = &济兹ZizBarony{}

func init() {
    f := BZiz济兹.(*济兹ZizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ziz",
		TitleName: "济兹",
		TitleCode: "b_ziz",
	}
}
