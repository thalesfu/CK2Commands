package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丹吉尔TangiersBarony struct {
	feud.BaseBarony
}

var BTangiers丹吉尔 feud.Barony = &丹吉尔TangiersBarony{}

func init() {
	f := BTangiers丹吉尔.(*丹吉尔TangiersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tangiers",
		TitleName: "丹吉尔",
		TitleCode: "b_tangiers",
	}
}
