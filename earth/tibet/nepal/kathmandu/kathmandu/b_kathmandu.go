package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加德满都KathmanduBarony struct {
	feud.BaseBarony
}

var BKathmandu加德满都 feud.Barony = &加德满都KathmanduBarony{}

func init() {
    f := BKathmandu加德满都.(*加德满都KathmanduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kathmandu",
		TitleName: "加德满都",
		TitleCode: "b_kathmandu",
	}
}
