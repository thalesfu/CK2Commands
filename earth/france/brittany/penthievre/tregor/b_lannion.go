package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗尼永LannionBarony struct {
	feud.BaseBarony
}

var BLannion朗尼永 feud.Barony = &朗尼永LannionBarony{}

func init() {
    f := BLannion朗尼永.(*朗尼永LannionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lannion",
		TitleName: "朗尼永",
		TitleCode: "b_lannion",
	}
}
