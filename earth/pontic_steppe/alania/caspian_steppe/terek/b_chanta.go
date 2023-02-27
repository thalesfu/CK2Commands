package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昌塔ChantaBarony struct {
	feud.BaseBarony
}

var BChanta昌塔 feud.Barony = &昌塔ChantaBarony{}

func init() {
    f := BChanta昌塔.(*昌塔ChantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chanta",
		TitleName: "昌塔",
		TitleCode: "b_chanta",
	}
}
