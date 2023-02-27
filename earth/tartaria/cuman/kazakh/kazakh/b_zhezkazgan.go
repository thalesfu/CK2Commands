package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热兹卡兹甘ZhezkazganBarony struct {
	feud.BaseBarony
}

var BZhezkazgan热兹卡兹甘 feud.Barony = &热兹卡兹甘ZhezkazganBarony{}

func init() {
    f := BZhezkazgan热兹卡兹甘.(*热兹卡兹甘ZhezkazganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhezkazgan",
		TitleName: "热兹卡兹甘",
		TitleCode: "b_zhezkazgan",
	}
}
