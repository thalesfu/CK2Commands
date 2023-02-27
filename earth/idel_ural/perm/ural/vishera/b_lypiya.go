package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷皮亚LypiyaBarony struct {
	feud.BaseBarony
}

var BLypiya雷皮亚 feud.Barony = &雷皮亚LypiyaBarony{}

func init() {
    f := BLypiya雷皮亚.(*雷皮亚LypiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lypiya",
		TitleName: "雷皮亚",
		TitleCode: "b_lypiya",
	}
}
