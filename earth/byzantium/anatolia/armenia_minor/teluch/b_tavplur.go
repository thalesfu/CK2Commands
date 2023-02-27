package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔夫普鲁尔TavplurBarony struct {
	feud.BaseBarony
}

var BTavplur塔夫普鲁尔 feud.Barony = &塔夫普鲁尔TavplurBarony{}

func init() {
    f := BTavplur塔夫普鲁尔.(*塔夫普鲁尔TavplurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tavplur",
		TitleName: "塔夫普鲁尔",
		TitleCode: "b_tavplur",
	}
}
