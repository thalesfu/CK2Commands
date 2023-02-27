package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈卡拉HaklaBarony struct {
	feud.BaseBarony
}

var BHakla哈卡拉 feud.Barony = &哈卡拉HaklaBarony{}

func init() {
    f := BHakla哈卡拉.(*哈卡拉HaklaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hakla",
		TitleName: "哈卡拉",
		TitleCode: "b_hakla",
	}
}
