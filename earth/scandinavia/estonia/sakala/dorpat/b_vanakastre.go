package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦纳_卡斯特雷VanakastreBarony struct {
	feud.BaseBarony
}

var BVanakastre瓦纳_卡斯特雷 feud.Barony = &瓦纳_卡斯特雷VanakastreBarony{}

func init() {
    f := BVanakastre瓦纳_卡斯特雷.(*瓦纳_卡斯特雷VanakastreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vanakastre",
		TitleName: "瓦纳_卡斯特雷",
		TitleCode: "b_vanakastre",
	}
}
