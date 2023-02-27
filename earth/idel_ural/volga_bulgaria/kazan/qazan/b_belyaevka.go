package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别利亚耶夫卡BelyaevkaBarony struct {
	feud.BaseBarony
}

var BBelyaevka别利亚耶夫卡 feud.Barony = &别利亚耶夫卡BelyaevkaBarony{}

func init() {
    f := BBelyaevka别利亚耶夫卡.(*别利亚耶夫卡BelyaevkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belyaevka",
		TitleName: "别利亚耶夫卡",
		TitleCode: "b_belyaevka",
	}
}
