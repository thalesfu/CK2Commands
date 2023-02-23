package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃格尔EgerBarony struct {
	feud.BaseBarony
}

var BEger埃格尔 feud.Barony = &埃格尔EgerBarony{}

func init() {
	f := BEger埃格尔.(*埃格尔EgerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eger",
		TitleName: "埃格尔",
		TitleCode: "b_eger",
	}
}
