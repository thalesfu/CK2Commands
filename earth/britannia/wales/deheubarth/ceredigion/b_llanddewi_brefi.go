package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰瑟威布雷维Llanddewi_brefiBarony struct {
	feud.BaseBarony
}

var BLlanddewi_brefi兰瑟威布雷维 feud.Barony = &兰瑟威布雷维Llanddewi_brefiBarony{}

func init() {
    f := BLlanddewi_brefi兰瑟威布雷维.(*兰瑟威布雷维Llanddewi_brefiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanddewi_brefi",
		TitleName: "兰瑟威布雷维",
		TitleCode: "b_llanddewi_brefi",
	}
}
