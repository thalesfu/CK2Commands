package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣吉扬_勒_代赛尔SaintguilhemledesertBarony struct {
	feud.BaseBarony
}

var BSaintguilhemledesert圣吉扬_勒_代赛尔 feud.Barony = &圣吉扬_勒_代赛尔SaintguilhemledesertBarony{}

func init() {
	f := BSaintguilhemledesert圣吉扬_勒_代赛尔.(*圣吉扬_勒_代赛尔SaintguilhemledesertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintguilhemledesert",
		TitleName: "圣吉扬_勒_代赛尔",
		TitleCode: "b_saintguilhemledesert",
	}
}
