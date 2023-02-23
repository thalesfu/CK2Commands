package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉让LabrangBarony struct {
	feud.BaseBarony
}

var BLabrang拉让 feud.Barony = &拉让LabrangBarony{}

func init() {
	f := BLabrang拉让.(*拉让LabrangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labrang",
		TitleName: "拉让",
		TitleCode: "b_labrang",
	}
}
