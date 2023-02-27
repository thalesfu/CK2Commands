package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷西加洛TresigalloBarony struct {
	feud.BaseBarony
}

var BTresigallo特雷西加洛 feud.Barony = &特雷西加洛TresigalloBarony{}

func init() {
    f := BTresigallo特雷西加洛.(*特雷西加洛TresigalloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tresigallo",
		TitleName: "特雷西加洛",
		TitleCode: "b_tresigallo",
	}
}
