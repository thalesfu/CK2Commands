package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丹丹乌依里克Dandan_uilikBarony struct {
	feud.BaseBarony
}

var BDandan_uilik丹丹乌依里克 feud.Barony = &丹丹乌依里克Dandan_uilikBarony{}

func init() {
    f := BDandan_uilik丹丹乌依里克.(*丹丹乌依里克Dandan_uilikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dandan_uilik",
		TitleName: "丹丹乌依里克",
		TitleCode: "b_dandan_uilik",
	}
}
