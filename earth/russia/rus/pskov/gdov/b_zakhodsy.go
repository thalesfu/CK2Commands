package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎霍齐ZakhodsyBarony struct {
	feud.BaseBarony
}

var BZakhodsy扎霍齐 feud.Barony = &扎霍齐ZakhodsyBarony{}

func init() {
    f := BZakhodsy扎霍齐.(*扎霍齐ZakhodsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zakhodsy",
		TitleName: "扎霍齐",
		TitleCode: "b_zakhodsy",
	}
}
