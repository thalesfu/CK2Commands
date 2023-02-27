package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因塔戈IttagiBarony struct {
	feud.BaseBarony
}

var BIttagi因塔戈 feud.Barony = &因塔戈IttagiBarony{}

func init() {
    f := BIttagi因塔戈.(*因塔戈IttagiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ittagi",
		TitleName: "因塔戈",
		TitleCode: "b_ittagi",
	}
}
