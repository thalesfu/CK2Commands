package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅那古利MainaguriBarony struct {
	feud.BaseBarony
}

var BMainaguri梅那古利 feud.Barony = &梅那古利MainaguriBarony{}

func init() {
	f := BMainaguri梅那古利.(*梅那古利MainaguriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mainaguri",
		TitleName: "梅那古利",
		TitleCode: "b_mainaguri",
	}
}
