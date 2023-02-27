package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基什孔豪洛什KiskunhalasBarony struct {
	feud.BaseBarony
}

var BKiskunhalas基什孔豪洛什 feud.Barony = &基什孔豪洛什KiskunhalasBarony{}

func init() {
    f := BKiskunhalas基什孔豪洛什.(*基什孔豪洛什KiskunhalasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiskunhalas",
		TitleName: "基什孔豪洛什",
		TitleCode: "b_kiskunhalas",
	}
}
