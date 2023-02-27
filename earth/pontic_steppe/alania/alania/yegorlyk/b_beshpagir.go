package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别什帕吉尔BeshpagirBarony struct {
	feud.BaseBarony
}

var BBeshpagir别什帕吉尔 feud.Barony = &别什帕吉尔BeshpagirBarony{}

func init() {
    f := BBeshpagir别什帕吉尔.(*别什帕吉尔BeshpagirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beshpagir",
		TitleName: "别什帕吉尔",
		TitleCode: "b_beshpagir",
	}
}
