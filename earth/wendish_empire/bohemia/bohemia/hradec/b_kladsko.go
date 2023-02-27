package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克沃兹科KladskoBarony struct {
	feud.BaseBarony
}

var BKladsko克沃兹科 feud.Barony = &克沃兹科KladskoBarony{}

func init() {
    f := BKladsko克沃兹科.(*克沃兹科KladskoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kladsko",
		TitleName: "克沃兹科",
		TitleCode: "b_kladsko",
	}
}
