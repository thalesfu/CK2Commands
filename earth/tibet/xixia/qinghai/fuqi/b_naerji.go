package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔几NaerjiBarony struct {
	feud.BaseBarony
}

var BNaerji纳尔几 feud.Barony = &纳尔几NaerjiBarony{}

func init() {
	f := BNaerji纳尔几.(*纳尔几NaerjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naerji",
		TitleName: "纳尔几",
		TitleCode: "b_naerji",
	}
}
