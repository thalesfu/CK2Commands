package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尼拉帕茨DonjilapacBarony struct {
	feud.BaseBarony
}

var BDonjilapac多尼拉帕茨 feud.Barony = &多尼拉帕茨DonjilapacBarony{}

func init() {
    f := BDonjilapac多尼拉帕茨.(*多尼拉帕茨DonjilapacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donjilapac",
		TitleName: "多尼拉帕茨",
		TitleCode: "b_donjilapac",
	}
}
