package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里卡塔TrikataBarony struct {
	feud.BaseBarony
}

var BTrikata特里卡塔 feud.Barony = &特里卡塔TrikataBarony{}

func init() {
    f := BTrikata特里卡塔.(*特里卡塔TrikataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trikata",
		TitleName: "特里卡塔",
		TitleCode: "b_trikata",
	}
}
