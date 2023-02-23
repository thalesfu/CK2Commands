package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡希纳CascinaBarony struct {
	feud.BaseBarony
}

var BCascina卡希纳 feud.Barony = &卡希纳CascinaBarony{}

func init() {
	f := BCascina卡希纳.(*卡希纳CascinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cascina",
		TitleName: "卡希纳",
		TitleCode: "b_cascina",
	}
}
