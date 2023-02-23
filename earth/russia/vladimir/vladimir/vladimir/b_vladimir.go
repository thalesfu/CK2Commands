package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉基米尔VladimirBarony struct {
	feud.BaseBarony
}

var BVladimir弗拉基米尔 feud.Barony = &弗拉基米尔VladimirBarony{}

func init() {
	f := BVladimir弗拉基米尔.(*弗拉基米尔VladimirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vladimir",
		TitleName: "弗拉基米尔",
		TitleCode: "b_vladimir",
	}
}
