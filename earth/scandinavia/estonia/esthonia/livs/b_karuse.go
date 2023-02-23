package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡鲁瑟KaruseBarony struct {
	feud.BaseBarony
}

var BKaruse卡鲁瑟 feud.Barony = &卡鲁瑟KaruseBarony{}

func init() {
	f := BKaruse卡鲁瑟.(*卡鲁瑟KaruseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karuse",
		TitleName: "卡鲁瑟",
		TitleCode: "b_karuse",
	}
}
