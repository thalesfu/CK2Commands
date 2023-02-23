package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔达尔ValldalBarony struct {
	feud.BaseBarony
}

var BValldal瓦尔达尔 feud.Barony = &瓦尔达尔ValldalBarony{}

func init() {
	f := BValldal瓦尔达尔.(*瓦尔达尔ValldalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valldal",
		TitleName: "瓦尔达尔",
		TitleCode: "b_valldal",
	}
}
