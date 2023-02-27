package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶加雪夫Yirga_chefeBarony struct {
	feud.BaseBarony
}

var BYirga_chefe耶加雪夫 feud.Barony = &耶加雪夫Yirga_chefeBarony{}

func init() {
    f := BYirga_chefe耶加雪夫.(*耶加雪夫Yirga_chefeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yirga_chefe",
		TitleName: "耶加雪夫",
		TitleCode: "b_yirga_chefe",
	}
}
