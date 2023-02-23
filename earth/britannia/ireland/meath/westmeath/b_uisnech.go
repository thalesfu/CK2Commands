package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊什纳赫UisnechBarony struct {
	feud.BaseBarony
}

var BUisnech伊什纳赫 feud.Barony = &伊什纳赫UisnechBarony{}

func init() {
	f := BUisnech伊什纳赫.(*伊什纳赫UisnechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uisnech",
		TitleName: "伊什纳赫",
		TitleCode: "b_uisnech",
	}
}
