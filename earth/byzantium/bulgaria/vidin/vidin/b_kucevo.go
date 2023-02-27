package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库切沃KucevoBarony struct {
	feud.BaseBarony
}

var BKucevo库切沃 feud.Barony = &库切沃KucevoBarony{}

func init() {
    f := BKucevo库切沃.(*库切沃KucevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kucevo",
		TitleName: "库切沃",
		TitleCode: "b_kucevo",
	}
}
