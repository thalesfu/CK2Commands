package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅葛布拉MergeblaBarony struct {
	feud.BaseBarony
}

var BMergebla梅葛布拉 feud.Barony = &梅葛布拉MergeblaBarony{}

func init() {
	f := BMergebla梅葛布拉.(*梅葛布拉MergeblaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mergebla",
		TitleName: "梅葛布拉",
		TitleCode: "b_mergebla",
	}
}
