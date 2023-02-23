package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利斯马ClysmaBarony struct {
	feud.BaseBarony
}

var BClysma克利斯马 feud.Barony = &克利斯马ClysmaBarony{}

func init() {
	f := BClysma克利斯马.(*克利斯马ClysmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clysma",
		TitleName: "克利斯马",
		TitleCode: "b_clysma",
	}
}
