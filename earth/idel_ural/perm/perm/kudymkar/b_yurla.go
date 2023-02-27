package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤尔拉YurlaBarony struct {
	feud.BaseBarony
}

var BYurla尤尔拉 feud.Barony = &尤尔拉YurlaBarony{}

func init() {
    f := BYurla尤尔拉.(*尤尔拉YurlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yurla",
		TitleName: "尤尔拉",
		TitleCode: "b_yurla",
	}
}
