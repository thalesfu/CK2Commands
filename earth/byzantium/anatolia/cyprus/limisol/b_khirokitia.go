package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基罗基蒂亚KhirokitiaBarony struct {
	feud.BaseBarony
}

var BKhirokitia基罗基蒂亚 feud.Barony = &基罗基蒂亚KhirokitiaBarony{}

func init() {
    f := BKhirokitia基罗基蒂亚.(*基罗基蒂亚KhirokitiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khirokitia",
		TitleName: "基罗基蒂亚",
		TitleCode: "b_khirokitia",
	}
}
