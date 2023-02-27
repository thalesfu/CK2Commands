package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏拉杂孔XalazakungBarony struct {
	feud.BaseBarony
}

var BXalazakung夏拉杂孔 feud.Barony = &夏拉杂孔XalazakungBarony{}

func init() {
    f := BXalazakung夏拉杂孔.(*夏拉杂孔XalazakungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xalazakung",
		TitleName: "夏拉杂孔",
		TitleCode: "b_xalazakung",
	}
}
