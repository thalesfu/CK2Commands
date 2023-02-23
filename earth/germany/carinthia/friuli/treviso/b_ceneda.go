package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡内达CenedaBarony struct {
	feud.BaseBarony
}

var BCeneda卡内达 feud.Barony = &卡内达CenedaBarony{}

func init() {
	f := BCeneda卡内达.(*卡内达CenedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ceneda",
		TitleName: "卡内达",
		TitleCode: "b_ceneda",
	}
}
