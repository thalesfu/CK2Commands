package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃伊尔EgilBarony struct {
	feud.BaseBarony
}

var BEgil埃伊尔 feud.Barony = &埃伊尔EgilBarony{}

func init() {
	f := BEgil埃伊尔.(*埃伊尔EgilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egil",
		TitleName: "埃伊尔",
		TitleCode: "b_egil",
	}
}
