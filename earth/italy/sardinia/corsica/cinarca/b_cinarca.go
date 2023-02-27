package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇纳尔卡CinarcaBarony struct {
	feud.BaseBarony
}

var BCinarca奇纳尔卡 feud.Barony = &奇纳尔卡CinarcaBarony{}

func init() {
    f := BCinarca奇纳尔卡.(*奇纳尔卡CinarcaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cinarca",
		TitleName: "奇纳尔卡",
		TitleCode: "b_cinarca",
	}
}
