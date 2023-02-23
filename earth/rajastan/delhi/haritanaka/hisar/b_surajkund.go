package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏罗阇君荼SurajkundBarony struct {
	feud.BaseBarony
}

var BSurajkund苏罗阇君荼 feud.Barony = &苏罗阇君荼SurajkundBarony{}

func init() {
	f := BSurajkund苏罗阇君荼.(*苏罗阇君荼SurajkundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surajkund",
		TitleName: "苏罗阇君荼",
		TitleCode: "b_surajkund",
	}
}
