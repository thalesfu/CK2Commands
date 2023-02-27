package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷拉BurelaBarony struct {
	feud.BaseBarony
}

var BBurela布雷拉 feud.Barony = &布雷拉BurelaBarony{}

func init() {
    f := BBurela布雷拉.(*布雷拉BurelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burela",
		TitleName: "布雷拉",
		TitleCode: "b_burela",
	}
}
