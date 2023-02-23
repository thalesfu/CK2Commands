package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂拉丘地TiliatepeBarony struct {
	feud.BaseBarony
}

var BTiliatepe蒂拉丘地 feud.Barony = &蒂拉丘地TiliatepeBarony{}

func init() {
	f := BTiliatepe蒂拉丘地.(*蒂拉丘地TiliatepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiliatepe",
		TitleName: "蒂拉丘地",
		TitleCode: "b_tiliatepe",
	}
}
