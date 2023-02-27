package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯特林KetteringBarony struct {
	feud.BaseBarony
}

var BKettering凯特林 feud.Barony = &凯特林KetteringBarony{}

func init() {
    f := BKettering凯特林.(*凯特林KetteringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kettering",
		TitleName: "凯特林",
		TitleCode: "b_kettering",
	}
}
