package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑弗灵HaveringBarony struct {
	feud.BaseBarony
}

var BHavering黑弗灵 feud.Barony = &黑弗灵HaveringBarony{}

func init() {
	f := BHavering黑弗灵.(*黑弗灵HaveringBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "havering",
		TitleName: "黑弗灵",
		TitleCode: "b_havering",
	}
}
