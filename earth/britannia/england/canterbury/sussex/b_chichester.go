package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇切斯特ChichesterBarony struct {
	feud.BaseBarony
}

var BChichester奇切斯特 feud.Barony = &奇切斯特ChichesterBarony{}

func init() {
    f := BChichester奇切斯特.(*奇切斯特ChichesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chichester",
		TitleName: "奇切斯特",
		TitleCode: "b_chichester",
	}
}
