package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡里萨CarissaBarony struct {
	feud.BaseBarony
}

var BCarissa卡里萨 feud.Barony = &卡里萨CarissaBarony{}

func init() {
	f := BCarissa卡里萨.(*卡里萨CarissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carissa",
		TitleName: "卡里萨",
		TitleCode: "b_carissa",
	}
}
