package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅梅尔MemelBarony struct {
	feud.BaseBarony
}

var BMemel梅梅尔 feud.Barony = &梅梅尔MemelBarony{}

func init() {
    f := BMemel梅梅尔.(*梅梅尔MemelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "memel",
		TitleName: "梅梅尔",
		TitleCode: "b_memel",
	}
}
