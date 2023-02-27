package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙金_吉雷ShagingireiBarony struct {
	feud.BaseBarony
}

var BShagingirei沙金_吉雷 feud.Barony = &沙金_吉雷ShagingireiBarony{}

func init() {
    f := BShagingirei沙金_吉雷.(*沙金_吉雷ShagingireiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shagingirei",
		TitleName: "沙金_吉雷",
		TitleCode: "b_shagingirei",
	}
}
