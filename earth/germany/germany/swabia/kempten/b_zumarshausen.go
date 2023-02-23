package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚斯马斯豪森ZumarshausenBarony struct {
	feud.BaseBarony
}

var BZumarshausen楚斯马斯豪森 feud.Barony = &楚斯马斯豪森ZumarshausenBarony{}

func init() {
	f := BZumarshausen楚斯马斯豪森.(*楚斯马斯豪森ZumarshausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zumarshausen",
		TitleName: "楚斯马斯豪森",
		TitleCode: "b_zumarshausen",
	}
}
