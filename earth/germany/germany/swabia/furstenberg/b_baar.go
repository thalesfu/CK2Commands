package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔BaarBarony struct {
	feud.BaseBarony
}

var BBaar巴尔 feud.Barony = &巴尔BaarBarony{}

func init() {
	f := BBaar巴尔.(*巴尔BaarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baar",
		TitleName: "巴尔",
		TitleCode: "b_baar",
	}
}
