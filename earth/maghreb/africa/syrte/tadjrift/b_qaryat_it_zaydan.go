package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加里耶特伊斋丹Qaryat_it_zaydanBarony struct {
	feud.BaseBarony
}

var BQaryat_it_zaydan加里耶特伊斋丹 feud.Barony = &加里耶特伊斋丹Qaryat_it_zaydanBarony{}

func init() {
    f := BQaryat_it_zaydan加里耶特伊斋丹.(*加里耶特伊斋丹Qaryat_it_zaydanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaryat_it_zaydan",
		TitleName: "加里耶特伊斋丹",
		TitleCode: "b_qaryat_it_zaydan",
	}
}
