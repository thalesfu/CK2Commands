package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马蒂拉MattilaBarony struct {
	feud.BaseBarony
}

var BMattila马蒂拉 feud.Barony = &马蒂拉MattilaBarony{}

func init() {
    f := BMattila马蒂拉.(*马蒂拉MattilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mattila",
		TitleName: "马蒂拉",
		TitleCode: "b_mattila",
	}
}
