package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特里日基StrizhkiBarony struct {
	feud.BaseBarony
}

var BStrizhki斯特里日基 feud.Barony = &斯特里日基StrizhkiBarony{}

func init() {
    f := BStrizhki斯特里日基.(*斯特里日基StrizhkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strizhki",
		TitleName: "斯特里日基",
		TitleCode: "b_strizhki",
	}
}
