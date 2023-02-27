package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜尔祖基亚DurdzukiaBarony struct {
	feud.BaseBarony
}

var BDurdzukia杜尔祖基亚 feud.Barony = &杜尔祖基亚DurdzukiaBarony{}

func init() {
    f := BDurdzukia杜尔祖基亚.(*杜尔祖基亚DurdzukiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durdzukia",
		TitleName: "杜尔祖基亚",
		TitleCode: "b_durdzukia",
	}
}
