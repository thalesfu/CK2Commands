package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥辛尼基OsinnikiBarony struct {
	feud.BaseBarony
}

var BOsinniki奥辛尼基 feud.Barony = &奥辛尼基OsinnikiBarony{}

func init() {
	f := BOsinniki奥辛尼基.(*奥辛尼基OsinnikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osinniki",
		TitleName: "奥辛尼基",
		TitleCode: "b_osinniki",
	}
}
