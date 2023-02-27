package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉纳凡福尔Llanafan_fawrBarony struct {
	feud.BaseBarony
}

var BLlanafan_fawr拉纳凡福尔 feud.Barony = &拉纳凡福尔Llanafan_fawrBarony{}

func init() {
    f := BLlanafan_fawr拉纳凡福尔.(*拉纳凡福尔Llanafan_fawrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanafan_fawr",
		TitleName: "拉纳凡福尔",
		TitleCode: "b_llanafan_fawr",
	}
}
