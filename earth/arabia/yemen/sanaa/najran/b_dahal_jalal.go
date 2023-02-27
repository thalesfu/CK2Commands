package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达哈勒贾拉勒Dahal_jalalBarony struct {
	feud.BaseBarony
}

var BDahal_jalal达哈勒贾拉勒 feud.Barony = &达哈勒贾拉勒Dahal_jalalBarony{}

func init() {
    f := BDahal_jalal达哈勒贾拉勒.(*达哈勒贾拉勒Dahal_jalalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dahal_jalal",
		TitleName: "达哈勒贾拉勒",
		TitleCode: "b_dahal_jalal",
	}
}
