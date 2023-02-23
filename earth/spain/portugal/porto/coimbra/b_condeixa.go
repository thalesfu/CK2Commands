package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔代沙CondeixaBarony struct {
	feud.BaseBarony
}

var BCondeixa孔代沙 feud.Barony = &孔代沙CondeixaBarony{}

func init() {
	f := BCondeixa孔代沙.(*孔代沙CondeixaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "condeixa",
		TitleName: "孔代沙",
		TitleCode: "b_condeixa",
	}
}
