package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯罕IshanBarony struct {
	feud.BaseBarony
}

var BIshan伊斯罕 feud.Barony = &伊斯罕IshanBarony{}

func init() {
    f := BIshan伊斯罕.(*伊斯罕IshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ishan",
		TitleName: "伊斯罕",
		TitleCode: "b_ishan",
	}
}
