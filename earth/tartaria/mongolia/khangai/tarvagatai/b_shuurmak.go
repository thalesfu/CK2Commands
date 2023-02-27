package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒尔马克ShuurmakBarony struct {
	feud.BaseBarony
}

var BShuurmak舒尔马克 feud.Barony = &舒尔马克ShuurmakBarony{}

func init() {
    f := BShuurmak舒尔马克.(*舒尔马克ShuurmakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shuurmak",
		TitleName: "舒尔马克",
		TitleCode: "b_shuurmak",
	}
}
