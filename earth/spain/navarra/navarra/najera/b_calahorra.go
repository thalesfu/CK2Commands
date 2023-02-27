package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉奥拉CalahorraBarony struct {
	feud.BaseBarony
}

var BCalahorra卡拉奥拉 feud.Barony = &卡拉奥拉CalahorraBarony{}

func init() {
    f := BCalahorra卡拉奥拉.(*卡拉奥拉CalahorraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calahorra",
		TitleName: "卡拉奥拉",
		TitleCode: "b_calahorra",
	}
}
