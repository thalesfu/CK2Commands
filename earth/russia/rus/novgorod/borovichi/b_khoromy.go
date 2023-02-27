package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗梅KhoromyBarony struct {
	feud.BaseBarony
}

var BKhoromy霍罗梅 feud.Barony = &霍罗梅KhoromyBarony{}

func init() {
    f := BKhoromy霍罗梅.(*霍罗梅KhoromyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khoromy",
		TitleName: "霍罗梅",
		TitleCode: "b_khoromy",
	}
}
