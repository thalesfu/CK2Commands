package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃德胡EdkhouBarony struct {
	feud.BaseBarony
}

var BEdkhou埃德胡 feud.Barony = &埃德胡EdkhouBarony{}

func init() {
	f := BEdkhou埃德胡.(*埃德胡EdkhouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edkhou",
		TitleName: "埃德胡",
		TitleCode: "b_edkhou",
	}
}
