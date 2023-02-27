package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔皮CarpiBarony struct {
	feud.BaseBarony
}

var BCarpi卡尔皮 feud.Barony = &卡尔皮CarpiBarony{}

func init() {
    f := BCarpi卡尔皮.(*卡尔皮CarpiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carpi",
		TitleName: "卡尔皮",
		TitleCode: "b_carpi",
	}
}
