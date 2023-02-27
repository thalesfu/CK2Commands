package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔瓦拉TawalaBarony struct {
	feud.BaseBarony
}

var BTawala塔瓦拉 feud.Barony = &塔瓦拉TawalaBarony{}

func init() {
    f := BTawala塔瓦拉.(*塔瓦拉TawalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tawala",
		TitleName: "塔瓦拉",
		TitleCode: "b_tawala",
	}
}
